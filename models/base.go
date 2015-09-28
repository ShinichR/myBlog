package models

import (
	//"github.com/Unknwon/com"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	//"os"
	//"path"
)

const (
	_DB_NAME        = "data/beeblog.db"
	_SQLITE3_DRIVER = "sqlite3"
)

func Init() {
	dbhost := beego.AppConfig.String("dbhost")
	dbport := beego.AppConfig.String("dbport")
	dbuser := beego.AppConfig.String("dbuser")
	dbpassword := beego.AppConfig.String("dbpassword")
	dbname := beego.AppConfig.String("dbname")
	if dbport == "" {
		dbport = "3306"
	}

	dsn := dbuser + ":" + dbpassword + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8"
	orm.RegisterDataBase("default", "mysql", dsn)
	orm.RegisterModel(new(User), new(Tag), new(TagPost), new(Post))
	createTable()
	/*
		if !com.IsExist(_DB_NAME) {
			os.MkdirAll(path.Dir(_DB_NAME), os.ModePerm)
			os.Create(_DB_NAME)
		}
		orm.RegisterModel(new(User),new(Tag),new(TagPost),new(Post))
		orm.RegisterDriver(_SQLITE3_DRIVER, orm.DR_Sqlite)
		orm.RegisterDataBase("default", _SQLITE3_DRIVER, _DB_NAME, 10)*/
}

func createTable() {
	name := "default"                          //数据库别名
	force := false                             //不强制建数据库
	verbose := true                            //打印建表过程
	err := orm.RunSyncdb(name, force, verbose) //建表
	if err != nil {
		beego.Error(err)
	}
}

//返回带前缀的表名
func TableName(str string) string {
	return beego.AppConfig.String("dbprefix") + str
}
