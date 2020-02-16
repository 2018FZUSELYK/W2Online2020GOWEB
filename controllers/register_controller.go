package controllers

import (
	"W2OlineWinterAssignmentTest/models"
	"W2OlineWinterAssignmentTest/utils"
	"fmt"
	"github.com/astaxie/beego"
)

type RegisterController struct{
	beego.Controller
}

//GET方法
func (this *RegisterController) Get(){
	this.TplName = "register2.html"
}

//POST方法
func (this *RegisterController) Post(){

	username := this.GetString("username")
	password := this.GetString("password")
	fmt.Println("username:",username)
	fmt.Println("password:",password)
	//判断该用户名是否被注册过 用户名是唯一的
	id := models.QueryUserWithUsername(username)
	fmt.Println(id)
	//fmt.Println("是否为ajax请求1",this.IsAjax())
	if id != 0{
		fmt.Println("id:",id,"   用户名已存在")
		this.Data["json"] = map[string]interface{}{"code":0,"message":"用户名已存在"}
		this.ServeJSON()
		return
	}

	//密码采用MD5哈希方式加密，在登录时也是把用户的密码MD5哈希后进行比较
	password = utils.MD5Hash(password)
	fmt.Println("MD5Hash:",password)

	user := models.User{0,username,password,0}
	_, err := models.InsertUser(user)

	if err != nil{
		this.Data["json"] = map[string]interface{}{"code":0,"message":"注册失败"}
	} else {
		this.Data["json"] = map[string]interface{}{"code":1,"message":"注册成功"}
	}
	this.ServeJSON()
	//fmt.Println("是否为ajax请求2",this.IsAjax())

	return
}

