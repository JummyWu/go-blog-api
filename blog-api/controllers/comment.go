package controllers

import (
	"go-blog-api/blog-api/models"
	"time"

	"github.com/astaxie/beego/orm"

	"github.com/astaxie/beego/logs"

	"github.com/gofrs/uuid"

	"github.com/astaxie/beego"
)

/*
CommentController 评论模块
*/
type CommentController struct {
	beego.Controller
}

/*
URLMapping 评论模块URL
*/
func (c *CommentController) URLMapping() {
	c.Mapping("PostComment", c.PostComment)
	c.Mapping("PostCommentReply", c.PostCommentReply)
	c.Mapping("GetPostComment", c.Get)
	c.Mapping("GetComment", c.GetComment)
	c.Mapping("GetCommentReply", c.GetCommentReply)
}

/*
PostComment /api/post/comment/
*/
// @router /api/post/comment/
func (c *CommentController) PostComment() {
	token := c.Ctx.Request.Header.Get("token")
	userID := models.GetUID(token)
	if userID == "" {
		c.Data["json"] = models.Message{Code: 200, Result: "没有这个用户", Data: nil}
		c.ServeJSON()
	}
	pid := c.GetString("pid")
	uid, error := uuid.NewV4()
	if error != nil {
		logs.Info(error)
	}
	cuid := "comment_" + uid.String()
	body := c.GetString("body")
	o := orm.NewOrm()
	post := models.Post{Uid: pid}
	err := o.Read(&post, "Uid")
	if err == orm.ErrNoRows {
		c.Data["json"] = models.Message{Code: 200, Result: "没有这文章,所以无法评论", Data: nil}
		c.ServeJSON()
	}
	comment := new(models.Comment)
	comment.Uid = cuid
	comment.UserId = userID
	comment.Body = body
	comment.PostId = pid
	comment.Time = time.Now()
	comment.Thumbs = 0
	logs.Info(comment)
	logs.Info(o.Insert(comment))
	c.Data["json"] = models.Message{Code: 200, Result: "评论成功", Data: comment}
	c.ServeJSON()
}

/*
PostCommentReply /api/post/comment_reply/
*/
// @router /api/post/comment_reply/
func (c *CommentController) PostCommentReply() {
	token := c.Ctx.Request.Header.Get("token")
	userID := models.GetUID(token)
	if userID == "" {
		c.Data["json"] = models.Message{Code: 200, Result: "没有这个用户", Data: nil}
		c.ServeJSON()
	}
	o := orm.NewOrm()
	pid := c.GetString("pid")
	post := models.Post{Uid: pid}
	err := o.Read(&post, "Uid")
	if err == orm.ErrNoRows {
		c.Data["json"] = models.Message{Code: 200, Result: "没有这文章,所以无法评论", Data: nil}
		c.ServeJSON()
	}
	commentID := c.GetString("commentId")
	comment := models.Comment{PostId: pid, Uid: commentID}
	errorComment := o.Read(&comment, "Uid")
	if errorComment == orm.ErrNoRows {
		c.Data["json"] = models.Message{Code: 200, Result: "没有这文章,所以无法评论", Data: nil}
		c.ServeJSON()
	}
	uid, err := uuid.NewV4()
	if err != nil {
		logs.Info(err)
	}
	cuid := "comment_rely_" + uid.String()
	body := c.GetString("body")
	commentReply := new(models.CommentReply)
	commentReply.Uid = cuid
	commentReply.Body = body
	commentReply.UserId = userID
	commentReply.Thumbs = 0
	commentReply.Time = time.Now()
	commentReply.CommentId = commentID
	number, err := o.Insert(commentReply)
	if err == orm.ErrNoRows {
		logs.Info(err)
	}
	logs.Info(number)
	c.Data["json"] = models.Message{Code: 200, Result: "二级评论成功", Data: commentReply}
	c.ServeJSON()
}

