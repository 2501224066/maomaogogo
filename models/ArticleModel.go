package models

import (
	"fmt"
	"regexp"
	"strings"

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
		FaceImg:  GetFaceImg(content),
		FaceText: GetFaceText(content),
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

func GetFaceImg(content string) string {
	face_img := ""

	re, _ := regexp.Compile("<img src=\"(.*)\" style")
	arr := re.FindStringSubmatch(content)
	if len(arr) > 1 {
		face_img = arr[1]
	}

	return face_img
}

func GetFaceText(content string) string {
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

func GetArticleList() []Article {
	var article []Article
	O.QueryTable(new(Article)).Filter("status", 1).OrderBy("-created_at").RelatedSel().All(&article)

	return article
}

func GetSelfArticleList(uid int) []Article {
	var article []Article
	O.QueryTable(new(Article)).Filter("status", 1).Filter("User", uid).OrderBy("-created_at").RelatedSel().All(&article)

	return article
}

func ArticleReadAddOne(article Article) {
	article.ReadNum = article.ReadNum + 1
	O.Update(&article, "read_num")
}
