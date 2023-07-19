package cloud

import (
	"crypto/rand"
	"github.com/valuetodays/go-common/utils"
	"go-micro.dev/v4/registry"
	"io"
	"math/big"
	"strconv"
)

func CallApiByService(registry registry.Registry, serviceName string, method string, path string, body io.Reader) (string, error) {
	services, err := registry.GetService(serviceName)
	if nil != err {
		return "could not get services", err
	}

	matchedService := services[RandomWith(len(services))]
	matchedNode := matchedService.Nodes[RandomWith(len(matchedService.Nodes))]

	resp, err := utils.DoCallApi(method, matchedNode.Address, path, body)
	return resp, err
}

// random value from [0, endExclude)
func RandomWith(endExclude int) int  {
	randomBigInt, _ := rand.Int(rand.Reader, big.NewInt(int64(endExclude)))
	// string转int
	randomInt, _ := strconv.Atoi(randomBigInt.String())
	return randomInt
}

