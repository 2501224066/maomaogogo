package models

import "github.com/astaxie/beego/orm"

func ArticleLikeCount(article_id, uid int) int64 {
	query := O.QueryTable(new(ArticleLike))

	if uid != 0 {
		query = query.Filter("uid", uid)
	}

	count, _ := query.Filter("article_id", article_id).Count()
	return count
}

func ArticleLikeUp(article_id, uid int) bool {
	articleLike := ArticleLike{
		ArticleId: article_id,
		Uid:       uid}
	_, e := O.Insert(&articleLike)
	if e != nil {
		return false
	}

	var article Article
	O.QueryTable(new(Article)).Filter("article_id", article_id).One(&article)
	_, err := O.QueryTable(new(Article)).Filter("article_id", article_id).Update(orm.Params{"like_num": article.LikeNum + 1})
	if err != nil {
		return false
	}

	return true
}

func ArticleLikeDown(article_id, uid int) bool {
	_, e := O.QueryTable(new(ArticleLike)).Filter("article_id", article_id).Filter("uid", uid).Delete()
	if e != nil {
		return false
	}

	var article Article
	O.QueryTable(new(Article)).Filter("article_id", article_id).One(&article)
	_, err := O.QueryTable(new(Article)).Filter("article_id", article_id).Update(orm.Params{"like_num": article.LikeNum - 1})
	if err != nil {
		return false
	}

	return true
}
