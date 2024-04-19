package api

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"myBlogWeb/server/models"
	"net/http"
	"os"
	"path/filepath"

	"myBlogWeb/config"
)

var imagePath string

func init() {
	imagePath = filepath.Join(config.AppLocalPath, "viewsrc/markdown/image/")
}

func (*Api) PostUploadFileApiResponse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

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
	dst, err := os.Create(filepath.Join(imagePath, handler.Filename))
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
	res.Url = dst.Name()
	resByte, _ := json.Marshal(res)
	var _, _ = w.Write(resByte)
	return
}
