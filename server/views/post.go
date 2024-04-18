package views

import (
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"

	"myBlogWeb/config"
	"myBlogWeb/server/models"
	"myBlogWeb/server/sql"
	"myBlogWeb/server/views/common"
)

func GetPostDetailResponseData(postID int) (*models.PostHtmlResponse, error) {
	// 博客详情
	post, err := sql.GetPostByID(postID)
	if err != nil {
		return nil, err
	}
	postMoreList, err := ConvertPostToPostMore([]models.Post{post})
	if err != nil {
		log.Printf("Convert Post to PostMore: %s\n", err)
		return nil, err
	}
	var postResponse = &models.PostHtmlResponse{
		Viewer:       config.GetConfig().Viewer,
		SystemConfig: config.GetConfig().System,
		Article:      postMoreList[0],
	}
	return postResponse, nil
}

// PostDetailHtmlResponse post详情响应
func (*HTMLApi) PostDetailHtmlResponse(w http.ResponseWriter, r *http.Request) {
	postDetailTemplate := common.GetHTMLTemplateCtl().Detail

	// 获取postID
	urlPath := r.URL.Path
	postIDStr := strings.TrimPrefix(urlPath, "/post/")
	postID, err := strconv.Atoi(postIDStr)
	if err != nil {
		postDetailTemplate.WriteError(w, errors.New("post的id非法"))
		return
	}

	postResponse, err := GetPostDetailResponseData(postID)
	if err != nil {
		postDetailTemplate.WriteError(w, errors.New("系统内部错误，请联系管理员！"))
		return
	}
	if err := postDetailTemplate.Execute(w, postResponse); err != nil {
		log.Println("post返回前端报错: ", err)
	}
}
