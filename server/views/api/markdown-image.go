package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"myBlogWeb/config"
	blogerr "myBlogWeb/error"
	"myBlogWeb/server/models"
	"myBlogWeb/server/views/common"
)

var imagePath string

const (
	imageLoadPath = "/viewsrc/markdown/image/"

	rspUrlPrefix = "/markdown/image/"
)

func init() {
	imagePath = filepath.Join(config.AppLocalPath, imageLoadPath)
}

func (*Api) PostUploadFileApiResponse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// 校验token
	var ok bool
	if _, ok = common.CheckIsLogin(r); !ok {
		_, _ = w.Write(ErrorRes(errors.New(blogerr.LoginOut)))
		return
	}

	// 限制上传文件的大小，例如10MB
	r.ParseMultipartForm(10 << 20)
	file, handler, err := r.FormFile("editormd-image-file")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	defer file.Close()

	//fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	//fmt.Printf("File Size: %+v\n", handler.Size)
	//fmt.Printf("MIME Header: %+v\n", handler.Header)

	// 创建文件
	imageName := handler.Filename[:strings.LastIndex(handler.Filename, ".")] +
		fmt.Sprintf("-time-%s", strconv.FormatInt(time.Now().Unix(), 10)) +
		handler.Filename[strings.LastIndex(handler.Filename, "."):]
	dst, err := os.Create(filepath.Join(imagePath, imageName))
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer dst.Close()

	// 将上传的文件内容拷贝到新创建的文件中
	if _, err := io.Copy(dst, file); err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var res models.FileApiResponse
	res.Success = 1
	res.Url = rspUrlPrefix + filepath.Base(dst.Name())
	resByte, _ := json.Marshal(res)
	var _, _ = w.Write(resByte)
	return
}