/*
RoplyViews 回复的视图
*/
type RoplyViews struct {
	Id        int         `json:"Id"`
	Uid       string      `json:"Uid"`
	User      models.User `json:"User"`
	CommentId string      `json:"CommentId"`
	Thumbs    int         `json:"Thumbs"`
	Body      string      `json:"Body"`
	Time      time.Time   `json:"Time"`
}

/*
CommentViews 一级评论的视图
*/
type CommentViews struct {
	Id     int          `json:"Id"`
	Uid    string       `json:"Uid"`
	User   models.User  `json: "User"`
	PostId string       `json:"PostId"`
	Thumbs int          `json:"Thumbs"`
	Body   string       `json:"Body"`
	Time   time.Time    `json:"Time"`
	Reply  []RoplyViews `json:"Reply"`
}

/*
Get /api/post/comment/:id([0-9]+)
*/
// @router /api/post/comment/
func (c *CommentController) Get() {
	pid := c.GetString("pid")
	o := orm.NewOrm()
	var jsonComment []CommentViews
	var comments []*models.Comment
	number, err := o.QueryTable("comment").OrderBy("-Id").Filter("post_id", pid).All(&comments)
	if err != nil {
		logs.Info(number)
	}
	for _, comment := range comments {
		var jsonCommentReply []RoplyViews
		var CommentReplys []*models.CommentReply
		logs.Info(comment)
		numbers, err := o.QueryTable("comment_reply").Filter("comment_id", comment.Uid).All(&CommentReplys)
		if err != nil {
			logs.Info(numbers)
		}
		for _, CommentReply := range CommentReplys {
			us := models.User{Uid: CommentReply.UserId}
			logs.Info(o.Read(&us, "Uid"))
			jsonCommentReply = append(jsonCommentReply, RoplyViews{Id: CommentReply.Id, Uid: CommentReply.Uid, User: us, CommentId: CommentReply.CommentId, Thumbs: CommentReply.Thumbs, Body: CommentReply.Body, Time: CommentReply.Time})
		}
		us := models.User{Uid: comment.UserId}
		logs.Info(o.Read(&us, "Uid"))
		jsonComment = append(jsonComment, CommentViews{Id: comment.Id, Uid: comment.Uid, User: us, PostId: comment.PostId, Thumbs: comment.Thumbs, Body: comment.Body, Time: comment.Time, Reply: jsonCommentReply})
	}
	logs.Info(pid)
	c.Data["json"] = models.Message{Code: 200, Result: "获取成功", Data: jsonComment}
	c.ServeJSON()
}

/*
GetComment /api/comment/
*/
func (c *CommentController) GetComment() {
	token := c.Ctx.Request.Header.Get("token")
	_, isAdmin := models.IsAdmin(token)
	if isAdmin == false {
		c.Data["json"] = models.Message{Code: 301, Result: "你不是超级管理员", Data: nil}
		c.ServeJSON()
	} else {
		var jsonComment []*models.Comment
		o := orm.NewOrm()
		number, err := o.QueryTable("comment").OrderBy("-Id").All(jsonComment)
		if err != nil {
			logs.Info(number)
		}
		c.Data["json"] = models.Message{Code: 200, Result: "查询成功", Data: jsonComment}
		c.ServeJSON()
	}
}

/*
GetCommentReply /api/comment/reply/
*/
func (c *CommentController) GetCommentReply() {
	token := c.Ctx.Request.Header.Get("token")
	_, isAdmin := models.IsAdmin(token)
	if isAdmin == false {
		c.Data["json"] = models.Message{Code: 301, Result: "你不是超级管理员", Data: nil}
		c.ServeJSON()
	} else {
		var jsonReply []*models.CommentReply
		o := orm.NewOrm()
		number, err := o.QueryTable("comment_reply").OrderBy("-Id").All(jsonReply)
		if err != nil {
			logs.Info(number)
		}
		c.Data["json"] = models.Message{Code: 200, Result: "查询成功", Data: jsonReply}
		c.ServeJSON()
	}
}
