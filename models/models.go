package models

import (
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

const (
	_DB_NAME      = "test"
	_MYSQL_DRIVER = "mysql"
)

type Category struct {
	Id              int64
	Title           string
	Created         time.Time `orm:"index"`
	Views           int64     `orm:"index"`
	TopicTime       time.Time `orm:"index"`
	TopicCount      int64
	TopicLastUserId int64
}

type Topic struct {
	Id               int64
	Uid              int64
	Title            string
	Content          string `orm:"size(3000)"`
	Attachment       string
	Created          time.Time `orm:"index"`
	Updateed         time.Time `orm:"index"`
	Views            int64     `orm:"index"`
	Author           string
	ReplyCount       int64
	ReplyTime        time.Time `orm:"index"`
	RepleyLastUserId int64
}

func RegistDB() {
	orm.RegisterModel(new(Category), new(Topic))
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "root:arvin@(127.0.0.1:3306)/test?charset=utf8")
}
