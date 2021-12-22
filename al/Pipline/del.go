package Pipline

import (
	"fmt"
	"go_spider/core/common/page"
	"go_spider/core/common/request"
	"net/http"
)

//详情页
type TDelProcesser struct {
	Id int64
}

func (this *TDelProcesser) Header(h map[string]string) http.Header {
	header := make(http.Header)
	header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36 Edg/96.0.1054.57")
	if h != nil {
		for k, v := range h {
			header.Add(k, v)
		}
	}
	return header
}
func NewTestdelProcesser(id int64) *TDelProcesser {
	return &TDelProcesser{
		id,
	}
}
func (this *TDelProcesser) Request(req *request.Request) {}
func (this *TDelProcesser) Process(p *page.Page) {
	if !p.IsSucc() {
		return
	}
	fmt.Println("delcode",p.GetStatusCode())
	if p.GetStatusCode() != 200 {
		p.SetStatus(true, fmt.Sprintf("status code：%d", p.GetStatusCode()))
		return
	}
	fmt.Println("delurl", p.GetBodyStr())
}

// 详情数据写入
