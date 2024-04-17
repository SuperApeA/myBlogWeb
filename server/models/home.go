package models

import "myBlogWeb/config"

type HomeHtmlResponse struct {
	config.Viewer
	CategoryList []Category
	Posts        []PostMore
	Total        int
	Page         int
	Pages        []int
	PageEnd      bool
}
