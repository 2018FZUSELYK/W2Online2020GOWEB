package controllers

import (
	"W2OlineWinterAssignmentTest/models"
	"fmt"
	"strconv"
)

type AddFavoriteBookController struct {
	BaseController
}
func (this *AddFavoriteBookController) Get(){
	Username := this.LoginUserStr
	idString := this.Ctx.Input.Param(":id")
	id,_ := strconv.Atoi(idString)
	fmt.Println(id)
	i,_ := models.AddBookByUserWithId(id,Username)
	if i>0{
		this.Data["json"] = map[string]interface{}{"code":1,"message":"收藏成功"}
	} else {
		this.Data["json"] = map[string]interface{}{"code":0,"message":"未知错误"}
	}
	this.ServeJSON()
}
//func (this *AddFavoriteBookController) Post(){
//	bookJson := &models.BookJson{}
//	_ = json.Unmarshal(this.Ctx.Input.RequestBody, bookJson)
//	fmt.Println("JSON:",bookJson)
//	Title := bookJson.BookName
//	Author := bookJson.BookAuthor
//	fmt.Println("Title:",Title)
//	fmt.Println("Author:",Author)
//	this.Data["json"] = map[string]interface{}{"code":1}
//	this.ServeJSON()
//}