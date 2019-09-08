package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

// FollowCount 关注状态
func FollowCount(userID, followUserID int) int64 {
	query := O.QueryTable(new(UserFollow))
	fmt.Println(userID)
	fmt.Println(followUserID)
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
