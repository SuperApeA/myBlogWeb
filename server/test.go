package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func IndexResponse(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var index IndexInfo
	index.Title = "AajBlogWeb"
	index.Desc = "程序员小艾的博客网站"
	jsonStr, _ := json.Marshal(index)
	if _, err := w.Write(jsonStr); err != nil {
		log.Println(err)
	}
	return
}
