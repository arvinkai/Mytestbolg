package models

import (
	"strconv"
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
	orm.RegisterDataBase("default", "mysql", "root:arvin@(192.168.80.154:3306)/test?charset=utf8")
}

func AddCategory(name string) error {
	o := orm.NewOrm()

	cate := &Category{Title: name}

	qs := o.QueryTable("category")
	err := qs.Filter("Title", name).One(cate)
	if err == nil {
		return err
	}

	_, err = o.Insert(cate)
	if err != nil {
		return err
	}

	return nil
}

func GetAllCategories() ([]*Category, error) {
	o := orm.NewOrm()

	cates := make([]*Category, 0)
	qs := o.QueryTable("category")
	_, err := qs.All(&cates)

	return cates, err
}

func DelCategory(id string) error {
	nId, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}

	o := orm.NewOrm()

	cate := &Category{Id: nId}
	_, err = o.Delete(cate)
	if err != nil {
		return err
	}

	return nil
}
