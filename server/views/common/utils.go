package common

import (
	"net/http"

	"myBlogWeb/server/utils"
)

func CheckIsLogin(r *http.Request) (int, bool) {
	token := r.Header.Get("Authorization")
	_, claims, err := utils.ParseToken(token)
	if err != nil {
		return -1, false
	}
	return claims.Uid, true
}
