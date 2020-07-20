package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Set("username", "")
	session.Set("password", "")
	session.Save()
	c.Redirect(302, "/")
}

func performLogin(c *gin.Context) {
	username := c.PostForm("username")
	username = strings.ToLower(username)
	password := c.PostForm("password")
	if isUserValid(username, password) {
		session := sessions.Default(c)
		session.Set("username", username)
		session.Set("password", password)
		session.Save()
		c.Redirect(302, "/news")
	} else if !isUserExists(username) {
		c.HTML(http.StatusBadRequest, "signin.html", gin.H{
			"ErrorTitle": "Wrong username",
			"ErrorMessage": "Username not exist",
		})
	} else {
		c.HTML(http.StatusBadRequest, "signin.html", gin.H{
			"ErrorTitle": "Wrong Account",
			"ErrorMessage": "Account information isn't correct",
		})
	}
}

func register(c *gin.Context) {
	username := c.PostForm("username")
	username = strings.ToLower(username)
	password := c.PostForm("password")
	passwordRepeat := c.PostForm("password-repeat")
	if !isUserExists(username) {
		if len(password) > 6 && len(username) > 6{
			if password == passwordRepeat {
				if !strings.ContainsAny(password, `='"<>`) {
					if err := registerNewUser(username, password); err == nil {
						session := sessions.Default(c)
						session.Set("username", username)
						session.Set("password", password)
						session.Save()
						c.Redirect(302, "/news")
					} else {
						c.HTML(http.StatusBadRequest, "signup.html", gin.H{
							"ErrorTitle": "Registration Failed",
							"ErrorMessage": err.Error(),
						})
					}
				} else {
					c.HTML(http.StatusBadRequest, "signup.html", gin.H{
						"ErrorTitle": "Wrong Password",
						"ErrorMessage": `You cant use = ' " < > in Password`,
					})
				}
			} else {
				c.HTML(http.StatusBadRequest, "signup.html", gin.H{
					"ErrorTitle": "Wrong Password",
					"ErrorMessage": "Passwords do not match",
				})
			}
		} else {
			c.HTML(http.StatusBadRequest, "signup.html", gin.H{
				"ErrorTitle": "Wrong Information",
				"ErrorMessage": "Password and Username longer than 6 characters",
			})
		}
	} else {
		c.HTML(http.StatusBadRequest, "signup.html", gin.H{
			"ErrorTitle": "Wrong Username",
			"ErrorMessage": "Username already exists",
		})
	}
}

func registerNewUser(username, password string) error {
	db, err := sql.Open("mysql", dbQuery)
	if err != nil {
		return err
	}
	defer db.Close()
	query := fmt.Sprintf("INSERT INTO users VALUES (\"%s\", \"%s\", \"300\", \"free\")", username, password)
	result, err := db.Exec(query)
		if err != nil {
		return err
	}
	err = registerDatabase(username)
	if err != nil {
		return err
	}
	n, err := result.RowsAffected()
	if n == 1 {
		return nil
	}
	return err
}

func registerDatabase(username string) error {
	db, err := sql.Open("mysql", dbQuery)
	if err != nil {
		return err
	}
	defer db.Close()
	query := fmt.Sprintf("CREATE TABLE %s_b (`os` text DEFAULT NULL, `ip` text DEFAULT NULL, `current_t` text DEFAULT NULL, `last_resp` text DEFAULT NULL) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;", username)
	_, err = db.Exec(query)
	if err != nil {
		return err
	}
	query = fmt.Sprintf("CREATE TABLE %s_t (`type` text DEFAULT NULL, `param` text DEFAULT NULL, `goal` text DEFAULT NULL, `now` text DEFAULT NULL, `tag` text DEFAULT NULL) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;", username)
	_, err = db.Exec(query)
	if err != nil {
		return err
	}
		return nil
}

