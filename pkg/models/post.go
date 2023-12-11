package models

import (
	"github.com/mattot-the-builder/totgram-api/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Post struct {
	gorm.Model
	UserId  int64  `json:"user_id"`
	Caption string `json:"caption"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Post{})
}

func (u *Post) CreatePost() *Post {
	db.Create(u)
	return u
}

func GetAllPosts() []Post {
	var posts []Post
	db.Find(&posts)
	return posts
}

func GetPostById(Id int64) (*Post, *gorm.DB) {
	var post Post
	db := db.Where("ID = ?", Id).Find(&post)
	return &post, db
}

func DeletePost(ID int64) Post {
	var post Post
	db.Where("ID = ?", ID).Delete(post)
	return post
}
