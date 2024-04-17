package router

import (
	"net/http"

	"myBlogWeb/config"
	"myBlogWeb/server/views"
)

// InitRouter 设置路由
func InitRouter() {
	http.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir(config.AppLocalPath+"/viewsrc/public/resource"))))
	http.HandleFunc("/index.html", views.HTML.IndexHtmlResponse)
	http.HandleFunc("/category/", views.HTML.CategoryHtmlResponse)
}
