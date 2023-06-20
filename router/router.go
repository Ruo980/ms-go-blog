// Package router @Program: ms-go-blog
// @Author: FunCoder
// @Create: 2023-06-15 17:34
// @Description:
package router

import (
	"ms-go-blog/views"
	"net/http"
)

// Router
//
//	@Description: 主要用于定义接口，对外响应页面、数据、静态资源
func Router() {
	//1.页面 2.数据 3.静态资源
	http.HandleFunc("/", views.HTML.Index)
	http.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir("public/resource/"))))
}
