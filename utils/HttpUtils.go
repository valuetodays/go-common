package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func DoPostJson(url string, req any, resp any) string {
	reqBytes, err := json.Marshal(req) // 把请求结构体解析为json
	if err != nil {
		fmt.Println("marshal failed. the error info: ", err)
	}

	// 调用rest接口
	post, err := http.Post(url, "application/json", bytes.NewBuffer(reqBytes))
	fmt.Println("respString=", post)
	if post == nil {
		return ""
	}
	defer post.Body.Close()

	postBody, err := ioutil.ReadAll(post.Body)
	if err != nil {
		fmt.Errorf("ReadAll failed, url: %s, reqBody: %s, err: %v", url, postBody, err)
	}

	bodyAsString := string(postBody)
	json.Unmarshal(postBody, &resp) // 解析请求到结构体

	return bodyAsString
}
