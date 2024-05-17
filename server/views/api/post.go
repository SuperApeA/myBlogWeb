package api

import (
	"errors"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	blogerr "myBlogWeb/error"
	"myBlogWeb/server/models"
	"myBlogWeb/server/sql"
	"myBlogWeb/server/views/common"
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

	var uid int
	var ok bool
	if uid, ok = common.CheckIsLogin(r); !ok {
		_, _ = w.Write(ErrorRes(errors.New(blogerr.LoginOut)))
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
		UserId:     uid,
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

func PutPostResponse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var ok bool
	if _, ok = common.CheckIsLogin(r); !ok {
		_, _ = w.Write(ErrorRes(errors.New(blogerr.LoginOut)))
		return
	}

	param, err := GetRequestJsonParam(r)
	if err != nil {
		log.Printf("Update post param failed, err: %v\n", err)
		_, _ = w.Write(ErrorRes(errors.New("请求参数非法！")))
		return
	}
	postId := param["pid"].(float64)
	categoryId := param["categoryId"].(float64)
	content := param["content"].(string)
	markdown := param["markdown"].(string)
	slug := param["slug"].(string)
	title := param["title"].(string)
	type_ := param["type"].(float64)

	post, err := sql.GetPostByID(int(postId))
	if err != nil {
		log.Printf("Get post by id  failed, err: %v\n", err)
		_, _ = w.Write(ErrorRes(errors.New("系统出错，发布文章失败，请联系管理员！")))
		return
	}
	post.Title = title
	post.Slug = slug
	post.Content = content
	post.Markdown = markdown
	post.CategoryId = int(categoryId)
	post.Type = int(type_)
	post.UpdateAt = time.Now()
	err = sql.UpdatePost(&post)
	if err != nil {
		log.Printf("Update post failed, err: %v\n", err)
		_, _ = w.Write(ErrorRes(errors.New("系统出错，发布文章失败，请联系管理员！")))
		return
	}
	_, _ = w.Write(SuccessRes(models.PostApiResponse{Pid: post.Pid}))
}

func GetPostResponse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var ok bool
	if _, ok = common.CheckIsLogin(r); !ok {
		_, _ = w.Write(ErrorRes(errors.New(blogerr.LoginOut)))
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
		Uid:        post.UserId,
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
	case http.MethodPut:
		PutPostResponse(w, r)
	case http.MethodGet:
		GetPostResponse(w, r)
	}
	return
}

func (*Api) PostSearchApiResponse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	err := r.ParseForm()
	if err != nil {
		_, _ = w.Write(ErrorRes(errors.New("系统内部错误，请联系管理员！")))
		return
	}
	searchKey := r.Form.Get("val")
	var res []models.SearchResp
	posts, err := sql.GetPostByCondition(searchKey)
	if err != nil {
		log.Printf("Get post failed, err: %v\n", err)
		_, _ = w.Write(ErrorRes(errors.New("系统出错，查询文章失败，请联系管理员！")))
		return
	}
	for _, post := range posts {
		res = append(res, models.SearchResp{
			Pid:   post.Pid,
			Title: post.Title,
		})
	}
	_, _ = w.Write(SuccessRes(res))
}
