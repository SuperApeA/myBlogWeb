package api

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
)

var API = &Api{}

type Api struct{}

type Message struct {
	Code  int         `json:"code"`
	Data  interface{} `json:"data"`
	Error string      `json:"error"`
}

func SuccessRes(data interface{}) []byte {
	var message Message
	message.Data = data
	message.Code = 200
	message.Error = ""
	response, _ := json.Marshal(message)
	return response
}

func ErrorRes(err error) []byte {
	var message Message
	message.Data = nil
	message.Code = 400
	message.Error = err.Error()
	response, _ := json.Marshal(message)
	return response
}

func GetRequestJsonParam(r *http.Request) (map[string]interface{}, error) {
	param := make(map[string]interface{})
	var body bytes.Buffer
	if _, err := io.Copy(&body, r.Body); err != nil {
		log.Printf("Copy request body failed: %v\n", err)
		return nil, err
	}
	// 读取请求体后，确保关闭它
	_ = r.Body.Close()
	_ = json.Unmarshal(body.Bytes(), &param)
	return param, nil
}

//func GetRequestJsonParam(r *http.Request) (map[string]interface{}, error) {
//	param := make(map[string]interface{})
//	body, err := io.ReadAll(r.Body)
//	if err != nil {
//		log.Printf("Get request body failed: %v\n", err)
//		return nil, err
//	}
//	// 读取请求体后，确保关闭它
//	_ = r.Body.Close()
//	_ = json.Unmarshal(body, &param)
//	return param, nil
//}
