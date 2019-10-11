package util

import (
	"go-blog-api/blog-api/models"
)

/*
UserToViews : user转userview
*/
func UserToViews(c models.User) *models.UserView {
	userView := new(models.UserView)
	userView.Id = c.Id
	userView.Uid = c.Uid
	userView.Username = c.Username
	userView.Email = c.Email
	userView.Avatar = c.Avatar
	userView.GithubId = c.GithubId
	userView.Blog = c.Blog
	userView.Bio = c.Bio
	userView.GitPath = c.GitPath
	userView.Location = c.Location
	userView.Time = c.Time
	return userView
}

/*
PostToViews : 文章转换类
*/
func PostToViews(c *models.Post, user models.UserView, tag models.TagView) *models.PostView {
	postView := new(models.PostView)
	postView.Id = c.Id
	postView.Uid = c.Uid
	postView.Title = c.Title
	postView.Desc = c.Desc
	postView.Image = c.Image
	postView.Pv = c.Pv
	postView.Likes = c.Likes
	postView.Content = c.Content
	postView.Markdown = c.Markdown
	postView.Time = c.Time
	postView.UserView = user
	postView.TagView = tag
	return postView
}

/*
TagToView : tag 转 tagview
*/
func TagToView(c models.Tag, user models.UserView) *models.TagView {
	tagView := new(models.TagView)
	tagView.Id = c.Id
	tagView.Name = c.Name
	tagView.Time = c.Time
	tagView.UserId = c.UserId
	tagView.UserView = user
	return tagView
}
