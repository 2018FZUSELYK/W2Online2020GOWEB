package models

import (
	"W2OlineWinterAssignmentTest/utils"
	"log"

	//"crypto/subtle"
	"fmt"
	"github.com/astaxie/beego"
	"strconv"
)

type Book struct {
	Id int
	Title string
	Author string
	Tags string
	Short string
	Content string
	Createtime int64
	Favotime int

}
type BookJson struct {
	BookName     string `json:"book_name"`
	BookAuthor   string `json:"book_author"`
	BookClasses  string `json:"book_classes"`
	IdentityCode string `json:"identityCode"`
}
//根据页码查询文章
func FindBookWithPage(page int) ([]Book,error) {
	num, _ :=beego.AppConfig.Int("booklistpagenum")
	page--
	fmt.Println("------page:",page)
	sql := fmt.Sprintf("limit %d,%d",page*num,num)
	return QueryBookWithCondition(sql)
}

func QueryBookWithCondition(sql string) ([]Book,error){
	sql = "SELECT ID,TITLE,AUTHOR,TAGS,SHORT,CONTENT,CREATETIME,FAVOTIME FROM NOVELS " + sql
	rows, err := utils.QueryDB(sql)
	if err != nil {
		return nil, err
	}
	var bookList []Book
	for rows.Next() {
		id := 0
		title := ""
		author := ""
		tags := ""
		short := ""
		content := ""
		var createtime int64
		createtime = 0
		var favotime int
		favotime = 0
		_ = rows.Scan(&id, &title, &author, &tags, &short, &content, &createtime, &favotime)
		book := Book{id,title,author,tags,short,content,createtime,favotime}
		bookList = append(bookList,book)
	}
	return bookList,nil
}


//-------------翻页功能部分-------------

var bookRowsNumber = 0

//首次获取表内行数使采取直接统计数量的方法
func GetBookRowsNumber() int{
	if bookRowsNumber == 0 {
		bookRowsNumber = QueryBookRowNumber()
	}
	return bookRowsNumber
}

//查询表内总行数
func QueryBookRowNumber() int {
	row := utils.QueryRowDB("SELECT COUNT(ID) FROM NOVELS")
	num := 0
	_ = row.Scan(&num)
	return num
}

//直接设置总行数
func SetBookRowNumber() {
	bookRowsNumber = QueryBookRowNumber()
}

//-------------查询文章-------------
func QueryBookWithId(id int) Book{
	row := utils.QueryRowDB("SELECT ID,TITLE,AUTHOR,TAGS,SHORT,CONTENT,CREATETIME,FAVOTIME FROM NOVELS WHERE ID="+ strconv.Itoa(id))
	title := ""
	author := ""
	tags := ""
	short := ""
	content := ""
	var createtime int64 = 0
	var favotime int = 0
	_ = row.Scan(&id, &title, &author, &tags, &short, &content, &createtime, &favotime)
	book := Book{id,title,author,tags,short,content,createtime,favotime}
	return book
}

//-------------查询标签-------------
func QueryTagsByStrings(tag string) []string{
	sql := fmt.Sprintf("SELECT %s FROM NOVELS",tag)
	rows,err := utils.QueryDB(sql)
	if err!=nil{
		log.Println(err)
	}
	var tagsList []string
	for rows.Next(){
		tags := ""
		_ = rows.Scan(&tags)
		tagsList = append(tagsList,tags)
	}
	return tagsList
}

//-------------书名模糊搜索-------------
func QueryBookByTitle(title string)([]Book,error){
	sql := "WHERE TITLE LIKE '%" + title + "%'"
	sql += " OR TITLE LIKE '%" + title + "'"
	sql += " OR TITLE LIKE '" + title + "%'"
	sql += " OR TITLE LIKE '" + title + "'"
	fmt.Println("模糊搜索指令:",sql)
	return QueryBookWithCondition(sql)
}

//-------------标题模糊搜索-------------
func QueryBookByTags(tags string)([]Book,error){
	sql := "WHERE TAGS LIKE '%&" + tags + "&%'"
	sql += " OR TAGS LIKE '%&" + tags + "'"
	sql += " OR TAGS LIKE '" + tags + "&%'"
	sql += " OR TAGS LIKE '" + tags + "'"
	fmt.Println("模糊搜索指令:",sql)
	return QueryBookWithCondition(sql)
}

