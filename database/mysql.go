package database

import (
	"fmt"
	"log"
	"strconv"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

// Test blablabla
type Test struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Chips int    `json:"chips"`
	Phone string `json:"phone"`
	Time  string `json:"time"`
}

var tests []Test

func init() {
	log.Println("<mysql init func>")
	//註冊驅動
	orm.RegisterDriver("mysql", orm.DRMySQL)
	//連線
	orm.RegisterDataBase("default", "mysql", "dev:dev123@tcp(192.168.50.253:3306)/CatHome?charset=utf8")
	//註冊定義的 model
	orm.RegisterModel(new(Test))

	// 建立 table
	orm.RunSyncdb("default", false, true)

}

//Query ...
func (test *Test) Query() []Test {
	log.Println("<mysql Test.Query func>")
	var j, i int64

	orm.Debug = true
	o := orm.NewOrm()
	o.Using("default") // 默认使用 default，你可以指定为其他数据库

	var tests []Test
	num, err := o.Raw("SELECT * FROM test ").QueryRows(&tests)
	fmt.Println(tests)

	if err == nil {
		fmt.Println("test nums: ", num)
	}

	i = num
	for j = 0; j < i; j++ {
		fmt.Println(j, tests[j])
	}
	return tests
}

// QueryById ...
func (test *Test) QueryById(url string) Test {
	log.Println("<mysql Test.QueryById func>")
	var j, i int64
	orm.Debug = true
	o := orm.NewOrm()
	o.Using("default") // 默认使用 default，你可以指定为其他数据库
	idx, err := strconv.Atoi(url)
	if err != nil {
		log.Println("QueryById strconv err,url:", url)

		return Test{}
	}
	var querydatas []Test

	num, err := o.Raw("SELECT * FROM test WHERE id = ?", idx).QueryRows(&querydatas)
	fmt.Println(querydatas)

	if err == nil {
		fmt.Println("user nums: ", num)
	}

	i = num
	for j = 0; j < i; j++ {
		fmt.Printf("Element[%d] = %d\n", j, querydatas[j])
	}

	return querydatas[0]
}

//Insert ...
/* 	使用方法 (database為套件暱稱)
	test := database.Test{
 	Name:  "test",
 	Chips: 5000,
 	Phone: "099999999",
 	Time:  time.Now().Format("2006-01-02 15:04:05"),
 	}
 test.Insert()*/
func (test *Test) Insert() {
	log.Println("<mysql Test.Insert func>")
	orm.Debug = true
	o := orm.NewOrm()
	o.Using("default") // 默认使用 default，你可以指定为其他数据库

	log.Println(o.Insert(test))

}

//Delete ...
/* 	使用方法 (database為套件暱稱)
	test := database.Test{
 	Name:  "test",
 	Chips: 5000,
 	Phone: "099999999",
 	Time:  time.Now().Format("2006-01-02 15:04:05"),
 	}
 test.Delete()*/
func (test *Test) Delete() {
	log.Println("<mysql Test.delete func>")
	orm.Debug = true
	o := orm.NewOrm()
	o.Using("default") // 默认使用 default，你可以指定为其他数据库
	log.Println(o.Delete(test))
}

//Update ...
/* 	使用方法 (database為套件暱稱)
	test := database.Test{
 	Name:  "test",
 	Chips: 5000,
 	Phone: "099999999",
 	Time:  time.Now().Format("2006-01-02 15:04:05"),
 	}
 test.Update()*/
func (test *Test) Update() {
	log.Println("<mysql Test.update func>")
	orm.Debug = true
	o := orm.NewOrm()
	o.Using("default") // 默认使用 default，你可以指定为其他数据库
	log.Println(o.Update(test))
}
