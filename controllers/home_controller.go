package controllers

import (
	"W2OlineWinterAssignmentTest/models"
	"fmt"
)

type HomeController struct {
	//beego.Controller
	BaseController
}
func (this *HomeController) Update(){
	fmt.Println("是否Ajax请求：",this.IsAjax())
	this.Data["json"] = map[string]interface{}{"status":this.IsLogin,
		  									   "name":this.LoginUser,
		  									   "isadmin":this.IsAdmin,
		  									   "iswriter":this.IsWriter}
	this.ServeJSON()
}
//请求类型为get 跳转到home页面
func (this *HomeController) Get(){
	//fmt.Println("IsLogin:",this.IsLogin,this.LoginUser)
	//---------获取小说以及页数部分---------
	page, _ := this.GetInt("page")
	if page <= 0 {
		page = 1
	}
	var bookList []models.Book
	bookList, _ = models.FindBookWithPage(page)
	this.Data["PageCode"] = models.SetPageFooterCode(page)
	this.Data["HasFooter"] = true

	//this.Data["BookList"] = bookList
	this.Data["BookList"] = models.MakeHomePageRecomd(bookList,this.IsLogin)



	//---------获取标签以及数量部分---------
	tags := models.QueryTagsByStrings("tags")
	fmt.Println(models.GetTagsMap(tags))
	this.Data["Tags"] = models.GetTagsMap(tags)


	//---------每日推荐书本 暂定随机抽取1本---------
	ranBook,_ := models.QueryBookByRandom()
	fmt.Println("每日推荐:",ranBook)
	this.Data["Daily"] = models.MakeHomePageRecomd(ranBook,true)
	this.TplName = "index.html"
}

func (this *HomeController) Post(){
	searchBook := this.GetString("searchBook")
	fmt.Println("所查找关键词：",searchBook)
	bookList,_ := models.QueryBookByTitle(searchBook)
	fmt.Println("搜索结果：",bookList)
	this.Data["json"] = map[string]interface{}{"code":1,"message":"查询成功"}
	this.ServeJSON()

}