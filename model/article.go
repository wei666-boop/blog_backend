package model

import (
	"gorm.io/gorm"
)

type AuthorStruct struct {
	gorm.Model
	Name      string `json:"name"`
	Telephone string `json:"telephone"`
}
type Article struct {
	Title      string       `gorm:"type:text;not null;unique" json:"title"`
	ArticleID  int          `gorm:"unique;not null" json:"article_id"`
	Author     AuthorStruct `json:"author"`
	Content    string       `gorm:"type:text;not null" json:"content"`
	CreateTime string       `gorm:"not null" json:"create_time"`
}
