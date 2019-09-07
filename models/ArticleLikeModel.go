package models

import "github.com/astaxie/beego/orm"

// ArticleLikeCount 点赞数量
func ArticleLikeCount(articleID, userID int) int64 {
	query := O.QueryTable(new(ArticleLike))

	if userID != 0 {
		query = query.Filter("user_id", userID)
	}

	count, _ := query.Filter("article_id", articleID).Count()
	return count
}

// ArticleLikeUp 点赞增加
func ArticleLikeUp(articleID, userID int) bool {
	articleLike := ArticleLike{
		ArticleID: articleID,
		UserID:    userID}
	_, e := O.Insert(&articleLike)
	if e != nil {
		return false
	}

	var article Article
	O.QueryTable(new(Article)).Filter("article_id", articleID).One(&article)
	_, err := O.QueryTable(new(Article)).Filter("article_id", articleID).Update(orm.Params{"like_num": article.LikeNum + 1})
	if err != nil {
		return false
	}

	return true
}

// ArticleLikeDown 点赞减少
func ArticleLikeDown(articleID, userID int) bool {
	_, e := O.QueryTable(new(ArticleLike)).Filter("article_id", articleID).Filter("user_id", userID).Delete()
	if e != nil {
		return false
	}

	var article Article
	O.QueryTable(new(Article)).Filter("article_id", articleID).One(&article)
	_, err := O.QueryTable(new(Article)).Filter("article_id", articleID).Update(orm.Params{"like_num": article.LikeNum - 1})
	if err != nil {
		return false
	}

	return true
}
