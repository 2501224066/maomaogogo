package models

import (
	"fmt"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

var O orm.Ormer

type User struct {
	Uid        int `orm:"pk"`
	Nickname   string
	Email      string
	Password   string
	AvatarUrl  string
	InShort    string
	Introduce  string
	Sex        string
	Birthday   time.Time
	Created_at time.Time  `orm:"auto_now_add;type(datetime)"`
	Updated_at time.Time  `orm:"auto_now;type(datetime)"`
	Article    []*Article `orm:"reverse(many)"`
}

type SystemSetting struct {
	Id    int `orm:"pk"`
	Key   string
	Name  string
	Value string
}

type Node struct {
	NodeId int `orm:"pk"`
	Name   string
}

type Article struct {
	ArticleId int `orm:"pk"`
	Title     string
	Content   string
	FaceImg   string
	FaceText  string
	Node      string
	Tag       string
	ReadNum   int
	LikeNum   int
	ThankNum  int
	Status    int
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt time.Time `orm:"auto_now;type(datetime)"`
	User      *User     `orm:"rel(fk)"`
}

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	dbUser := beego.AppConfig.String("db::user")
	dbPassword := beego.AppConfig.String("db::password")
	dbHost := beego.AppConfig.String("db::host")
	dbPort := beego.AppConfig.String("db::port")
	dbDatabase := beego.AppConfig.String("db::database")
	dbCharset := beego.AppConfig.String("db::charset")
	maxIdle := beego.AppConfig.DefaultInt("db::maxIdle", 50)
	maxConn := beego.AppConfig.DefaultInt("db:: maxConn", 300)
	loc := "Local"
	dbLink := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&loc=%v", dbUser, dbPassword, dbHost, dbPort, dbDatabase, dbCharset, loc)
	orm.RegisterDataBase("default", "mysql", dbLink, maxIdle, maxConn)

	orm.RegisterModel(
		new(User),
		new(SystemSetting),
		new(Node),
		new(Article))
	O = orm.NewOrm()
}
