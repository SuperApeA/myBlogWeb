package sql

import (
	"fmt"
	"log"

	"myBlogWeb/server/models"
)

func GetAllCategory() ([]models.Category, error) {
	sqlStr := fmt.Sprint("select * from blog_category;")
	rows, err := DB.Query(sqlStr)
	if err != nil {
		log.Println("Query all category data error")
		return nil, err
	}
	var categoryList []models.Category
	for rows.Next() {
		var category models.Category
		err := rows.Scan(
			&category.Cid,
			&category.Name,
			&category.CreateAt,
			&category.UpdateAt,
		)
		if err != nil {
			log.Println("Scan category data error")
			return nil, err
		}
		categoryList = append(categoryList, category)
	}
	return categoryList, nil
}
