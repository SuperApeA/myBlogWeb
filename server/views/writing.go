package views

import (
	"errors"
	"log"
	"net/http"

	"myBlogWeb/config"
	"myBlogWeb/server/models"
	"myBlogWeb/server/sql"
	"myBlogWeb/server/views/common"
)

func GetWritingResponseData() (*models.WritingHtmlResponse, error) {
	// 类别
	categoryList, err := sql.GetAllCategory()
	if err != nil {
		log.Printf("Get all category failed: %s\n", err)
		return &models.WritingHtmlResponse{}, err
	}
	var writingResponse = &models.WritingHtmlResponse{
		Viewer:       config.GetConfig().Viewer,
		SystemConfig: config.GetConfig().System,
		CategoryList: categoryList,
	}
	return writingResponse, nil
}

// WritingHtmlResponse 编辑界面响应
func (*HTMLApi) WritingHtmlResponse(w http.ResponseWriter, r *http.Request) {
	writingTemplate := common.GetHTMLTemplateCtl().Writing

	writingResponse, err := GetWritingResponseData()
	if err != nil {
		writingTemplate.WriteError(w, errors.New("系统内部错误，请联系管理员！"))
		return
	}
	if err := writingTemplate.Execute(w, writingResponse); err != nil {
		log.Println("writing页面解析报错: ", err)
	}
}
