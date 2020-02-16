package controllers

import (
	"W2OlineWinterAssignmentTest/models"
	"W2OlineWinterAssignmentTest/utils"
	"fmt"
	"github.com/astaxie/beego"
)

type LoginController struct {
	beego.Controller
}

func (this *LoginController) Get() {
	this.TplName = "login.html"
}

func (this *LoginController) Post(){
	username := this.GetString("username")
	password := this.GetString("password")
	fmt.Println("username:",username,"password:",password)

	id := models.QueryUserWithUnAndPwd(username,utils.MD5Hash(password))
	fmt.Println("id:",id)
	ifadmin := models.QueryAdmin(username)
	fmt.Println("ifadmin:",ifadmin)
	fmt.Println("isAdmin",ifadmin)
	if id!=0 {
		//设置session可以将数据设置到cookie，再由cookie来判断用户是谁
		this.SetSession("loginuser",username)
		if ifadmin == 1{
			this.SetSession("isadmin",ifadmin)
		} else if ifadmin == 2{
			this.SetSession("iswriter",ifadmin)
		}
		this.Data["json"] = map[string]interface{}{"code": 1, "message": "登录成功","isadmin":ifadmin}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 0, "message": "登录失败","isadmin":ifadmin}
	}
	this.ServeJSON()
}