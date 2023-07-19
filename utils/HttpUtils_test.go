package utils

import (
	"fmt"
	"github.com/valuetodays/go-common/rest"
	"testing"
)


type IndexResp struct {
	rest.R
	Data IndexData `json:"data"`
}

type IndexData struct {
	TimestampMs uint64 `json:"timestampMs"`
	PoweredBy string `json:"poweredBy"`
	IntValue uint32 `json:"intValue"`
}

func TestDoPostJson(t *testing.T) {
	fmt.Println("please run simple-http-server in docs/simple-http-server.7z first!!!!")

	const url = "http://localhost:18080"
	const req uint64 = 1
	indexResp := IndexResp{}
	got := DoPostJson(url, req, &indexResp)
	fmt.Println("got[", got, "]")
	fmt.Println("indexResp", indexResp)
	indexData := indexResp.Data
	fmt.Printf("indexData-pointer=%p\n", &indexData)
	fmt.Printf("indexData=%v\n", &indexData)
}
