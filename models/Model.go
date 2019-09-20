package models

import (
	"fmt"
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	// mysql使用
	_ "github.com/go-sql-driver/mysql"
)

// O orm实例
var O orm.Ormer

// User 用户表
type User struct {
	UserID    int `orm:"column(user_id);pk"`
	Nickname  string
	Email     string
	Password  string
	AvatarURL string `orm:"column(avatar_url)"`
	InShort   string
	Introduce string
	QrImg     string
	Sex       string
	Birthday  time.Time
	FollowNum int
	FansNum   int
	CreatedAt time.Time  `orm:"auto_now_add;type(datetime)"`
	UpdatedAt time.Time  `orm:"auto_now;type(datetime)"`
	Article   []*Article `orm:"reverse(many)"`
	Comment   []*Comment `orm:"reverse(many)"`
}

// SystemSetting 设置表
type SystemSetting struct {
	ID    int `orm:"column(id);pk"`
	Key   string
	Name  string
	Value string
}

// Node 节点表
type Node struct {
	NodeID int `orm:"column(node_id);pk"`
	Name   string
	About  string
	Img    string
}

// Article 文章表
type Article struct {
	ArticleID  int   `orm:"column(article_id);pk"`
	User       *User `orm:"rel(fk)"`
	Title      string
	Content    string
	FaceImg    string
	FaceText   string
	Node       *Node `orm:"rel(one)"`
	Tag        string
	ReadNum    int
	LikeNum    int
	CollectNum int
	CommentNum int
	Status     int
	CreatedAt  time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt  time.Time `orm:"auto_now;type(datetime)"`
}

// ArticleLike 文章点赞表
type ArticleLike struct {
	ID        int       `orm:"column(id);pk"`
	ArticleID int       `orm:"column(article_id)"`
	UserID    int       `orm:"column(user_id)"`
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt time.Time `orm:"auto_now;type(datetime)"`
}

// ArticleCollect 文章收藏表
type ArticleCollect struct {
	ID        int       `orm:"column(id);pk"`
	ArticleID int       `orm:"column(article_id)"`
	UserID    int       `orm:"column(user_id)"`
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt time.Time `orm:"auto_now;type(datetime)"`
}

// Comment 评论表
type Comment struct {
	CommentID int      `orm:"column(comment_id);pk"`
	Article   *Article `orm:"rel(fk)"`
	User      *User    `orm:"rel(fk)"`
	Content   string
	CreatedAt time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt time.Time `orm:"auto_now;type(datetime)"`
}

// UserFollow 用户关注表
type UserFollow struct {
	ID           int       `orm:"column(id);pk"`
	UserID       int       `orm:"column(user_id)"`
	FollowUserID int       `orm:"column(follow_user_id)"`
	CreatedAt    time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt    time.Time `orm:"auto_now;type(datetime)"`
}

// Notice 通知表
type Notice struct {
	NoticeID   int `orm:"column(notice_id);pk"`
	UserID     int `orm:"column(user_id)"`
	Content    string
	ReadStatus int
	CreatedAt  time.Time `orm:"auto_now_add;type(datetime)"`
	UpdatedAt  time.Time `orm:"auto_now;type(datetime)"`
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
		new(Article),
		new(ArticleLike),
		new(ArticleCollect),
		new(Comment),
		new(UserFollow),
		new(Notice))

	orm.Debug = false
	O = orm.NewOrm()
}
