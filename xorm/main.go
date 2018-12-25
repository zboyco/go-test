package main

import (
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/core"
	"github.com/go-xorm/xorm"
)

type User struct {
	ID        int `xorm:"pk autoincr"`
	UserName  string
	UserPwd   string
	CreatedAt time.Time `xorm:"created"`
	Flag      int
}

var engine *xorm.Engine
var err error
var affected int64
var has bool

func main() {

	// 新建orm引擎
	engine, err = xorm.NewEngine("mysql", "root:tamanya233@tcp(192.168.2.100:3307)/XormTest?charset=utf8")

	// 有错退出
	checkErr(err)

	// 设置字段映射方法
	engine.SetMapper(core.SameMapper{})

	// 插入
	insertID := insert()

	// // 查询
	query()

	// // 存在
	exist()

	// // 更新
	update()

	// //删除
	del(insertID)
}

func insert() int {

	// 事务
	session := engine.NewSession()
	defer session.Close()

	err = session.Begin()
	checkErr(err)

	// 插入数据
	newUser := User{
		UserName: "root",
		UserPwd:  "test",
		Flag:     1,
	}
	affected, err = engine.Insert(&newUser)
	if err != nil {
		session.Rollback()
		checkErr(err)
	}
	err = session.Commit()
	checkErr(err)
	log.Println("新插入", affected, "条数据,ID:", newUser.ID)
	return newUser.ID
}

func query() {

	var user User
	// ID获取单条
	engine.ID(1).Get(&user)
	log.Println("查出数据,ID:", user.ID)

	// where获取单条
	has, err = engine.Where("UserName = ?", "root").Get(&user)
	checkErr(err)
	log.Println("查出数据", has, "ID:", user.ID)

	// user结构体中已有的非空数据来获得单条数据
	user = User{UserName: "root"}
	engine.Get(&user)
	log.Println("查出数据,ID:", user.ID)

	// find 获取多条 where
	var users []User
	engine.Where("UserName = ?", "root").Find(&users)
	for _, item := range users {
		log.Println("查出数据:", item)
	}

	// find 获取多条 结构体中已有的非空数据来获得
	engine.Find(&users, &User{
		UserName: "root",
	})
	for _, item := range users {
		log.Println("查出数据:", item)
	}

	// find map
	usersMap := make(map[int]User, 0)
	err = engine.Find(&usersMap)
	checkErr(err)
	for _, item := range usersMap {
		log.Println("查出数据:", item)
	}

	// find 查询单个字段
	var ids []int
	err := engine.Table(new(User)).Cols("ID").Find(&ids)
	checkErr(err)
	log.Println("查出IDs:", ids)
}

func exist() {
	// Exist 判断是否存在
	has, err = engine.Exist(&User{
		UserName: "root",
	})
	checkErr(err)
	log.Println("数据是否存在:", has)
	has, err = engine.Where("UserName = ?", "root").Exist(&User{})
	checkErr(err)
	log.Println("数据是否存在:", has)
	has, err = engine.Table(new(User)).Where("UserName = ?", "root").Exist()
	checkErr(err)
	log.Println("数据是否存在:", has)
}

func update() {

	// user := &User{
	// 	UserName: "update",
	// 	UserPwd:  "update",
	// }

	user := new(User)

	user.UserName = "update"
	user.UserPwd = "update"

	// 非0值更新
	affected, err = engine.ID(1).Update(user)
	checkErr(err)
	log.Println("更新", affected, "条数据")

	// 指定列更新
	affected, err = engine.ID(2).Cols("UserName").Update(user)
	checkErr(err)
	log.Println("更新", affected, "条数据")

	// map更新
	affected, err = engine.Table(new(User)).ID(3).Update(map[string]interface{}{"UserPwd": "update"})
	checkErr(err)
	log.Println("更新", affected, "条数据")
}

func del(id int) {
	user := new(User)
	affected, err = engine.ID(id).Delete(user)
	checkErr(err)
	log.Println("删除ID为", id, "的", affected, "条数据")
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
