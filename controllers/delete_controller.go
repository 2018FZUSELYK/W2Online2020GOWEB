package controllers

import (
	"W2OlineWinterAssignmentTest/models"
	"encoding/json"
	"fmt"
)

type DeleteController struct {
	BaseController
}

func (this *DeleteController) Post(){
	IsAdmin := this.IsAdmin
	bookJson := &models.BookJson{}
	_ = json.Unmarshal(this.Ctx.Input.RequestBody, bookJson)
	fmt.Println("JSON:",bookJson)
	Title := bookJson.BookName
	Author := bookJson.BookAuthor
	fmt.Println("Title:",Title)
	fmt.Println("Author:",Author)
	var i int64
	if IsAdmin == false{
		i,_ = models.DelBookByUserWithTitleAndAuthor(Title,Author)
	} else {
		i,_ = models.DelBookByAdminWithTitleAndAuthor(Title,Author)
	}
	if i > 0{
		this.Data["json"] = map[string]interface{}{"code":1,"message":"删除成功"}
	} else {
		this.Data["json"] = map[string]interface{}{"code":0,"message":"不存在此书"}
	}

	this.ServeJSON()
}
