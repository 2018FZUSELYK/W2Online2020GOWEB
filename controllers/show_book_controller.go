package controllers

import (
	"W2OlineWinterAssignmentTest/models"
	"fmt"
	"strconv"
)

type ShowBookController struct {
	BaseController
}

func (this *ShowBookController) Get(){
	idString := this.Ctx.Input.Param(":id")

	id, _ := strconv.Atoi(idString)
	fmt.Println("ID:",id)

	//获取id所对应文章信息
	book := models.QueryBookWithId(id)

	this.Data["Title"] = book.Title
	this.Data["Short"] = book.Short
	this.Data["Author"] = book.Author
	this.Data["Link"] = "/book/"+strconv.Itoa(id)
	this.Data["AddLink"] = "/addfavo/" + strconv.Itoa(id)
	this.TplName = "bookInfo.html"
}

func (this *ShowBookController) Post(){
	IsLogin := this.IsLogin
	IsAdmin := this.IsAdmin
	IsWriter := this.IsWriter
	fmt.Println(this.IsAjax())
	this.Data["json"] = map[string]interface{}{"code":IsLogin,"isadmin":IsAdmin,"iswriter":IsWriter}
	this.ServeJSON()
}