//-------------数据库随机读取-------------
func QueryBookByRandom()([]Book,error){
	sql := "select * from novels order by rand() limit 1"
	rows,err := utils.QueryDB(sql)
	if err != nil{
		return nil,err
	}
	var bookList []Book
	for rows.Next(){
		id := 0
		title := ""
		author := ""
		tags := ""
		short := ""
		content := ""
		var createtime int64
		createtime = 0
		var favotime int
		favotime = 0
		_ = rows.Scan(&id, &title, &author, &tags, &short, &content, &createtime, &favotime)
		book := Book{id,title,author,tags,short,content,createtime,favotime}
		bookList = append(bookList,book)
	}

	return bookList,nil

}

//-------------用户根据id收藏书籍-------------
func AddBookByUserWithId(id int,username string)(int64,error){
	book := QueryBookWithId(id)
	sql := "INSERT INTO favorite(USERNAME,BOOKID,TITLE,AUTHOR,TAGS,SHORT,CONTENT,CREATETIME)VALUES(?,?,?,?,?,?,?,?)"
	i, err := utils.ModifyDB(sql,username,book.Id,book.Title,book.Author,book.Tags,book.Short,book.Content,book.Createtime)
	return i,err
}

//-------------根据根据用户名查询收藏书籍-------------
func QueryBookByUsername(username string)([]Book,error){
	fmt.Println("Username:",username)
	sql := fmt.Sprintf("SELECT BOOKID,TITLE,AUTHOR,TAGS,SHORT,CONTENT,CREATETIME FROM FAVORITE WHERE USERNAME='%s'",username)
	rows,err := utils.QueryDB(sql)
	if err != nil{
		return nil,err
	}
	var bookList []Book
	for rows.Next() {
		id := 0
		title := ""
		author := ""
		tags := ""
		short := ""
		content := ""
		var createtime int64
		createtime = 0
		var favotime int
		favotime = 0
		_ = rows.Scan(&id, &title, &author, &tags, &short, &content, &createtime)
		book := Book{id,title,author,tags,short,content,createtime,favotime}
		bookList = append(bookList,book)
	}
	return bookList,nil
}

//-------------用户根据书名和作者删除书籍-------------
func DelBookByUserWithTitleAndAuthor (title string,author string)(int64,error){
	sql := fmt.Sprintf("DELETE FROM favorite WHERE title='%s' and author='%s'",title,author)
	fmt.Println(sql)
	i, err := utils.ModifyDB(sql)
	return i,err
}

//-------------管理员根据书名和作者删除书籍-------------
func DelBookByAdminWithTitleAndAuthor (title string,author string)(int64,error){
	sql := fmt.Sprintf("DELETE FROM novels WHERE title='%s' and author='%s'",title,author)
	fmt.Println(sql)
	i, err := utils.ModifyDB(sql)
	SetBookRowNumber()
	return i,err
}

//-------------管理员显示所有书籍-------------
func QueryAllBooksByAdmin()([]Book,error){
	sql := "SELECT * FROM NOVELS WHERE 1=1"
	rows,err := utils.QueryDB(sql)
	if err != nil{
		return nil,err
	}
	var bookList []Book
	for rows.Next(){
		id := 0
		title := ""
		author := ""
		tags := ""
		short := ""
		content := ""
		var createtime int64
		createtime = 0
		var favotime int
		favotime = 0
		_ = rows.Scan(&id, &title, &author, &tags, &short, &content, &createtime, &favotime)
		book := Book{id,title,author,tags,short,content,createtime,favotime}
		bookList = append(bookList,book)
	}
	return bookList,nil
}

//-------------修改书籍-------------
func ModifyBook(book Book)(int64,error){
	sql := "update novels set title=?,tags=?,short=?,content=? where id=?"
	i,err := utils.ModifyDB(sql,book.Title,book.Tags,book.Short,book.Content,book.Id)
	return i,err
}


