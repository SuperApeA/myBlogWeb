package models

import "myBlogWeb/config"

type PigeonholeResponse struct {
	config.Viewer
	config.SystemConfig
	CategoryList []Category
	PostMap      map[string][]Post
}
