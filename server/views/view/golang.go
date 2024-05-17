package view

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

// GetOnePageBySlugPostMore 获取博客信息
func GetOnePageBySlugPostMore(pageNumber int, slug string) ([]models.PostMore, int, int, error) {
	// 总条数
	totalPosts, err := sql.CountAllPostBySlug(slug)
	if err != nil {
		log.Printf("Count all post failed: %s\n", err)
		return nil, 0, 0, err
	}
	// 总页数
	totalPages := (totalPosts-1)/DefaultPageSize + 1
	if totalPages < pageNumber {
		return []models.PostMore{}, totalPosts, totalPages, nil
	}
	postList, err := sql.GetOnePagePostBySlug(pageNumber, DefaultPageSize, slug)
	if err != nil {
		log.Printf("Get one page post failed: %s\n", err)
		return nil, 0, 0, err
	}
	postMoreList, err := ConvertPostToPostMore(postList)
	if err != nil {
		log.Printf("Convert Post to PostMore: %s\n", err)
		return nil, 0, 0, err
	}
	CutPostMoreContext(postMoreList)
	return postMoreList, totalPosts, totalPages, nil
}

// GetSlugIndexResponseData 获取相应slug的博客信息
func GetSlugIndexResponseData(pageNumber int, slug string) (*models.HomeHtmlResponse, error) {
	// 类别信息
	categoryList, err := sql.GetAllCategory()
	if err != nil {
		return nil, err
	}

	// 博客信息
	postList, totalPosts, totalPages, err := GetOnePageBySlugPostMore(pageNumber, slug)
	if err != nil {
		return nil, err
	}
	var pages []int
	for i := 0; i < totalPages; i++ {
		pages = append(pages, i+1)
	}
	var hr = &models.HomeHtmlResponse{
		Viewer:       config.GetConfig().Viewer,
		CategoryList: categoryList,
		Posts:        postList,
		Total:        totalPosts,
		Page:         pageNumber,
		Pages:        pages,
		PageEnd:      pageNumber < totalPages,
	}
	return hr, nil
}

// GolangHtmlResponse golang界面响应
func (*HTMLApi) GolangHtmlResponse(w http.ResponseWriter, r *http.Request) {
	indexTemplate := common.GetHTMLTemplateCtl().Index

	// 获取页数
	err := r.ParseForm()
	if err != nil {
		indexTemplate.WriteError(w, errors.New("系统内部错误，请联系管理员！"))
		return
	}
	page := r.Form.Get("page")
	pageNumber := 1
	if page != "" {
		pageNumber, _ = strconv.Atoi(page)
	}
	// 获取slug
	slug := strings.TrimLeft(r.URL.Path, "/")
	hr, err := GetSlugIndexResponseData(pageNumber, slug)
	if err != nil {
		indexTemplate.WriteError(w, errors.New("系统内部错误，请联系管理员！"))
		return
	}

	if err := indexTemplate.Execute(w, hr); err != nil {
		log.Println("index返回前端报错: ", err)
	}
}
