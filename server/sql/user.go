package sql

import (
	"fmt"
	"log"
	"strings"

	"myBlogWeb/server/models"
	"myBlogWeb/server/utils"
)

func GetUserNameByIds(userIds []int) ([]models.User, error) {
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
