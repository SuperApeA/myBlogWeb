package router

import (
	"myBlogWeb/server/views/api"
	"myBlogWeb/server/views/view"
	"net/http"

	"myBlogWeb/config"
)

// InitRouter 设置路由
func InitRouter() {
	http.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir(config.AppLocalPath+"/viewsrc/public/resource"))))
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir(config.AppLocalPath+"/viewsrc/public/resource/css"))))
	http.Handle("/fonts/", http.StripPrefix("/fonts/", http.FileServer(http.Dir(config.AppLocalPath+"/viewsrc/public/resource/fonts"))))
	http.Handle("/images/", http.StripPrefix("/lib/", http.FileServer(http.Dir(config.AppLocalPath+"/viewsrc/public/resource/images"))))
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir(config.AppLocalPath+"/viewsrc/public/resource/js"))))
	http.Handle("/lib/", http.StripPrefix("/lib/", http.FileServer(http.Dir(config.AppLocalPath+"/viewsrc/public/resource/lib"))))
	http.Handle("/plugins/", http.StripPrefix("/plugins/", http.FileServer(http.Dir(config.AppLocalPath+"/viewsrc/public/resource/plugins"))))
	http.Handle("/markdown/image/", http.StripPrefix("/markdown/image/", http.FileServer(http.Dir(config.AppLocalPath+"/viewsrc/markdown/image"))))

	http.HandleFunc("/index.html", view.HTML.IndexHtmlResponse)
	http.HandleFunc("/category/", view.HTML.CategoryHtmlResponse)
	http.HandleFunc("/post/", view.HTML.PostDetailHtmlResponse)
	http.HandleFunc("/login", view.HTML.LoginHtmlResponse)
	http.HandleFunc("/writing", view.HTML.WritingHtmlResponse)

	http.HandleFunc("/api/v1/login", api.API.LoginApiResponse)
	http.HandleFunc("/api/v1/post", api.API.PostApiResponse)
	http.HandleFunc("/api/v1/post/", api.API.PostApiResponse)
	http.HandleFunc("/api/v1/post-uploadfile", api.API.PostUploadFileApiResponse)

}
