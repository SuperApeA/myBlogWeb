package models

import "myBlogWeb/config"

type LoginHtmlResponse struct {
	config.Viewer
}

type LoginApiResponse struct {
	Token    string   `json:"token"`
	UserInfo UserInfo `json:"userInfo"`
}
