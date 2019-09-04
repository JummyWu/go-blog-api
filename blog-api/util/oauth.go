package util

import (
	"encoding/json"
	"fmt"
	"go-blog-api/blog-api/models"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

/*
GetGitHubToken 获取access_token
*/
func GetGitHubToken(code string) string {
	url := "https://github.com/login/oauth/access_token?client_id=" + "0c40c7742f27ca32cc50" + "&client_secret=" + "dc9a8c3ba4da1f5469096e3f0bbf089a0578c546" + "&code=" + code
	fmt.Println(url)
	client := &http.Client{}
	reqest, err := http.NewRequest("POST", url, nil)
	reqest.Header.Add("accept", "application/json")
	if err != nil {
		panic(err)
	}
	resp, _ := client.Do(reqest)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}
	var v interface{}
	json.Unmarshal(body, &v)
	data := v.(map[string]interface{})
	return data["access_token"].(string)
}

/*
GetGitHubUser 获取用户信息保存
*/
func GetGitHubUser(accessToken string) *models.User {
	url := "https://api.github.com/user?access_token=" + accessToken
	fmt.Println(url)
	client := &http.Client{}
	reqest, err := http.NewRequest("GET", url, nil)
	reqest.Header.Add("accept", "application/json")
	if err != nil {
		panic(err)
	}
	resp, _ := client.Do(reqest)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// handle error
	}
	var v interface{}
	json.Unmarshal(body, &v)
	data := v.(map[string]interface{})
	user := models.User{GithubId: data["id"].(float64)}
	o := orm.NewOrm()
	logs.Info(o.Read(&user, "GithubId"))
	fmt.Println(o.Read(&user, "GithubId"))
	if user.Id == 0 {
		user := new(models.User)
		user.Uid = models.UID()
		user.IsAdmin = 0
		user.Avatar = data["avatar_url"].(string)
		user.Bio = data["bio"].(string)
		user.Blog = data["blog"].(string)
		user.Email = data["email"].(string)
		user.GitPath = data["html_url"].(string)
		user.GithubId = data["id"].(float64)
		user.Location = data["location"].(string)
		user.Username = data["login"].(string)
		user.Time = time.Now()
		logs.Info(o.Insert(user))
		return user
	} else if user.Id > 0 {
		user.Avatar = data["avatar_url"].(string)
		user.Bio = data["bio"].(string)
		user.Blog = data["blog"].(string)
		user.Email = data["email"].(string)
		user.GitPath = data["html_url"].(string)
		user.Location = data["location"].(string)
		user.Username = data["login"].(string)
		user.Time = time.Now()
		logs.Info(o.Update(&user, "GithubId"))
		return &user
	} else {
		return nil
	}
}
