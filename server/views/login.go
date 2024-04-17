package views

import (
	"log"
	"myBlogWeb/config"
	"myBlogWeb/server/models"
	"myBlogWeb/server/views/common"
	"net/http"
)

func (*HTMLApi) LoginHtmlResponse(w http.ResponseWriter, r *http.Request) {
	loginTemplate := common.GetHTMLTemplateCtl().Login

	var loginResponse models.LoginHtmlResponse
	loginResponse.Viewer = config.GetConfig().Viewer

	if err := loginTemplate.Execute(w, loginResponse); err != nil {
		log.Println("index返回前端报错: ", err)
	}
}
