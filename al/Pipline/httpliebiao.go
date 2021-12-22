package Pipline

import (
	"fmt"
	"github.com/gogf/gf/util/gconv"
	"github.com/tidwall/gjson"
	"go_spider/core/common/com_interfaces"
	"go_spider/core/common/page"
	"go_spider/core/common/page_items"
	"go_spider/core/common/request"
	"go_spider/core/spider"
	"main/Model"
	"main/service"
	"net/http"
	"strings"
)
var ids []int64

type TestPageProcesser struct{}

func (this *TestPageProcesser) Header(h map[string]string) http.Header {
	header := make(http.Header)
	header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/96.0.4664.110 Safari/537.36 Edg/96.0.1054.57")
	if h != nil {
		for k, v := range h {
			header.Add(k, v)
		}
	}
	return header
}

func TestContentPageProcesser() *TestPageProcesser           { return &TestPageProcesser{} }
func (this *TestPageProcesser) Request(req *request.Request) {}

func (this *TestPageProcesser) Process(p *page.Page) {
	if !p.IsSucc() {
		return
	}

	if p.GetStatusCode()!=200 {
		p.SetStatus(true,fmt.Sprintf("status code：%d",p.GetStatusCode()))
		return
	}

	//fmt.Print("http", p.GetStatusCode())
	//fmt.Println("json", p.GetBodyStr())

	if p.GetRequest().Urltag == "first" {
		for i := 2; i <= 4; i++ {
			url := fmt.Sprintf("https://alixiaozhan.net/api/discussions?include=user%%2ClastPostedUser%%2Ctags%%2Ctags.parent%%2CfirstPost&sort=&page%%5Boffset%%5D=%d", i)
			fmt.Println("循环的url", url)
			req := request.NewRequest(url, "json", "", "GET", "", this.Header(nil), nil, nil, this)
			p.AddTargetRequestWithParams(req)

		}
	}
	// lists :=p.GetRequest().GetMeta().(*Testinfo)
	var dicList []*Model.Posts
	for _, node := range gjson.Get(p.GetBodyStr(), "data").Array() {
		fmt.Println("数据id:", node.Get("id").String())
		id := node.Get("id")
		title := node.Get("attributes.title").String()
		htmlcontext := node.Get("attributes.createdAt").String()
		tags := gconv.Int(node.Get("attributes.commentCount").String())
		statusInfo := 1
		var str [] string = [] string{"https://alixiaozhan.net/d/", id.String()}
		s3 := strings.Join(str, "")
		fmt.Println("del_url", s3)
		ids = append(ids, id.Int())
		create := node.Get("attributes.createdAt").String()
		dicList = append(dicList, &Model.Posts{
			Title:       title,
			Request_Id:  id.String(),
			HtmlContext: htmlcontext,
			Tags:        tags,
			Del_url:     s3,
			Status:      statusInfo,
			CreateOn:    create,
		})
		p.AddField("postlist", dicList)
	}

}

// 列表写入数据
type TsPipeline struct{}
func NewTsPipeline() *TsPipeline { return &TsPipeline{} }
func (this *TsPipeline) Process(items *page_items.PageItems, t com_interfaces.Task) {
	allitems := items.GetAll()
	postlist := allitems["postlist"].([]*Model.Posts)
	//循环
	for _, node := range postlist {
		fmt.Println("node的值", node)
		err := service.SavePost(node)
		fmt.Println("SaveErr", err)
	}
}


func (test *Testinfo) Info() {
	spider1 := spider.NewSpider(TestContentPageProcesser(), "").SetThreadnum(2).SetSleepTime("rand", 300, 600).AddPipeline(NewTsPipeline())
	url := fmt.Sprintf("https://alixiaozhan.net/api/discussions?include=user%%2ClastPostedUser%%2Ctags%%2Ctags.parent%%2CfirstPost&sort=")
	fmt.Println("转义", url)
	req := request.NewRequest(url, "json", "first", "GET", "", test.Header(nil), nil, nil, test)
	spider1.AddRequest(req)
	spider1.Run()

	for _,id:=range ids{
		fmt.Println("ids", id)
		spiders := spider.NewSpider(NewTestdelProcesser(id), "").SetThreadnum(2).SetSleepTime("rand", 300, 600)//.AddPipeline(())
		url := fmt.Sprintf("https://alixiaozhan.net/d/%d", id)
		fmt.Println("详情的url", url)
		req := request.NewRequest(url, "text", "", "GET", "", nil, nil, nil, id)
		spiders.AddRequest(req)
		spiders.Run()
	}
}
