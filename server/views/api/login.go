package api

import (
	"errors"
	"log"
	"net/http"

	"myBlogWeb/server/models"
	"myBlogWeb/server/sql"
	"myBlogWeb/server/utils"
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
	passwd = utils.Md5Crypt(passwd)

	user, err := sql.GetUserByNameAndPasswd(userName, passwd)
	if err != nil || user.UserName != userName {
		_, _ = w.Write(ErrorRes(errors.New("用户不存在或密码错误！")))
		return
	}
	var res models.LoginApiResponse
	//生成token  jwt技术进行生成 令牌  A.B.C
	token, err := utils.Award(&user.UserID)
	if err != nil {
		_, _ = w.Write(ErrorRes(errors.New("系统内部错误，请联系管理员！")))
		return
	}
	var userInfo models.UserInfo
	userInfo.UserID = user.UserID
	userInfo.UserName = user.UserName
	userInfo.Avatar = user.Avatar
	res.Token = token
	res.UserInfo = userInfo
	var _, _ = w.Write(SuccessRes(res))
	return
}
