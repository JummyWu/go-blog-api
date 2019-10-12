package models

import "time"

// PostView : 文章展示列表
type PostView struct {
	Id           int          `json:"id"`
	Uid          string       `json:"uid"`
	Title        string       `json:"title"`
	Desc         string       `json:"desc"`
	Image        string       `json:"image"`
	Pv           int          `json:"pv"`
	Likes        int          `json:"likes"`
	Content      string       `json:"content"`
	Markdown     string       `json:"markdown"`
	Time         time.Time    `json:"time"`
	TagView      TagView      `json:"tag"`
	UserView     UserView     `json:"user"`
	CategoryView CategoryView `json:"category"`
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
	Email    string    `json:"email"`
	Avatar   string    `json:"avatar"`
	GithubId float64   `json:"githubId"`
	Blog     string    `json:"blog"`
	Bio      string    `json:"bio"`
	GitPath  string    `json:"gitPath"`
	Location string    `json:"location"`
	Time     time.Time `json:"time"`
}

//CategoryView : 分类列表
type CategoryView struct {
	Id       int       `json:"id"`
	Uid      string    `json:"uid"`
	Name     string    `json:"name"`
	UserId   string    `json:"userId"`
	UserView UserView  `json:"user"`
	Time     time.Time `json:"time"`
}
