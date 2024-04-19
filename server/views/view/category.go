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

// GetOnePagePostMoreByCategoryID 通过类别获取博客信息
func GetOnePagePostMoreByCategoryID(categoryID int, pageNumber int) ([]models.PostMore, int, int, error) {
	// 总条数
	totalPosts, err := sql.CountAllPostByCategoryIDs([]int{categoryID})
	if err != nil {
		log.Printf("Count all post by category id failed: %s\n", err)
		return nil, 0, 0, err
	}
	// 总页数
	totalPages := (totalPosts-1)/DefaultPageSize + 1
	if totalPages < pageNumber {
		return []models.PostMore{}, totalPosts, totalPages, nil
	}
	postList, err := sql.GetOnePagePostByCategoryIDs([]int{categoryID}, pageNumber, DefaultPageSize)
	if err != nil {
		log.Printf("Get one page post by category id failed: %s\n", err)
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

func GetCategoryResponseData(categoryID int, pageNumber int) (*models.CategoryHtmlResponse, error) {
	// 类别信息
	categoryList, err := sql.GetAllCategory()
	if err != nil {
		return nil, err
	}

	// 博客信息
	postList, totalPosts, totalPages, err := GetOnePagePostMoreByCategoryID(categoryID, pageNumber)
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
	var categoryResponse = &models.CategoryHtmlResponse{
		HomeHtmlResponse: *hr,
	}
	if len(postList) > 0 {
		categoryResponse.CategoryName = postList[0].CategoryName
	}
	return categoryResponse, nil
}

// CategoryHtmlResponse category界面响应
func (*HTMLApi) CategoryHtmlResponse(w http.ResponseWriter, r *http.Request) {
	categoryTemplate := common.GetHTMLTemplateCtl().Category

	// 获取categoryID
	urlPath := r.URL.Path
	categoryIDStr := strings.TrimPrefix(urlPath, "/category/")
	categoryID, err := strconv.Atoi(categoryIDStr)
	if err != nil {
		categoryTemplate.WriteError(w, errors.New("category的id非法"))
		return
	}

	// 获取页数
	err = r.ParseForm()
	if err != nil {
		categoryTemplate.WriteError(w, errors.New("系统内部错误，请联系管理员！"))
		return
	}
	page := r.Form.Get("page")
	pageNumber := 1
	if page != "" {
		pageNumber, _ = strconv.Atoi(page)
	}
	categoryResponse, err := GetCategoryResponseData(categoryID, pageNumber)
	if err != nil {
		categoryTemplate.WriteError(w, errors.New("系统内部错误，请联系管理员！"))
		return
	}
	if err := categoryTemplate.Execute(w, categoryResponse); err != nil {
		log.Println("category返回前端报错: ", err)
	}
}
