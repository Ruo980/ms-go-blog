// Package views @Program: ms-go-blog
// @Author: FunCoder
// @Create: 2023-06-15 17:42
// @Description:
package views

import (
	"html/template"
	"log"
	"ms-go-blog/config"
	"ms-go-blog/models"
	"net/http"
	"time"
)

type IndexData struct {
	Title string
	Desc  string
}

func IsODD(num int) bool {
	return num%2 == 0
}

func GetNextName(strs []string, index int) string {
	return strs[index+1]
}

func Date(layout string) string {
	return time.Now().Format(layout)
}

// 根目录请求处理函数
func (*HTMLApi) Index(w http.ResponseWriter, r *http.Request) {
	//配置模板渲染时数据：这些数据是给模板而不是前端因此不需要进行序列化
	var indexData IndexData
	indexData.Title = "博客"
	indexData.Desc = "欢迎来到博客系统"
	//将资源放置到go web中
	t := template.New("index.html")
	//1.拿到当前的路径
	path := config.Cfg.System.CurrentDir
	home := path + "/template/home.html"
	header := path + "/template/layout/header.html"
	footer := path + "/template/layout/footer.html"
	pagination := path + "/template/layout/pagination.html"
	personal := path + "/template/layout/personal.html"
	post := path + "/template/layout/post-list.html"
	//将定义的函数渲染到模板中的写法
	t.Funcs(template.FuncMap{"isODD": IsODD, "getNextName": GetNextName, "date": Date})
	t, err := t.ParseFiles(path+"/template/index.html", home, header, footer, pagination, personal, post)
	if err != nil {
		log.Println(err)
	}
	//页面上涉及到的所有的数据，必须有定义
	var categorys = []models.Category{
		{
			Cid:  1,
			Name: "go",
		},
	}
	var posts = []models.PostMore{
		{
			Pid:          1,
			Title:        "go博客",
			Content:      "内容",
			UserName:     "码神",
			ViewCount:    123,
			CreateAt:     "2022-02-20",
			CategoryId:   1,
			CategoryName: "go",
			Type:         0,
		},
	}
	var hr = &models.HomeResponse{
		config.Cfg.Viewer,
		categorys,
		posts,
		1,
		1,
		[]int{1},
		true,
	}
	//将数据带入模板并进行渲染、响应
	t.Execute(w, hr)
}
