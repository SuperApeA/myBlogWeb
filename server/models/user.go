package models

type User struct {
	UserID   int    `json:"UserID"`
	UserName string `json:"userName"`
	Password string `json:"password"`
	Avatar   string `json:"avatar"`
	CreateAt string `json:"createAt"`
	UpdateAt string `json:"updateAt"`
}

type UserInfo struct {
	UserID   int    `json:"UserID"`
	UserName string `json:"userName"`
	Avatar   string `json:"avatar"`
}
