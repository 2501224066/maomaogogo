package models

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/astaxie/beego"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func ArticleRead(id int) (Article, error) {
	var article Article
	err := O.QueryTable(new(Article)).Filter("article_id", id).RelatedSel().One(&article)

	return article, err
}

func ArticleInsert(uid int, title, content, node, tag string) bool {
	// 去除两端空格
	strings.TrimSpace(title)
	strings.TrimSpace(content)
	strings.TrimSpace(tag)

	user := UserRead(uid)

	article := Article{
		User:     &user,
		Title:    title,
		Content:  content,
		FaceImg:  FaceImg(content),
		FaceText: FaceText(content),
		Node:     node,
		Tag:      tag,
		Status:   1}

	_, err := O.Insert(&article)
	if err != nil {
		fmt.Println(err)
		return false
	}

	return true
}

func ArticleUpdate(ArticleId int, title, content, node, tag string) bool {
	// 去除两端空格
	strings.TrimSpace(title)
	strings.TrimSpace(content)
	strings.TrimSpace(tag)

	_, err := O.QueryTable(new(Article)).Filter("article_id", ArticleId).Update(orm.Params{
		"title":     title,
		"content":   content,
		"face_img":  FaceImg(content),
		"face_text": FaceText(content),
		"node":      node,
		"tag":       tag})
	if err != nil {
		return false
	}

	return true
}

func FaceImg(content string) string {
	face_img := ""

	re, _ := regexp.Compile("<img src=\"(.*)\" style")
	arr := re.FindStringSubmatch(content)
	if len(arr) > 1 {
		face_img = arr[1]
	}

	return face_img
}

func FaceText(content string) string {
	// 将HTML标签全转换成小写
	re, _ := regexp.Compile("\\<[\\S\\s]+?\\>")
	content = re.ReplaceAllStringFunc(content, strings.ToLower)

	// 去除STYLE
	re, _ = regexp.Compile("\\<style[\\S\\s]+?\\</style\\>")
	content = re.ReplaceAllString(content, "")

	// 去除SCRIPT
	re, _ = regexp.Compile("\\<script[\\S\\s]+?\\</script\\>")
	content = re.ReplaceAllString(content, "")

	// 去除所有尖括号内的HTML代码
	re, _ = regexp.Compile("\\<[\\S\\s]+?\\>")
	content = re.ReplaceAllString(content, "")

	// 去除连续的换行符
	re, _ = regexp.Compile("\\s{2,}")
	content = re.ReplaceAllString(content, "")

	return content
}

// 文章总数量
func ArticleCount(nodeId, uid int) int64 {
	query := O.QueryTable(new(Article))

	if nodeId != 0 {
		var node Node
		O.QueryTable(new(Node)).Filter("node_id", nodeId).One(&node)
		query = query.Filter("node", node.Name)
	}

	if uid != 0 {
		query = query.Filter("User", uid)
	}

	count, _ := query.Filter("status", 1).Count()
	return count
}

// 文章列表
func ArticleList(nodeId, uid, p int) []Article {
	var article []Article
	query := O.QueryTable(new(Article))

	if nodeId != 0 {
		var node Node
		O.QueryTable(new(Node)).Filter("node_id", nodeId).One(&node)
		query = query.Filter("node", node.Name)
	}

	if uid != 0 {
		query = query.Filter("User", uid)
	}

	if p != 0 {
		pageNum, _ := beego.AppConfig.Int("page::num")
		offset := pageNum * (p - 1)
		query = query.Limit(pageNum, offset)
	}

	query.Filter("status", 1).OrderBy("-created_at").RelatedSel().All(&article)

	return article
}

// 文章阅读量增加
func ArticleReadNumUp(article Article) {
	article.ReadNum = article.ReadNum + 1
	O.Update(&article, "read_num")
}

// 文章删除
func ArticleDel(article_id int) bool {
	_, err := O.QueryTable(new(Article)).Filter("article_id", article_id).Update(orm.Params{"status": 0})
	if err != nil {
		return false
	}

	return true
}
