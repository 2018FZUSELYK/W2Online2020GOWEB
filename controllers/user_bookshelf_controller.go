package controllers

import (
	"W2OlineWinterAssignmentTest/models"
	"fmt"
)

type UserBookShelfController struct {
	BaseController
}

func (this *UserBookShelfController) Get(){
	IsLogin := this.IsLogin
	IsAdmin := this.IsAdmin
	IsWriter := this.IsWriter
	LoginUser := this.LoginUserStr
	var bookList []models.Book
	if IsAdmin == false && IsWriter ==false{
		bookList,_ = models.QueryBookByUsername(LoginUser)
		fmt.Println("用户收藏书籍：",bookList)
	} else if IsAdmin == true && IsWriter == false {
		bookList,_ =models.QueryAllBooksByAdmin()
		fmt.Println("所有书籍：",bookList)
	} else {
		bookList,_ = models.QueryBookByUsername(LoginUser)
		fmt.Println("作家收藏书籍：",bookList)
	}
	BookShelf := models.MakeHomePageRecomd(bookList,IsLogin)
	this.Data["BookShelf"] = BookShelf
	this.Data["IsLogin"] = IsLogin
	this.TplName = "newbookShelf.html"
}

func (this *UserBookShelfController) Post(){
	IsLogin := this.IsLogin
	IsAdmin := this.IsAdmin
	IsWriter := this.IsWriter

	this.Data["json"] = map[string]interface{}{"code":IsLogin,"isadmin":IsAdmin,"iswriter":IsWriter}
	this.ServeJSON()

}
