package models

import (
	"time"

	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

/*
Post ID, 文章的UID, 文章的标题, 文章的简介, 文章的标签, 文章的分类, 文章的头图, 阅读量, 点赞数, 正文, markdown, 时间
*/
type Post struct {
	Id       int
	Uid      string
	Title    string
	Desc     string
	Tid      int
	Image    string
	Pv       int
	Likes    int
	Content  string `orm:"type(text)"`
	Markdown string `orm:"type(text)"`
	UserId   string
	Time     time.Time
}

/*
Tag ID, 标签名, 时间
*/
type Tag struct {
	Id     int
	Name   string
	UserId string
	Time   time.Time
}

/*
User Id, 用户UID, 用户名, 密码, 邮箱, 是否为管理员(0-否,1-是), 头像, githubId, 主页, 个人简介, GitHub地址, 地址, 注册时间
*/
type User struct {
	Id       int
	Uid      string
	Username string
	Password string
	Email    string
	IsAdmin  int
	Avatar   string
	GithubId float64
	Blog     string
	Bio      string
	GitPath  string
	Location string
	Time     time.Time
}

/*
Links Uid 站长邮箱 网站名字 网站的图片 网站链接 网站介绍 时间
*/
type Links struct {
	Id        int
	Uid       string
	Is        int
	Email     string
	HostName  string
	HostImage string
	HostUrl   string
	Body      string `orm:"type(text)"`
	Time      time.Time
}

/*
Comment id 评论的UID 评论人UID  评论文章 点赞数 正文 时间
*/
type Comment struct {
	Id     int
	Uid    string
	UserId string
	PostId string
	Thumbs int
	Body   string
	Time   time.Time
}

/*
CommentReply ID 回复评论的UID 用户ID 上级评论的UID 点赞数 正文 时间
*/
type CommentReply struct {
	Id        int
	Uid       string
	UserId    string
	CommentId string
	Thumbs    int
	Body      string
	Time      time.Time
}

/*
Category : ID UUID 用户的ID 分类名 时间
*/
type Category struct {
	Id     int
	Uid    string
	UserId string
	Name   string
	Time   time.Time
}

func init() {
	orm.RegisterModel(new(User), new(Post), new(Tag), new(Links), new(Comment), new(CommentReply), new(Category))
	re := orm.RunSyncdb("default", false, true)
	logs.Info(re)
	logs.Info("创建数据表成功")
}
