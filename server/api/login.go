package api

import (
	"errors"
	"log"
	"myBlogWeb/server/models"
	"myBlogWeb/server/sql"
	"net/http"
)

func (*Api) LoginApiResponse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	param, err := GetRequestJsonParam(r)
	if err != nil {
		log.Printf("Get login param failed, err: %v\n", err)
		_, _ = w.Write(ErrorRes(errors.New("请求参数非法！")))
		return
	}
	userName := param["username"].(string)
	passwd := param["passwd"].(string)

	user, err := sql.GetUserByNameAndPasswd(userName, passwd)
	if err != nil || user.UserName != userName {
		_, _ = w.Write(ErrorRes(errors.New("用户不存在或密码错误！")))
		return
	}
	var res models.LoginApiResponse
	var _, _ = w.Write(SuccessRes(res))
	return
}
