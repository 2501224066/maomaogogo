package models

import "github.com/astaxie/beego/orm"

// ArticleCollectCount 收藏数量
func ArticleCollectCount(articleID, userID int) int64 {
	query := O.QueryTable(new(ArticleCollect))

	if articleID != 0 {
		query = query.Filter("article_id", articleID)
	}

	if userID != 0 {
		query = query.Filter("user_id", userID)
	}

	count, _ := query.Count()
	return count
}

// ArticleCollectUp 收藏增加
func ArticleCollectUp(articleID, userID int) bool {
	articleCollect := ArticleCollect{
		ArticleID: articleID,
		UserID:    userID}
	_, e := O.Insert(&articleCollect)
	if e != nil {
		return false
	}

	var article Article
	O.QueryTable(new(Article)).Filter("article_id", articleID).One(&article)
	_, err := O.QueryTable(new(Article)).Filter("article_id", articleID).Update(orm.Params{"collect_num": article.CollectNum + 1})
	if err != nil {
		return false
	}

	return true
}

// ArticleCollectDown 收藏减少
func ArticleCollectDown(articleID, userID int) bool {
	_, e := O.QueryTable(new(ArticleCollect)).Filter("article_id", articleID).Filter("user_id", userID).Delete()
	if e != nil {
		return false
	}

	var article Article
	O.QueryTable(new(Article)).Filter("article_id", articleID).One(&article)
	_, err := O.QueryTable(new(Article)).Filter("article_id", articleID).Update(orm.Params{"collect_num": article.CollectNum - 1})
	if err != nil {
		return false
	}

	return true
}



