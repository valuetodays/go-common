package utils
//
//import (
//	"fmt"
//	"github.com/valuetodays/go-common/rest"
//	"testing"
//)
//
//func TestDoPostJson(t *testing.T) {
//	const url = "http://api.valuetodays.cn/myLink/feign/listTreeByUserId.do"
//	const req uint64 = 1
//	myLinkDataR := rest.MyLinkDataR{}
//	got := DoPostJson(url, req, &myLinkDataR)
//	fmt.Println("got[", got, "]")
//	fmt.Println("myLinkDataR", myLinkDataR)
//	list := myLinkDataR.Data.ItemList
//	PrintItemList(list)
//	//fmt.Printf("r=%p\n\n", &myLinkDataR)
//	//fmt.Printf("r2=%p\n\n", &myLinkDataR2)
//}
//
//func PrintItemList(itemList []rest.ItemList) {
//	for _, item := range itemList {
//		fmt.Println("-> " + item.Title + item.Url)
//		children := item.Children
//		if len(children) > 0 {
//			PrintItemList(children)
//		}
//	}
//}
