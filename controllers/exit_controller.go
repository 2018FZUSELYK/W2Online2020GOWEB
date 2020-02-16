package controllers

import "fmt"

type ExitController struct {
	BaseController
}

func (this *ExitController) Get(){
	//删除用户“loginuser”的session 退出登录
	this.DelSession("loginuser")
	isadmin := this.GetSession("isadmin")
	iswriter := this.GetSession("iswriter")
	if isadmin != nil{
		this.DelSession("isadmin")
		fmt.Println("IsAdmin session del success!")
	}
	if iswriter != nil{
		this.DelSession("iswriter")
		fmt.Println("IsWriter session del success!")
	}
	//重定向回主页
	this.Redirect("/",302)
}