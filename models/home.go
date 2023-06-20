// Package models @Program: ms-go-blog
// @Author: FunCoder
// @Create: 2023-06-11 22:11
// @Description:
package models

import "ms-go-blog/config"

// HomeResponse
//
//	@Description: 对应主页数据返回
type HomeResponse struct {
	config.Viewer
	Categorys []Category
	Posts     []PostMore
	Total     int
	Page      int
	Pages     []int
	PageEnd   bool
}
