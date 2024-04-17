package views

import (
	"errors"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"myBlogWeb/config"
	"myBlogWeb/server/models"
	"myBlogWeb/server/sql"
	"myBlogWeb/server/utils"
	"myBlogWeb/server/views/common"
)

const DefaultPageSize = 10

func ConvertPostToPostMore(p []models.Post) ([]models.PostMore, error) {
	var postMoreList []models.PostMore
	var userIds, categoryIds []int
	for _, post := range p {
		userIds = append(userIds, post.UserId)
		categoryIds = append(categoryIds, post.CategoryId)
	}
	userList, err := sql.GetUserByIds(userIds)
	userMap := make(map[int]models.User)
	for _, user := range userList {
		userMap[user.UserID] = user
	}
	categoryList, err := sql.GetCategoryByIds(categoryIds)
	categoryMap := make(map[int]models.Category)
	for _, category := range categoryList {
		categoryMap[category.Cid] = category
	}
	if err != nil {
		log.Printf("Get user name by id failed: %s", err)
		return nil, err
	}
	for _, post := range p {
		var postMore models.PostMore
		postMore.Pid = post.Pid
		postMore.Title = post.Title
		postMore.Slug = post.Slug
		content := []rune(post.Content)
		if len(content) > 100 {
			content = content[:100]
		}
		postMore.Content = template.HTML(content)
		postMore.CategoryId = post.CategoryId
		postMore.ViewCount = post.ViewCount
		postMore.Type = post.Type
		postMore.CreateAt = utils.FormatTime(post.CreateAt, "")
		postMore.UpdateAt = utils.FormatTime(post.UpdateAt, "")
		postMore.UserName = userMap[post.UserId].UserName
		postMore.CategoryName = categoryMap[post.CategoryId].Name
		postMoreList = append(postMoreList, postMore)
	}
	return postMoreList, nil
}

// GetOnePagePostMore 获取博客信息
func GetOnePagePostMore(pageNumber int) ([]models.PostMore, int, int, error) {
	// 总条数
	totalPosts, err := sql.CountAllPost()
	if err != nil {
		log.Printf("Count all post failed: %s\n", err)
		return nil, 0, 0, err
	}
	// 总页数
	totalPages := (totalPosts-1)/DefaultPageSize + 1
	if totalPages < pageNumber {
		return []models.PostMore{}, totalPosts, totalPages, nil
	}
	postList, err := sql.GetOnePagePost(pageNumber, DefaultPageSize)
	if err != nil {
		log.Printf("Get one page post failed: %s\n", err)
		return nil, 0, 0, err
	}
	postMoreList, err := ConvertPostToPostMore(postList)
	if err != nil {
		log.Printf("Convert Post to PostMore: %s\n", err)
		return nil, 0, 0, err
	}
	return postMoreList, totalPosts, totalPages, nil
}

// GetIndexResponseData 获取首页信息
func GetIndexResponseData(pageNumber int) (*models.HomeResponse, error) {
	// 类别信息
	categoryList, err := sql.GetAllCategory()
	if err != nil {
		return nil, err
	}

	// 博客信息
	postList, totalPosts, totalPages, err := GetOnePagePostMore(pageNumber)
	if err != nil {
		return nil, err
	}
	var pages []int
	for i := 0; i < totalPages; i++ {
		pages = append(pages, i+1)
	}
	var hr = &models.HomeResponse{
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

// IndexHtmlResponse index界面响应，home页
func (*HTMLApi) IndexHtmlResponse(w http.ResponseWriter, r *http.Request) {
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
	hr, err := GetIndexResponseData(pageNumber)
	if err != nil {
		indexTemplate.WriteError(w, errors.New("系统内部错误，请联系管理员！"))
		return
	}

	if err := indexTemplate.Execute(w, hr); err != nil {
		log.Println("index返回前端报错: ", err)
	}
}
