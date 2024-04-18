package sql

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"myBlogWeb/server/models"
	"myBlogWeb/server/utils"
	"strings"
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

func CountAllPostByCategoryIDs(categoryIds []int) (int, error) {
	categoryIds = utils.RemoveDuplicates(categoryIds)
	placeholders := make([]string, len(categoryIds))
	args := make([]interface{}, len(categoryIds))
	for i, uid := range categoryIds {
		placeholders[i] = "?"
		args[i] = uid
	}
	sqlStr := fmt.Sprintf("select count(*) from blog_post where category_id in (%s);", strings.Join(placeholders, ","))
	row := DB.QueryRow(sqlStr, args...)
	total := 0
	if err := row.Scan(&total); err != nil {
		log.Printf("Count all post data by category id error: %s\n", err)
		return 0, err
	}
	return total, nil
}

// GetOnePagePost
func GetOnePagePost(pageNumber int, pageSize int) ([]models.Post, error) {
	// 计算 OFFSET 值，pageNumber从1开始算
	offset := (pageNumber - 1) * pageSize
	sqlStr := fmt.Sprintf("select * from blog_post limit %v offset %v;", pageSize, offset)
	rows, err := DB.Query(sqlStr)
	if err != nil {
		log.Printf("Query all post data error: %s\n", err)
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
		if errors.Is(err, sql.ErrNoRows) {
			return []models.Post{}, nil
		}
		if err != nil {
			log.Printf("Scan post data error: %s\n", err)
			return nil, err
		}
		postList = append(postList, post)
	}
	return postList, nil
}

// GetOnePagePostByCategoryIDs
func GetOnePagePostByCategoryIDs(categoryIds []int, pageNumber int, pageSize int) ([]models.Post, error) {
	categoryIds = utils.RemoveDuplicates(categoryIds)
	placeholders := make([]string, len(categoryIds))
	args := make([]interface{}, len(categoryIds))
	for i, uid := range categoryIds {
		placeholders[i] = "?"
		args[i] = uid
	}
	// 计算 OFFSET 值，pageNumber从1开始算
	offset := (pageNumber - 1) * pageSize
	sqlStr := fmt.Sprintf("select * from blog_post where category_id in (%s) limit %v offset %v;", strings.Join(placeholders, ","), pageSize, offset)
	rows, err := DB.Query(sqlStr, args...)
	if err != nil {
		log.Printf("Query all post data error: %s\n", err)
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
		if errors.Is(err, sql.ErrNoRows) {
			return []models.Post{}, nil
		}
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
		log.Println("Query all pos data error")
		return nil, err
	}
	var postList []models.Post
	for rows.Next() {
		var post models.Post
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
		if errors.Is(err, sql.ErrNoRows) {
			return []models.Post{}, nil
		}
		if err != nil {
			log.Println("Scan post data error")
			return nil, err
		}
		postList = append(postList, post)
	}
	return postList, nil
}

func GetPostByID(postID int) (models.Post, error) {
	sqlStr := fmt.Sprint("select * from blog_post where pid = ?;")
	row := DB.QueryRow(sqlStr, postID)
	var post models.Post
	err := row.Scan(
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
	if errors.Is(err, sql.ErrNoRows) {
		return models.Post{}, nil
	}
	if err != nil {
		log.Printf("Scan post data error: %v\n", err)
		return models.Post{}, err
	}
	return post, nil
}

func CreatePost(post *models.Post) error {
	sqlStr := fmt.Sprint("insert into blog_post " +
		"(title, content, markdown, category_id, user_id, view_count, type, slug, create_at, update_at) " +
		"values(?,?,?,?,?,?,?,?,?,?);")
	res, err := DB.Exec(sqlStr,
		post.Title,
		post.Content,
		post.Markdown,
		post.CategoryId,
		post.UserId,
		post.ViewCount,
		post.Type,
		post.Slug,
		post.CreateAt,
		post.UpdateAt)
	if err != nil {
		log.Printf("Insert post error: %v\n", err)
		return err
	}
	pid, _ := res.LastInsertId()
	post.Pid = int(pid)
	return nil
}
