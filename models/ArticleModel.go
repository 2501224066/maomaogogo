package models

import (
	"regexp"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"

	// mysql 使用
	_ "github.com/go-sql-driver/mysql"
)

// ArticleRead 单条文章信息
func ArticleRead(articleID int) (Article, error) {
	var article Article
	err := O.QueryTable(new(Article)).Filter("article_id", articleID).RelatedSel().One(&article)

	return article, err
}

// ArticleInsert 查询文章
func ArticleInsert(userID, nodeID int, title, content, tag string) bool {
	// 去除两端空格
	strings.TrimSpace(title)
	strings.TrimSpace(content)
	strings.TrimSpace(tag)

	user := UserRead(userID)
	node := NodeRead(nodeID)

	article := Article{
		User:     &user,
		Title:    title,
		Content:  content,
		FaceImg:  FaceImg(content),
		FaceText: FaceText(content),
		Node:     &node,
		Tag:      tag,
		Status:   1}

	_, err := O.Insert(&article)
	if err != nil {
		return false
	}

	return true
}

// ArticleUpdate 修改文章
func ArticleUpdate(ArticleID int, title, content, node, tag string) bool {
	// 去除两端空格
	strings.TrimSpace(title)
	strings.TrimSpace(content)
	strings.TrimSpace(tag)

	_, err := O.QueryTable(new(Article)).Filter("article_id", ArticleID).Update(orm.Params{
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

// FaceImg 封面图片
func FaceImg(content string) string {
	faceImg := ""

	re, _ := regexp.Compile("<img src=\"(.*)\" style")
	arr := re.FindStringSubmatch(content)
	if len(arr) > 1 {
		faceImg = arr[1]
	}

	return faceImg
}

// FaceText 封面文本
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

// ArticleCount 文章总数量
func ArticleCount(nodeID, userID int) int64 {
	query := O.QueryTable(new(Article))

	if nodeID != 0 {
		query = query.Filter("node_id", nodeID)
	}

	if userID != 0 {
		query = query.Filter("User", userID)
	}

	count, _ := query.Filter("status", 1).Count()
	return count
}

// ArticleList 文章列表
func ArticleList(nodeID, userID, p int) []Article {
	var article []Article
	query := O.QueryTable(new(Article))

	if nodeID != 0 {
		query = query.Filter("node_id", nodeID)
	}

	if userID != 0 {
		query = query.Filter("User", userID)
	}

	if p != 0 {
		pageNum, _ := beego.AppConfig.Int("page::num")
		offset := pageNum * (p - 1)
		query = query.Limit(pageNum, offset)
	}

	query.Filter("status", 1).OrderBy("-created_at").RelatedSel().All(&article)
	return article
}

// ArticleReadNumUp 文章阅读量增加
func ArticleReadNumUp(article Article) {
	article.ReadNum = article.ReadNum + 1
	O.Update(&article, "read_num")
}

// ArticleDel 文章删除
func ArticleDel(articleID int) bool {
	_, err := O.QueryTable(new(Article)).Filter("article_id", articleID).Update(orm.Params{"status": 0})
	if err != nil {
		return false
	}

	return true
}

// ArticleCommentNumUp 文章评论数增加
func ArticleCommentNumUp(articleID int) {
	article, _ := ArticleRead(articleID)
	article.CommentNum = article.CommentNum + 1
	O.Update(&article, "comment_num")
}