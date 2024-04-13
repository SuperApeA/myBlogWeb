package main

import (
	"log"
	"net/http"

	"myBlogWeb/server/router"
	views "myBlogWeb/server/views/common"
)

type IndexInfo struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

func main() {

	server := http.Server{
		//Addr: "127.0.0.1:12345",
		Addr: "0.0.0.0:12345",
	}

	// 加载前端模板
	views.InitHTMLTemplateCtl()
	// 设置路由
	router.InitRouter()

	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}

	return
}
