package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

// AddNotice 添加通知
func AddNotice(userID int, content string) {
	notice := Notice{
		UserID:     userID,
		Content:    content,
		ReadStatus: 1}

	O.Insert(&notice)
}

// UnreadNoticeCount 用户未读消息数量
func UnreadNoticeCount(userID, readStatus int) int64 {
	query := O.QueryTable(new(Notice))

	if readStatus != 0 {
		query = query.Filter("read_status", readStatus)
	}

	count, _ := query.Filter("user_id", userID).Count()
	return count
}

// NoticeList 通知列表
func NoticeList(userID, p int) []Notice {
	var notice []Notice
	query := O.QueryTable(new(Notice)).Filter("user_id", userID)

	if p != 0 {
		pageNum, _ := beego.AppConfig.Int("page::num")
		offset := pageNum * (p - 1)
		query = query.Limit(pageNum, offset)
	}

	query.OrderBy("-created_at").All(&notice)

	return notice
}

// NoticeAllRead 通知全部已读
func NoticeAllRead(userID int) {
	O.QueryTable(new(Notice)).Update(orm.Params{"read_status": 2})
}
