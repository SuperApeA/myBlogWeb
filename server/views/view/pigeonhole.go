package view

import (
	"errors"
	"log"
	"net/http"

	"myBlogWeb/config"
	blogerr "myBlogWeb/error"
	"myBlogWeb/server/models"
	"myBlogWeb/server/sql"
	"myBlogWeb/server/views/common"
)

func GetPigeonholeResponseData() (*models.PigeonholeResponse, error) {
	// 文章
	postList, err := sql.GetAllPost()
	if err != nil {
		log.Printf("Get all post failed: %s\n", err)
		return &models.PigeonholeResponse{}, err
	}
	postMap := make(map[string][]models.Post)
	// 根据前端需要赋值
	for i, _ := range postList {
		var post models.Post
		post.Pid = postList[i].Pid
		post.Title = postList[i].Title
		post.CreateAt = postList[i].CreateAt
		monthStr := postList[i].CreateAt.Format("2006-01")
		if _, ok := postMap[monthStr]; !ok {
			postMap[monthStr] = []models.Post{}
		}
		postMap[monthStr] = append(postMap[monthStr], post)
	}
	// 类别
	categoryList, err := sql.GetAllCategory()
	if err != nil {
		log.Printf("Get all category failed: %s\n", err)
		return &models.PigeonholeResponse{}, err
	}

	var pigeonholeResponse = &models.PigeonholeResponse{
		Viewer:       config.GetConfig().Viewer,
		SystemConfig: config.GetConfig().System,
		CategoryList: categoryList,
		PostMap:      postMap,
	}
	return pigeonholeResponse, nil
}

// PigeonholeHtmlResponse 归档界面响应
func (*HTMLApi) PigeonholeHtmlResponse(w http.ResponseWriter, r *http.Request) {
	pigeonholeTemplate := common.GetHTMLTemplateCtl().Pigeonhole

	pigeonholeResponse, err := GetPigeonholeResponseData()
	if err != nil {
		if blogerr.GetClientError()[err.Error()] {
			pigeonholeTemplate.WriteError(w, err)
		} else {
			pigeonholeTemplate.WriteError(w, errors.New("系统内部错误，请联系管理员！"))
		}
		return
	}
	if err := pigeonholeTemplate.Execute(w, pigeonholeResponse); err != nil {
		log.Println("pigeonhole页面解析报错: ", err)
	}
}
