package models

import (
	"fmt"
	"github.com/astaxie/beego"
	//"os"
	"strconv"
	"strings"
)

type HomePageNovelRecommend struct {
	Id int
	Title string
	Author string
	Tags []TagLink
	Short string
	Content string
	CreateTime string
	Link string
	IsLogin bool
	Num int
	FavoTime int
	ClassId string
	DelLink string
}

type TagLink struct {
	TagName string
	TagUrl string
}

type PageFooter struct {
	HasPre bool //是否有上一页
	HasNext bool //是否有下一页
	ShowPage string //目前页数/总页数
	PreLink string //上一页的链接
	NextLink string //下一页的链接
}

//-------------主页显示推荐小说-------------
func MakeHomePageRecomd(book []Book, isLogin bool) []HomePageNovelRecommend {
	//htmlHome := ""
	var homeList []HomePageNovelRecommend
	for k,tmp := range book{
		Id := tmp.Id
		Title := tmp.Title
		Author := tmp.Author
		Tags := createTagsLinks(tmp.Tags)
		Short := tmp.Short
		Content := tmp.Content
		CreateTime := "2020-1-26 13:22"
		//homeRecomd.CreateTime =
		Link := "/bookinfo/" + strconv.Itoa(tmp.Id)
		IsLogin := isLogin
		ClassId := "num" + strconv.Itoa(k+1)
		Num := k+1
		FavoTime := tmp.Favotime
		DelLink := "/delete/" + strconv.Itoa(tmp.Id)
		tmp := HomePageNovelRecommend{Id,Title,Author,Tags,Short,Content,CreateTime,Link,IsLogin,Num,FavoTime,ClassId,DelLink}
		homeList = append(homeList,tmp)
	}

	//fmt.Println("htmlHomePage: ",htmlHome)
	return homeList
}


//-------------标签&标签链接-------------
func createTagsLinks(tags string) []TagLink {
	var tagLink []TagLink
	//用&分割标签 添加标签时用&隔开即可
	tagsTmp := strings.Split(tags,"&")
	for _,tag := range tagsTmp {
		tagLink = append(tagLink,TagLink{tag,"/?tag=" + tag})
	}
	return tagLink
}

//-------------翻页功能-------------
//page是当前所在的页数
func SetPageFooterCode(page int) PageFooter{
	pageCode := PageFooter{}
	//查询总页数
	num := GetBookRowsNumber()
	//从配置文件中读取每页显示的小说本数
	pageRow, _ := beego.AppConfig.Int("booklistpagenum")
	//计算出总页数
	pageRowSum := (num-1)/pageRow + 1
	//在网页上输出页数
	pageCode.ShowPage = fmt.Sprintf("%d/%d",page,pageRowSum)
	//如果当前页数是第一页 则没有上一页 上一页的按钮不能点击
	if page <= 1 {
		pageCode.HasPre = false
	} else {
		pageCode.HasPre = true
	}
	//如果当前页数是最后一页 则没有下一页 下一页的按钮不能点击
	if page >= pageRowSum {
		pageCode.HasNext = false
	} else {
		pageCode.HasNext = true
	}
	//处理上下页的url
	pageCode.PreLink = "/?page=" + strconv.Itoa(page-1)
	pageCode.NextLink = "/?page=" + strconv.Itoa(page+1)

	return pageCode
}




