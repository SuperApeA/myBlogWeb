package sql

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"strings"

	"myBlogWeb/server/models"
	"myBlogWeb/server/utils"
)

func GetAllCategory() ([]models.Category, error) {
	sqlStr := fmt.Sprint("select * from blog_category;")
	rows, err := DB.Query(sqlStr)
	if err != nil {
		log.Printf("Query all category data error: %s\n", err)
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
		if errors.Is(err, sql.ErrNoRows) {
			return []models.Category{}, nil
		}
		if err != nil {
			log.Printf("Scan category data error: %s\n", err)
			return nil, err
		}
		categoryList = append(categoryList, category)
	}
	return categoryList, nil
}

func GetCategoryByIds(categoryIds []int) ([]models.Category, error) {
	categoryIds = utils.RemoveDuplicates(categoryIds)
	placeholders := make([]string, len(categoryIds))
	args := make([]interface{}, len(categoryIds))
	for i, uid := range categoryIds {
		placeholders[i] = "?"
		args[i] = uid
	}
	sqlStr := fmt.Sprintf("select * from blog_category where cid in (%s);", strings.Join(placeholders, ","))
	rows, err := DB.Query(sqlStr, args...)
	if err != nil {
		log.Printf("Query category data error: %s\n", err)
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
		if errors.Is(err, sql.ErrNoRows) {
			return []models.Category{}, nil
		}
		if err != nil {
			log.Printf("Scan category data error: %s\n", err)
			return nil, err
		}
		categoryList = append(categoryList, category)
	}
	return categoryList, nil
}
