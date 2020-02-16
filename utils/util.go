package utils

import (
	"crypto/md5"
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var db *sql.DB //定义一个全局db变量

func InitMysql() {
	fmt.Println("InitMySql......")
	DriverName := beego.AppConfig.String("driverName")

	//注册数据库驱动
	_ = orm.RegisterDriver(DriverName, orm.DRMySQL)

	//数据库连接
	user := beego.AppConfig.String("MySqlUser")
	password := beego.AppConfig.String("MySqlPwd")
	host := beego.AppConfig.String("host")
	port := beego.AppConfig.String("port")
	dbname := beego.AppConfig.String("DbName")
	//连接数据库用
	dbConnect := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8"
	//dbConnect := "root:123456@tcp(127.0.0.1:3306)/mynovel?charset=utf8"

	db1, err := sql.Open(DriverName,dbConnect)

	if err != nil {
		fmt.Println(err.Error())
	} else {
		db = db1
		CreateTableWithUser()
		CreateTableWithNovel()
		CreateTableWithFavorite()
	}
	return

}
//对数据库进行操作 参数为可变变量
func ModifyDB(sql string, args ...interface{}) (int64, error){
	result, err := db.Exec(sql, args...)
	if err != nil{
		log.Println(err)
		return 0,err
	}
	cnt, err := result.RowsAffected()
	if err != nil{
		log.Println(err)
		return 0,err
	}
	return cnt,nil
}
//创建用户表
//设计用户表，username为用户名，password为密码，ifadmin为是否为管理员，id是主键
func CreateTableWithUser() {
	sql := `CREATE TABLE IF NOT EXISTS users(
			id INT(4) PRIMARY KEY AUTO_INCREMENT NOT NULL,
			username VARCHAR(64),
			password VARCHAR(64),
			ifadmin INT(4)
			);`
	_, _ = ModifyDB(sql)
}

//创建小说表
//设计小说表，title为小说书名，author为作者，tags为分类，short为简介，content为内容，createtime为发布日期,favotime为收藏次数
func CreateTableWithNovel() {
	sql := `CREATE TABLE IF NOT EXISTS novels(
			id INT(4) PRIMARY KEY AUTO_INCREMENT NOT NULL,
			title VARCHAR(30),
			author VARCHAR(20),
			tags VARCHAR(30),
			short VARCHAR(255),
			content LONGTEXT,
			createtime INT(10),
			favotime INT(10)
			);`
	_, _ = ModifyDB(sql)
}

//创建收藏表
//设计收藏表
//收藏表排序 select * from favorite order by username asc,bookid asc, createtime asc;
func CreateTableWithFavorite() {
	sql := `CREATE TABLE IF NOT EXISTS favorite(
			id INT(4) PRIMARY KEY AUTO_INCREMENT NOT NULL,
			username VARCHAR(64),
			bookId INT(4),
			title VARCHAR(30),
			author VARCHAR(20),
			tags VARCHAR(30),
			short VARCHAR(255),
			content LONGTEXT,
			createTime INT(10)
			);`
	_, _ = ModifyDB(sql)
}
//查询 用于之后的数据库操作
func QueryRowDB(sql string) *sql.Row{
	return db.QueryRow(sql)
}

func QueryDB(sql string) (*sql.Rows, error) {
	return db.Query(sql)
}

//MD5加密
func MD5Hash(str string) string{
	data := []byte(str)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)
	return md5str
}