package sql

import (
	"fmt"
	"log"

	"myBlogWeb/server/models"
)

func CountAllPost() (int, error) {
	sqlStr := fmt.Sprint("select count(*) from blog_post;")
	row := DB.QueryRow(sqlStr)
	total := 0
	if err := row.Scan(&total); err != nil {
		log.Printf("Count all post data error: %s\n", err)
		return 0, err
	}
	return total, nil
}

// GetOnePagePost
func GetOnePagePost(pageNumber int, pageSize int) ([]models.Post, error) {
	// 计算 OFFSET 值，pageNumber从1开始算
	offset := (pageNumber - 1) * pageSize
	sqlStr := fmt.Sprint("select * from blog_post limit ? offset ?;")
	rows, err := DB.Query(sqlStr, pageSize, offset)
	if err != nil {
		log.Printf("Select all post data error: %s\n", err)
		return nil, err
	}
	var postList []models.Post
	for rows.Next() {
		var post models.Post
		// Scan读取的列位置需要和变量名保持一致
		err := rows.Scan(
			&post.Pid,
			&post.Title,
			&post.Content,
			&post.Markdown,
			&post.CategoryId,
			&post.UserId,
			&post.ViewCount,
			&post.Type,
			&post.Slug,
			&post.CreateAt,
			&post.UpdateAt,
		)
		if err != nil {
			log.Printf("Scan post data error: %s\n", err)
			return nil, err
		}
		postList = append(postList, post)
	}
	return postList, nil
}

func GetAllPost() ([]models.Post, error) {
	sqlStr := fmt.Sprint("select * from blog_post;")
	rows, err := DB.Query(sqlStr)
	if err != nil {
		log.Println("Select all pos data error")
		return nil, err
	}
	var postList []models.Post
	for rows.Next() {
		var post models.Post
		err := rows.Scan(
			&post.Pid,
			&post.Title,
			&post.Slug,
			&post.Content,
			&post.Markdown,
			&post.CategoryId,
			&post.UserId,
			&post.ViewCount,
			&post.Type,
			&post.CreateAt,
			&post.UpdateAt,
		)
		if err != nil {
			log.Println("Scan post data error")
			return nil, err
		}
		postList = append(postList, post)
	}
	return postList, nil
}
