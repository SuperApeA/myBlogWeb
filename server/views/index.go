package views

import (
	"log"
	"net/http"

	"myBlogWeb/config"
	"myBlogWeb/server/models"
	"myBlogWeb/server/views/common"
)

// IndexHtmlResponse index界面相应
func (*HTMLApi) IndexHtmlResponse(w http.ResponseWriter, r *http.Request) {
	t := common.GetHTMLTemplateCtl().Index

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
