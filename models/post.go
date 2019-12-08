package models

import (
	helper "github.com/AliBedaer/go-blog/helpers"
)

type Post struct {
	Id      uint   `validate:"required"`
	Title   string `json:"title" validate:"required,alphanum"`
	Desc    string `json:"desc" validate:"required,alphanum"`
	Content string `json:"content" validate:"required,alphanum"`
}

func (p Post) CreatePost(us *Post) {
	db := helper.DBConnect()

	db.Create(&us)

}
