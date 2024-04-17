package sql

import (
	"fmt"
	"log"
	"strings"

	"myBlogWeb/server/models"
	"myBlogWeb/server/utils"
)

func GetUserByNameAndPasswd(userName string, passwd string) (models.User, error) {
	sqlStr := fmt.Sprint("select * from blog_user where user_name = ? and passwd = ? limit 1;")
	log.Println(sqlStr)
	row := DB.QueryRow(sqlStr, userName, passwd)
	if err := row.Err(); err != nil {
		log.Printf("Query user data by name error: %s\n", err)
		return models.User{}, err
	}
	var user models.User
	err := row.Scan(
		&user.UserID,
		&user.UserName,
		&user.Password,
		&user.Avatar,
		&user.CreateAt,
		&user.UpdateAt,
	)
	if err != nil {
		log.Printf("Scan user data by name error: %s\n", err)
		return models.User{}, err
	}
	return user, nil
}

func GetUserByIds(userIds []int) ([]models.User, error) {
	userIds = utils.RemoveDuplicates(userIds)
	placeholders := make([]string, len(userIds))
	args := make([]interface{}, len(userIds))
	for i, uid := range userIds {
		placeholders[i] = "?"
		args[i] = uid
	}
	sqlStr := fmt.Sprintf("select * from blog_user where uid in (%s);", strings.Join(placeholders, ","))
	rows, err := DB.Query(sqlStr, args...)
	if err != nil {
		log.Printf("Query user data error: %s\n", err)
		return nil, err
	}
	if err := rows.Err(); err != nil {
		log.Printf("Select user data error: %s\n", err)
		return nil, err
	}
	var userList []models.User
	for rows.Next() {
		var user models.User
		err := rows.Scan(
			&user.UserID,
			&user.UserName,
			&user.Password,
			&user.Avatar,
			&user.CreateAt,
			&user.UpdateAt,
		)
		if err != nil {
			log.Printf("Scan user data error: %s\n", err)
			return nil, err
		}
		userList = append(userList, user)
	}
	return userList, nil
}
