package models

type Category struct {
	Cid      int    `json:"cid"`
	Name     string `json:"name"`
	CreateAt string `json:"createAt"`
	UpdateAt string `json:"updateAt"`
}

type CategoryResponse struct {
	HomeResponse
	CategoryName string `json:"categoryName"`
}
