package models

import (
	"W2OlineWinterAssignmentTest/utils"
	"fmt"
)

type User struct {
	ID int
	Username string
	Password string
	IfAdmin int // 0 普通用户，1 管理员
}

//----------------数据库操作----------------

//插入操作
func InsertUser(user User)(int64,error){
	return utils.ModifyDB("insert into users(USERNAME,PASSWORD,IFADMIN) values(?,?,?)",
		user.Username,user.Password,user.IfAdmin)
}

//按条件查询
func QueryUserWithCondition(condition string) int {
	sql := fmt.Sprintf("select id from users %s",condition)
	fmt.Println(sql)
	row := utils.QueryRowDB(sql)
	id := 0
	_ = row.Scan(&id)
	return id
}

//根据用户名查询id
func QueryUserWithUsername(username string) int{
	sql := fmt.Sprintf("where username='%s'",username)
	return QueryUserWithCondition(sql)
}

//根据用户名与对应密码查询id
func QueryUserWithUnAndPwd(username ,password string) int {
	sql := fmt.Sprintf("where username='%s' and password='%s'",username,password)
	return QueryUserWithCondition(sql)
}

func QueryAdmin(username string) int {
	sql := fmt.Sprintf("where username='%s'",username)
	sql2 := fmt.Sprintf("select ifadmin from users %s",sql)
	fmt.Println(sql2)
	row := utils.QueryRowDB(sql2)
	var ifadmin int
	_ =row.Scan(&ifadmin)
	fmt.Println("ifadmin:",ifadmin)
	return ifadmin
}





