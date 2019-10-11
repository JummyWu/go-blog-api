package models

import "time"

// PostView : 文章展示列表
type PostView struct {
	Id       int       `json:"id"`
	Uid      string    `json:"uid"`
	Title    string    `json:"title"`
	Desc     string    `json:"desc"`
	Tid      int       `json:"tid"`
	Image    string    `json:"image"`
	Pv       int       `json:"pv"`
	Likes    int       `json:"likes"`
	Content  string    `json:"content"`
	Markdown string    `json:"markdown"`
	UserId   string    `json:"userId"`
	Time     time.Time `json:"time"`
	TagView  TagView   `json:"Tag"`
	UserView UserView  `json:"User"`
}

//TagView : 便签展示列表
type TagView struct {
	Id       int       `json:"id"`
	Name     string    `json:"name"`
	UserId   string    `json:"userId"`
	UserView UserView  `json:"user"`
	Time     time.Time `json:"time"`
}

// UserView : 用户展示列表
type UserView struct {
	Id       int       `json:"id"`
	Uid      string    `json:"uid"`
	Username string    `json:"username"`
	Password string    `json:"password"`
	Email    string    `json:"email"`
	IsAdmin  int       `json:"isAdmin"`
	Avatar   string    `json:"avatar"`
	GithubId float64   `json:"githubId"`
	Blog     string    `json:"blog"`
	Bio      string    `json:"bio"`
	GitPath  string    `json:"gitPath"`
	Location string    `json:"location"`
	Time     time.Time `json:"time"`
}
