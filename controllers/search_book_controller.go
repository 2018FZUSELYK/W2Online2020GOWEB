package controllers

import (
	"W2OlineWinterAssignmentTest/models"
	"fmt"
)

type SearchBookController struct {
	BaseController
}

func (this *SearchBookController) Post(){
	searchBook := this.GetString("searchBook")
	book,_ := models.QueryBookByTitle(searchBook)
	resultBookList := models.MakeHomePageRecomd(book,true)
	fmt.Println(book)
	this.Data["Result"] = resultBookList
	this.TplName = "searchBook.html"
}