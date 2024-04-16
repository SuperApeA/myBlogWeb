package models

type User struct {
	UserID   int    `json:"user_id"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
	CreateAt string `json:"createAt"`
	UpdateAt string `json:"updateAt"`
}
