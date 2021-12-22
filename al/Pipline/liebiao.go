package Pipline

import (
	"fmt"
	"github.com/tidwall/gjson"
	"net/http"
)

type Testinfo struct {
}

func NewTestInfo() *Testinfo {
	return &Testinfo{}
}
func (test *Testinfo) Header(h map[string]string) http.Header {
	header := make(http.Header)
	header.Set("User-Agent","Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36 Edg/96.0.1054.57")
	if h != nil {
		for k, v := range h {
			header.Add(k, v)
		}
	}
	return header
}

func (test *Testinfo) GetStatus(result gjson.Result) string {
	fmt.Println("result", result.Get("first").String())
	return " null"
}

func (test *Testinfo) GetTest(result gjson.Result) string {
	fmt.Println("TestInfo", result)
	return "null2"
}
