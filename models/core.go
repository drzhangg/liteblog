package models

import (
	"github.com/jinzhu/gorm"
	"os"
	"time"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var (
	db *gorm.DB
)

func init() {
	//logs.Info("hello hello hello")

	/*var err error

	if db,err = gorm.Open("mysql","root:root@tcp(127.0.0.1:3306)/liteblog?charset=utf8");err != nil{
		fmt.Println(err)
		return
	}

	defer db.Close()*/



	var err error
	if err = os.MkdirAll("data", 0777); err != nil {
		panic("failed " + err.Error())
	}
	db, err = gorm.Open("sqlite3", "data/data.db")
	if err != nil {
		panic("failed to connect database")
	}
	//自动添加表结构
	db.AutoMigrate(&User{},&Note{},&Message{})

	//如果数据库里面没有用户数据，我们新增一条admin记录
	var count int
	if err := db.Model(&User{}).Count(&count).Error; err == nil && count == 0 {
		db.Create(&User{
			Name:"admin",
			Email:"admin@qq.com",
			Pwd:"123",
			Avatar:"/static/images/info-img.png",
			Role:0,
		})
	}

}

type Model struct {
	ID uint		`gorm:"primary_key" json:"id"`
	CreatedAt time.Time		`json:"created_at"`
	UpdatedAt time.Time		`json:"updated_at"`
	DeletedAt *time.Time	`sql:"index" json:"deleted_at"`
}