func isExpExists(username string) bool {
	db, err := sql.Open("mysql", dbQuery)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	query := fmt.Sprintf("SELECT EXISTS(SELECT * from expdate WHERE username=\"%s\")", username)
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


func isUserExists(username string) bool {
	db, err := sql.Open("mysql", dbQuery)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	query := fmt.Sprintf("SELECT EXISTS(SELECT * from users WHERE username=\"%s\")", username)
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

func isUserValid(username, password string) bool {
	db, err := sql.Open("mysql", dbQuery)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	if isUserExists(username) {
		query := fmt.Sprintf("SELECT password from users WHERE username=\"%s\"", username)
		rows := db.QueryRow(query)
				var pw string
		err = rows.Scan(&pw)
		if err != nil {
			fmt.Println(err)
		}
		if pw == password {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

func isUserVIP(username, password string) bool {
	if isUserValid(username, password) {
		db, err := sql.Open("mysql", dbQuery)
		if err != nil {
			fmt.Println(err)
		}
		defer db.Close()
		if isUserExists(username) {
			query := fmt.Sprintf("SELECT status from users WHERE username=\"%s\"", username)
			rows := db.QueryRow(query)
						var status string
			err = rows.Scan(&status)
			if err != nil {
				fmt.Println(err)
			}
			if status == "vip" {
				return true
			} else {
				return false
			}
		} else {
			return false
		}
	} else {
		return false
	}
}

func getUserInfo(username string) string {
	db, err := sql.Open("mysql", dbQuery)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	if isUserExists(username) {
		query := fmt.Sprintf("SELECT status from users WHERE username=\"%s\"", username)
		rows := db.QueryRow(query)
				var status string
		err = rows.Scan(&status)
		if err != nil {
			fmt.Println(err)
		}
		return status
	} else {
		return ""
	}
	return ""
}

func getInterval(username string) string {
	db, err := sql.Open("mysql", dbQuery)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	if isUserExists(username) {
		query := fmt.Sprintf("SELECT inteval from users WHERE username=\"%s\"", username)
		rows := db.QueryRow(query)
				var interval string
		err = rows.Scan(&interval)
		if err != nil {
			fmt.Println(err)
		}
		return interval
	} else {
		return ""
	}
	return ""
}


func getExpDate(username string) string {
	db, err := sql.Open("mysql", dbQuery)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	if isUserExists(username) {
		query := fmt.Sprintf("SELECT date from expdate WHERE username=\"%s\"", username)
		rows := db.QueryRow(query)
				var exp string
		err = rows.Scan(&exp)
		if err != nil {
			fmt.Println(err)
		}
		return exp
	}
	return ""
}

func setVIP(username string) {
	db, err := sql.Open("mysql", dbQuery)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	query := fmt.Sprintf("UPDATE users SET status=\"vip\" WHERE username=\"%s\"", username)
	_, err = db.Exec(query)
		if err != nil {
		fmt.Println(err)
	}
}

func setFree(username string) {
	db, err := sql.Open("mysql", dbQuery)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	query := fmt.Sprintf("UPDATE users SET status=\"free\" WHERE username=\"%s\"", username)
	_, err = db.Exec(query)
		if err != nil {
		fmt.Println(err)
	}
}

func addMonth(username string) {
	var target string
	if isExpExists(username) {
		temp := getExpDate(username)
		fmt.Println(temp)
		tempTime, _ := time.Parse("2006-01-02 15:04:05", temp)
		loc, _ := time.LoadLocation("UTC")
		t := time.Now().In(loc)
		if tempTime.After(t) {
			tempTime = tempTime.Add(time.Hour * 24 * 30)
			target = tempTime.Format("2006-01-02 15:04:05")
		} else {
			loc, _ := time.LoadLocation("UTC")
			t = time.Now().In(loc)
			target = t.Add(time.Hour * 24 * 30).Format("2006-01-02 15:04:05")
		}
		fmt.Println(target)
		db, err := sql.Open("mysql", dbQuery)
		if err != nil {
			fmt.Println(err)
		}
		defer db.Close()
		query := fmt.Sprintf("UPDATE expdate SET date=\"%s\" WHERE username=\"%s\"", target, username)
		_, err = db.Exec(query)
				if err != nil {
			fmt.Println(err)
		}
	} else {
		loc, _ := time.LoadLocation("UTC")
		t := time.Now().In(loc)
		target = t.Add(time.Hour * 24 * 30).Format("2006-01-02 15:04:05")
		db, err := sql.Open("mysql", dbQuery)
		if err != nil {
			fmt.Println(err)
		}
		defer db.Close()
		query := fmt.Sprintf("INSERT INTO expdate VALUES (\"%s\", \"%s\")", username, target)
		result, err := db.Exec(query)
				if err != nil {
			fmt.Println(err)
		}
		n, err := result.RowsAffected()
		if n == 1 {
			return
		}
	}
}

func updateStatus(username string) {
	if isExpExists(username) {
		expdate := getExpDate(username)
		expDate, _ := time.Parse("2006-01-02 15:04:05", expdate)
		loc, _ := time.LoadLocation("UTC")
		t := time.Now().In(loc)
		if expDate.After(t) {
			setVIP(username)
		} else {
			setFree(username)
		}
	} else {
		setFree(username)
	}
}

func checkSerial(c *gin.Context) {
	session := sessions.Default(c)
	username := session.Get("username")
	password := session.Get("password")
	serialCode := c.PostForm("serial")
	if username != nil && password != nil {
		valid := isUserValid(username.(string), password.(string))
		if valid {
			if isSerialValid(serialCode) {
				addMonth(username.(string))
			}
			c.Redirect(302, "/user")
		} else {
			c.Redirect(302, "/signin")
		}
	} else {
		c.Redirect(302, "/signin")
	}
}

func postUser(c *gin.Context) {
	session := sessions.Default(c)
	username := session.Get("username")
	password := session.Get("password")
	if username != nil || password != nil {
		if isUserExists(username.(string)) {
			if isUserValid(username.(string), password.(string)) {
				if c.PostForm("interval") != "" {
					buildStub(c)
				} else if c.PostForm("serial") != "" {
					checkSerial(c)
				}
			}
		}
	}
}

func isSerialValid(serial string) bool {
	db, err := sql.Open("mysql", dbQuery)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	query := fmt.Sprintf("SELECT EXISTS(SELECT * from serial WHERE code=\"%s\")", serial)
	rows := db.QueryRow(query)
	var exist string
	err = rows.Scan(&exist)
	if err != nil {
		fmt.Println(err)
	}
	if exist == "1" {
		deleteSerial(serial)
		return true
	} else {
		return false
	}
}

func deleteSerial(serial string) {
	db, err := sql.Open("mysql", dbQuery)
	if err != nil {
		fmt.Print(err)
	}
	defer db.Close()
	query := fmt.Sprintf("DELETE FROM serial WHERE code=\"%s\"", serial)
	_, err = db.Exec(query)
	if err != nil {
		fmt.Print(err)
	}
}

func createSerial(c *gin.Context) {
	var list string
	count, _ := strconv.Atoi(c.Param("count"))
	for i := 1; i <= count; i++ {
		list += newCode() + "\n"
	}
	c.String(200, list)
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const (
	letterIdxBits = 6                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
)

func RandStringBytesMask(n int) string {
	b := make([]byte, n)
	for i := 0; i < n; {
		if idx := int(rand.Int63() & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i++
		}
	}
	return string(b)
}

func newCode() string {
	code := RandStringBytesMask(5) + "-" + RandStringBytesMask(5) + "-" + RandStringBytesMask(5) + "-" + RandStringBytesMask(5)
	db, err := sql.Open("mysql", dbQuery)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	query := fmt.Sprintf("INSERT INTO serial VALUES (\"%s\")", code)
	_, err = db.Exec(query)
	if err != nil {
		fmt.Println(err)
	}
	return code
}