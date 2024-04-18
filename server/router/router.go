package router

import (
	"net/http"

	"myBlogWeb/config"
	"myBlogWeb/server/api"
	"myBlogWeb/server/views"
)

// InitRouter 设置路由
func InitRouter() {
	http.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir(config.AppLocalPath+"/viewsrc/public/resource"))))

	http.HandleFunc("/index.html", views.HTML.IndexHtmlResponse)
	http.HandleFunc("/category/", views.HTML.CategoryHtmlResponse)
	http.HandleFunc("/post/", views.HTML.PostDetailHtmlResponse)
	http.HandleFunc("/login", views.HTML.LoginHtmlResponse)
	http.HandleFunc("/writing", views.HTML.WritingHtmlResponse)

	http.HandleFunc("/api/v1/login", api.API.LoginApiResponse)
	http.HandleFunc("/api/v1/post", api.API.PostApiResponse)
	http.HandleFunc("/api/v1/post/", api.API.PostApiResponse)
}
