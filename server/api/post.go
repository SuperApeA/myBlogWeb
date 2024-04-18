package api

import (
	"errors"
	"log"
	"myBlogWeb/server/models"
	"myBlogWeb/server/sql"
	"myBlogWeb/server/utils"
	"net/http"
	"strconv"
	"strings"
	"time"
)

/*
请求入参格式：
categoryId : "1"
content : "<p>test123</p>\n"
markdown : "test123\n"
slug : "123"
title : "test123"
type : 0
*/
func CreatPostResponse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	token := r.Header.Get("Authorization")
	_, claims, err := utils.ParseToken(token)
	if err != nil {
		_, _ = w.Write(ErrorRes(errors.New("用户未登录或登录已过期！")))
		return
	}
	param, err := GetRequestJsonParam(r)
	if err != nil {
		log.Printf("Get post param failed, err: %v\n", err)
		_, _ = w.Write(ErrorRes(errors.New("请求参数非法！")))
		return
	}
	categoryId := param["categoryId"].(string)
	content := param["content"].(string)
	markdown := param["markdown"].(string)
	slug := param["slug"].(string)
	title := param["title"].(string)
	type_ := param["type"].(float64)

	cid, _ := strconv.Atoi(categoryId)
	post := models.Post{
		Title:      title,
		Slug:       slug,
		Content:    content,
		Markdown:   markdown,
		CategoryId: cid,
		Type:       int(type_),
		UserId:     claims.Uid,
		ViewCount:  0,
		CreateAt:   time.Now(),
		UpdateAt:   time.Now(),
	}
	err = sql.CreatePost(&post)
	if err != nil {
		log.Printf("Create post failed, err: %v\n", err)
		_, _ = w.Write(ErrorRes(errors.New("系统出错，发布文章失败，请联系管理员！")))
		return
	}
	_, _ = w.Write(SuccessRes(models.PostApiResponse{Pid: post.Pid}))
}

func GetPostResponse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	token := r.Header.Get("Authorization")
	_, _, err := utils.ParseToken(token)
	if err != nil {
		_, _ = w.Write(ErrorRes(errors.New("用户未登录或登录已过期！")))
		return
	}
	// 获取postID
	urlPath := r.URL.Path
	postIDStr := strings.TrimPrefix(urlPath, "/api/v1/post/")
	postID, err := strconv.Atoi(postIDStr)
	if err != nil {
		_, _ = w.Write(ErrorRes(errors.New("category的id非法")))
		return
	}

	post, err := sql.GetPostByID(postID)
	if err != nil {
		log.Printf("Get post failed, err: %v\n", err)
		_, _ = w.Write(ErrorRes(errors.New("系统出错，查询文章失败，请联系管理员！")))
		return
	}
	res := models.PostApiResponse{
		Pid:        post.Pid,
		Title:      post.Title,
		CategoryId: post.CategoryId,
		Type:       post.Type,
		Slug:       post.Slug,
		Markdown:   post.Markdown,
	}
	_, _ = w.Write(SuccessRes(res))
}

func (*Api) PostApiResponse(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		CreatPostResponse(w, r)
	case http.MethodGet:
		GetPostResponse(w, r)
	}
	return
}
