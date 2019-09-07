package models

import (
	"strings"
)

// CommentList 评论列表
func CommentList(articleID int) interface{} {
	var comment []*Comment
	O.QueryTable(new(Comment)).Filter("article_id", articleID).RelatedSel().All(&comment)
	return comment
}

// CommentInsert 添加评论
func CommentInsert(articleID, userID int, content string) bool {
	// 去除两端空格
	strings.TrimSpace(content)

	user := UserRead(userID)
	article, _ := ArticleRead(articleID)

	comment := Comment{
		User:    &user,
		Article: &article,
		Content: content}

	_, err := O.Insert(&comment)
	if err != nil {
		return false
	}

	return true
}
