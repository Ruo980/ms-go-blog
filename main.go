// @program: ms-go-blog
// @author: FunCoder
// @create: 2023-06-09 13:16
// @Description: 程序入口，一个项目只能有一个入口

package main

import (
	"log"
	"ms-go-blog/router"
	"net/http"
)

func main() {
	//程序入口，一个项目只能有一个入口
	//web 程序，http 协议 ip port
	server := http.Server{
		Addr: "127.0.0.1:8081",
	}
	router.Router() //将相关功能代码放入一个模块中调用处理
	//server.ListenAndServe()启动服务并监听端口，返回启动服务时发生的错误
	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
