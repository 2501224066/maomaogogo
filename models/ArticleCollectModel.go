package models

import "github.com/astaxie/beego/orm"

func ArticleCollectCount(article_id, uid int) int64 {
	query := O.QueryTable(new(ArticleCollect))

	if article_id != 0 {
		query = query.Filter("article_id", article_id)
	}

	if uid != 0 {
		query = query.Filter("uid", uid)
	}

	count, _ := query.Count()
	return count
}

// 收藏增加
func ArticleCollectUp(article_id, uid int) bool {
	articleCollect := ArticleCollect{
		ArticleId: article_id,
		Uid:       uid}
	_, e := O.Insert(&articleCollect)
	if e != nil {
		return false
	}

	var article Article
	O.QueryTable(new(Article)).Filter("article_id", article_id).One(&article)
	_, err := O.QueryTable(new(Article)).Filter("article_id", article_id).Update(orm.Params{"collect_num": article.CollectNum + 1})
	if err != nil {
		return false
	}

	return true
}

// 收藏减少
func ArticleCollectDown(article_id, uid int) bool {
	_, e := O.QueryTable(new(ArticleCollect)).Filter("article_id", article_id).Filter("uid", uid).Delete()
	if e != nil {
		return false
	}

	var article Article
	O.QueryTable(new(Article)).Filter("article_id", article_id).One(&article)
	_, err := O.QueryTable(new(Article)).Filter("article_id", article_id).Update(orm.Params{"collect_num": article.CollectNum - 1})
	if err != nil {
		return false
	}

	return true
}
