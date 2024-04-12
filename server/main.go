package main

import (
	"html/template"
	"log"
	"myBlogWeb/config"
	"myBlogWeb/server/models"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

type IndexInfo struct {
	Title string `json:"title"`
	Desc  string `json:"desc"`
}

func isODD(num int) bool {
	return num%2 == 0
}

func getNextName(nameList []string, index int) string {
	return nameList[index+1]
}

func getDate(layout string) string {
	return time.Now().Format(layout)
}

func IndexHtmlResponse(w http.ResponseWriter, r *http.Request) {
	t := template.New("index.html")

	//1.拿到项目的根路径
	path, _ := os.Getwd()
	path = filepath.Dir(path)
	index := filepath.Join(path, "/view/template/index.html")
	home := filepath.Join(path, "/view/template/home.html")
	header := filepath.Join(path, "/view/template/layout/header.html")
	footer := filepath.Join(path, "/view/template/layout/footer.html")
	personal := filepath.Join(path, "/view/template/layout/personal.html")
	post := filepath.Join(path, "/view/template/layout/post-list.html")
	pagination := filepath.Join(path, "/view/template/layout/pagination.html")
	t.Funcs(template.FuncMap{"isODD": isODD, "getNextName": getNextName, "date": getDate})
	t, err := t.ParseFiles(index, home, header, footer, personal, post, pagination)
	if err != nil {
		log.Println("解析首页报错", err)
	}

	var categoryList = []models.Category{
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
	var hr = &models.HomeData{
		Viewer:       config.GetConfig().Viewer,
		CategoryList: categoryList,
		Posts:        posts,
		Total:        1,
		Page:         1,
		Pages:        []int{1},
		PageEnd:      true,
	}

	if err := t.Execute(w, hr); err != nil {
		log.Println("index返回前端报错", err)
	}

}

func main() {

	server := http.Server{
		//Addr: "127.0.0.1:12345",
		Addr: "0.0.0.0:12345",
	}

	http.HandleFunc("/", IndexHtmlResponse)
	http.Handle("/resource/", http.StripPrefix("/resource", http.FileServer(http.Dir(config.AppLocalPath+"/view/public/resource"))))

	if err := server.ListenAndServe(); err != nil {
		log.Println(err)
	}

	return
}
