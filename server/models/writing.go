package models

import "myBlogWeb/config"

type WritingHtmlResponse struct {
	config.Viewer
	config.SystemConfig
	CategoryList []Category `json:"categoryList"`
}
