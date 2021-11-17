package main

import (
	"fmt"
	"github.com/sethvargo/go-githubactions"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"time"
)

func main() {
	email := githubactions.GetInput("email")
	passwd := githubactions.GetInput("passwd")
	checkIn(email, passwd, "")
}

func checkIn(email string, passwd string, code string) {
	client := http.Client{
		Timeout: time.Second * 100,
	}
	jar, err := cookiejar.New(nil)
	if err != nil {
		panic(err)
	}
	client.Jar = jar

	_, err = client.PostForm("https://www.cordcloud.biz/auth/login", url.Values{"email": {email}, "passwd": {passwd}, "code": {code}})
	if err != nil {
		fmt.Println("登录失败，请重试")
	}
	resp, err := client.Post("https://www.cordcloud.biz/user/checkin", "application/json", nil)
	if err != nil {
		fmt.Println("续命失败，请重试")
	} else {
		defer resp.Body.Close()
		bodyStr, _ := ioutil.ReadAll(resp.Body)
		fmt.Println(string(bodyStr))
	}
}
