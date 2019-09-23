package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

// FollowCount 关注状态
func FollowCount(userID, followUserID int) int64 {
	query := O.QueryTable(new(UserFollow))
	count, _ := query.Filter("user_id", userID).Filter("follow_user_id", followUserID).Count()
	return count
}

// UserFollowUp 关注
func UserFollowUp(userID, followUserID int) bool {
	// 添加关注记录
	userFollow := UserFollow{
		UserID:       userID,
		FollowUserID: followUserID}
	_, e := O.Insert(&userFollow)
	if e != nil {
		return false
	}

	// 用户关注数上升
	var user User
	O.QueryTable(new(User)).Filter("user_id", userID).One(&user)
	_, err := O.QueryTable(new(User)).Filter("user_id", userID).Update(orm.Params{"follow_num": user.FollowNum + 1})
	if err != nil {
		return false
	}

	// 被关注着粉丝数上升
	var user2 User
	O.QueryTable(new(User)).Filter("user_id", followUserID).One(&user2)
	_, err2 := O.QueryTable(new(User)).Filter("user_id", followUserID).Update(orm.Params{"fans_num": user2.FansNum + 1})
	if err2 != nil {
		return false
	}

	return true
}

// UserFollowDown 取消关注
func UserFollowDown(userID, followUserID int) bool {
	// 删除关注记录
	_, e := O.QueryTable(new(UserFollow)).Filter("user_id", userID).Filter("follow_user_id", followUserID).Delete()
	if e != nil {
		return false
	}

	// 用户关注数下降
	var user User
	O.QueryTable(new(User)).Filter("user_id", userID).One(&user)
	_, err := O.QueryTable(new(User)).Filter("user_id", userID).Update(orm.Params{"follow_num": user.FollowNum - 1})
	if err != nil {
		return false
	}

	// 被关注着粉丝数下降
	var user2 User
	O.QueryTable(new(User)).Filter("user_id", followUserID).One(&user2)
	_, err2 := O.QueryTable(new(User)).Filter("user_id", followUserID).Update(orm.Params{"fans_num": user2.FansNum - 1})
	if err2 != nil {
		return false
	}

	return true
}

// FollowUserList 关注用户列表
func FollowUserList(userID, p int) []User {
	var userFollow []UserFollow
	query := O.QueryTable(new(UserFollow)).Filter("user_id", userID)

	if p != 0 {
		pageNum, _ := beego.AppConfig.Int("page::num")
		offset := pageNum * (p - 1)
		query = query.Limit(pageNum, offset)
	}

	query.OrderBy("-created_at").All(&userFollow)

	// 被关注人信息
	var (
		userFollowList []User
		user           User
	)
	for _, v := range userFollow {
		user.UserID = v.FollowUserID
		O.Read(&user)
		userFollowList = append(userFollowList, user)
	}

	return userFollowList
}

// FansList 粉丝列表
func FansList(userID, p int) []User {
	var userFollow []UserFollow
	query := O.QueryTable(new(UserFollow)).Filter("follow_user_id", userID)

	if p != 0 {
		pageNum, _ := beego.AppConfig.Int("page::num")
		offset := pageNum * (p - 1)
		query = query.Limit(pageNum, offset)
	}

	query.OrderBy("-created_at").All(&userFollow)

	// 粉丝信息
	var (
		FansList []User
		user     User
	)
	for _, v := range userFollow {
		user.UserID = v.UserID
		O.Read(&user)
		FansList = append(FansList, user)
	}

	return FansList
}
