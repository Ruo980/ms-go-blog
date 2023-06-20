// Package models @Program: ms-go-blog
// @Author: FunCoder
// @Create: 2023-06-11 11:03
// @Description:
package models

import (
	//标准库
	"html/template"
	"time"

	//第三方库

	//本地库
	"ms-go-blog/config"
)

type Post struct {
	Pid        int       `json:"pid"`        // æç« ID
	Title      string    `json:"title"`      // æç« ID
	Slug       string    `json:"slug"`       // èªå®ä¹é¡µé¢ path
	Content    string    `json:"content"`    // æç« çhtml
	Markdown   string    `json:"markdown"`   // æç« çMarkdown
	CategoryId int       `json:"categoryId"` //åç±»id
	UserId     int       `json:"userId"`     //ç¨æ·id
	ViewCount  int       `json:"viewCount"`  //æ¥çæ¬¡æ°
	Type       int       `json:"type"`       //æç« ç±»å 0 æ®éï¼1 èªå®ä¹æç«
	CreateAt   time.Time `json:"createAt"`   // åå»ºæ¶é´
	UpdateAt   time.Time `json:"updateAt"`   // æ´æ°æ¶é´
}

type PostMore struct {
	Pid          int           `json:"pid"`          // æç« ID
	Title        string        `json:"title"`        // æç« ID
	Slug         string        `json:"slug"`         // èªå®ä¹é¡µé¢ path
	Content      template.HTML `json:"content"`      // æç« çhtml
	CategoryId   int           `json:"categoryId"`   // æç« çMarkdown
	CategoryName string        `json:"categoryName"` // åç±»å
	UserId       int           `json:"userId"`       // ç¨æ·id
	UserName     string        `json:"userName"`     // ç¨æ·å
	ViewCount    int           `json:"viewCount"`    // æ¥çæ¬¡æ°
	Type         int           `json:"type"`         // æç« ç±»å 0 æ®éï¼1 èªå®ä¹æç«
	CreateAt     string        `json:"createAt"`
	UpdateAt     string        `json:"updateAt"`
}

type PostReq struct {
	Pid        int    `json:"pid"`
	Title      string `json:"title"`
	Slug       string `json:"slug"`
	Content    string `json:"content"`
	Markdown   string `json:"markdown"`
	CategoryId int    `json:"categoryId"`
	UserId     int    `json:"userId"`
	Type       int    `json:"type"`
}

type SearchResp struct {
	Pid   int    `orm:"pid" json:"pid"` // æç« ID
	Title string `orm:"title" json:"title"`
}

type PostRes struct {
	config.Viewer
	config.SystemConfig
	Article PostMore
}
