package models

import (
	"strconv"
	"time"

	"github.com/astaxie/beego"
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
	Created          string `orm:"index"`
	Updated          string `orm:"index"`
	Views            int64  `orm:"index"`
	Author           string
	ReplyCount       int64
	ReplyTime        string `orm:"index"`
	RepleyLastUserId int64
	Category         string
}
type Reply struct {
	Id         int64
	Tid        int64
	Nickname   string
	Content    string
	Createtime string
}

func RegistDB() {
	host := beego.AppConfig.String("DBhost")
	dataSource := "root:arvin@(" + host + ":3306)/test?charset=utf8"
	orm.RegisterModel(new(Category), new(Topic), new(Reply))
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", dataSource)
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

func AddTopic(title, content, cateTitle string) error {
	o := orm.NewOrm()
	timestamp := time.Now().Unix()
	tm := time.Unix(timestamp, 0)
	var s string = tm.Format("2006-01-02 15:04:05")

	topic := &Topic{
		Title:     title,
		Content:   content,
		Created:   s,
		Updated:   s,
		ReplyTime: s,
		Category:  cateTitle,
	}

	_, err := o.Insert(topic)

	if err != nil {
		return err
	}
	UpdateCategoryCount(cateTitle, "add")
	return err
}

func GetAllTopics(cate string, isDesc bool) ([]*Topic, error) {
	o := orm.NewOrm()

	Topics := make([]*Topic, 0)

	qs := o.QueryTable("topic")

	var err error
	if isDesc {
		if len(cate) > 0 {
			qs = qs.Filter("category", cate)
		}
		_, err = qs.OrderBy("-created").All(&Topics)
	} else {
		_, err = qs.All(&Topics)
	}

	return Topics, err
}

func GetTopic(uId string) (*Topic, error) {
	nId, err := strconv.ParseInt(uId, 10, 64)
	if err != nil {
		return nil, err
	}
	o := orm.NewOrm()

	qs := o.QueryTable("topic")

	topic := new(Topic)
	err = qs.Filter("id", nId).One(topic)
	if err != nil {
		return nil, err
	}

	topic.Views++

	_, err = o.Update(topic)
	if err != nil {
		return nil, err
	}

	return topic, nil
}

func ModifyTopic(uId, title, content string) error {
	nId, err := strconv.ParseInt(uId, 10, 64)
	if err != nil {
		return err
	}

	o := orm.NewOrm()
	topic := &Topic{Id: nId}
	err = o.Read(topic)
	if err != nil {
		return err
	}
	topic.Title = title
	topic.Content = content
	timetamp := time.Now().Unix()
	tm := time.Unix(timetamp, 0)
	topic.Updated = tm.Format("2006-01-02 15:04:05")
	o.Update(topic)
	return nil
}

func DelTopic(tId string) error {
	nId, err := strconv.ParseInt(tId, 10, 64)
	if err != nil {
		return err
	}

	o := orm.NewOrm()

	topic := &Topic{Id: nId}
	o.Read(topic)
	_, err = o.Delete(topic)
	if err != nil {
		return err
	}

	UpdateCategoryCount(topic.Category, "del")
	return nil
}

func AddReply(tId, nick, content string) error {
	tIdnum, err := strconv.ParseInt(tId, 10, 64)
	if err != nil {
		return err
	}

	timetamp := time.Now().Unix()
	tm := time.Unix(timetamp, 0)

	reply := &Reply{Tid: tIdnum, Nickname: nick, Content: content, Createtime: tm.Format("2006-01-02 15:04:05")}
	o := orm.NewOrm()
	_, err = o.Insert(reply)
	if err != nil {
		return err
	}
	return nil

}

func GetAllReplys(tId string, IsDesc bool) ([]*Reply, error) {
	tIdnum, err := strconv.ParseInt(tId, 10, 64)
	if err != nil {
		return nil, err
	}

	replys := make([]*Reply, 0)

	o := orm.NewOrm()

	qs := o.QueryTable("reply")

	if IsDesc {
		qs.OrderBy("-createtime").All(&replys)
	} else {
		_, err = qs.Filter("tid", tIdnum).All(&replys)
		if err != nil {
			return nil, err
		}

	}

	return replys, nil
}

func DelReply(id string) error {
	idNum, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}

	reply := &Reply{Id: idNum}
	o := orm.NewOrm()
	_, err = o.Delete(reply)
	return err
}

func UpdateCategoryCount(cateTitle, op string) error {
	o := orm.NewOrm()
	switch op {
	case "add":
		{
			cate := &Category{}

			qs := o.QueryTable("category")
			err := qs.Filter("title", cateTitle).One(cate)
			if err != nil {
				return err
			}
			cate.TopicCount++
			_, err = o.Update(cate)
		}
	case "del":
		{
			cate := &Category{}

			qs := o.QueryTable("category")
			err := qs.Filter("title", cateTitle).One(cate)
			if err != nil {
				return err
			}
			cate.TopicCount--
			_, err = o.Update(cate)
		}
	}
	return nil
}
