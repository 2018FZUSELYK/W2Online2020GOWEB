package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"reflect"
)

type BaseController struct{
	beego.Controller
	IsLogin bool
	IsAdmin bool
	IsWriter bool
	LoginUser interface{}
	LoginUserStr string
}

//判断是否登录，重写prepare方法用于获取session
func (this *BaseController) Prepare() {
	//获取登录session
	loginuser := this.GetSession("loginuser")
	isadmin := this.GetSession("isadmin")
	iswriter := this.GetSession("iswriter")
	if loginuser != nil {
		//如果session有效则视为登录成功
		this.IsLogin = true
		this.LoginUser = loginuser
		this.LoginUserStr = loginuser.(string)
		fmt.Println("LoginUserStrType:",reflect.TypeOf(this.LoginUserStr))
		fmt.Println("LoginUserStr:",this.LoginUserStr)
		//对用户身份进行判断
		if isadmin != nil{
			this.IsAdmin = true
		} else {
			this.IsAdmin= false
		}
		if iswriter != nil{
			this.IsWriter = true
		} else {
			this.IsWriter= false
		}
	} else {
		this.IsLogin = false
	}
	fmt.Println("LoginUser:",loginuser)
	fmt.Println("IsAdmin:",this.IsAdmin)
	fmt.Println("IsLogin:",this.IsLogin)
	fmt.Println("IsWriter:",this.IsWriter)
	this.Data["LoginUser"] = this.LoginUser
	this.Data["IsLogin"] = this.IsLogin
	this.Data["IsAdmin"] = this.IsAdmin
	this.Data["IsWriter"] = this.IsWriter
}

