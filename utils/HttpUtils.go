package utils

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"strings"
)


// 简单封装一个请求api的方法
func DoCallApi(method string, ipAndPort string, path string, body io.Reader) (string, error) {
	// 1.如果没有http开头就给它加一个
	if !strings.HasPrefix(ipAndPort, "http://") && !strings.HasPrefix(ipAndPort, "https://") {
		ipAndPort = "http://" + ipAndPort
	}
	// 2. 新建一个request
	req, _ := http.NewRequest(method, ipAndPort + path, body)

	// 3. 新建httpclient，并且传入request
	client := http.DefaultClient
	res, err := client.Do(req)
	if err != nil {
		return "", err
	}

	defer res.Body.Close()

	// 4. 获取请求结果
	buff, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(buff), nil
}

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
