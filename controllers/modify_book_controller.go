package controllers

import (
	"W2OlineWinterAssignmentTest/models"
	"fmt"
)

type ModifyBookController struct {
	BaseController
}

func (this *ModifyBookController) Get(){
	IsAdmin := this.IsAdmin
	fmt.Println("IsAdmin:",IsAdmin)

	id,_ := this.GetInt("id")
	fmt.Println("Id:",id)
	book := models.QueryBookWithId(id)
	fmt.Println("所查询书籍信息：",book)
	this.Data["Book"] = book
	this.TplName = "modifybook.html"

}

func (this *ModifyBookController) Post(){
	Id, _ := this.GetInt("id")
	fmt.Println("Id:", Id)
	book := models.QueryBookWithId(Id)
	Title := this.GetString("Title")
	Author := this.GetString("Author")
	Tags := this.GetString("Tags")
	Short := this.GetString("Short")
	Content := this.GetString("Content")
	newbook := models.Book{Id:Id,Title:Title,Author:Author,Tags:Tags,Short:Short,Content:Content,Createtime:book.Createtime,Favotime:book.Favotime}
	_,err := models.ModifyBook(newbook)

	if err != nil{
		this.Data["json"] = map[string]interface{}{"code":0,"message":"修改成功"}
	} else {
		this.Data["json"] = map[string]interface{}{"code":1,"message":"修改失败"}
	}
	this.ServeJSON()
}
