package Model

type Posts struct{
	// Id int //`json:"id"`
    Title string //`json:"title"`
	HtmlContext string //`json:"HtmlCo"`
	Tags int
	Request_Id string `json:"Request_id"`
	Aliyun_Url string `json:"aliyun_url"`
	Del_url string `json:"del_url"`
	Status int
	CreateOn string
}
