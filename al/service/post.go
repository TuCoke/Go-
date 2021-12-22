package service

import (
	"database/sql"
	"fmt"
	"github.com/gogf/gf/frame/g"
	"main/Model"
	"main/util"
	"strings"
	"time"
)

func SavePost(post *Model.Posts) error {
	//fmt.Println("savepost", post)
	post.CreateOn = util.FormatDate(time.Now(), 1)

	// ail := fmt.Sprintf("https://alixiaozhan.net/d/", post.Request_Id)
	var str [] string = [] string{"https://alixiaozhan.net/d/", post.Request_Id}
	s3 := strings.Join(str, "")
	fmt.Println("del_url", s3)

	isExit := g.DB().Table("posts").Fields("del_url").Struct(post.Del_url, "del_url=?", s3)
	fmt.Println("isexit:", isExit)
	if isExit == sql.ErrNoRows{
		Id, err := g.DB().Table("posts").InsertAndGetId(post)
		fmt.Println("ErrInsert", err)
		fmt.Println("返回的id", Id)
	}else if isExit == nil{
		fmt.Println("isExitnil")
	}

	return isExit
}


