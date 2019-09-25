package models

import (
	"maomaogogo/helper"
	"time"

	"github.com/astaxie/beego/orm"

	// mysql使用
	_ "github.com/go-sql-driver/mysql"
)

// UserRead 用户信息
func UserRead(userID int) User {
	user := User{UserID: userID}
	O.Read(&user)
	return user
}

// UserInsert 添加用户
func UserInsert(nickname, email, password string) bool {
	systemSetting := SystemSettingRead("default_avatar_url")
	avatarURL := systemSetting.Value

	local, _ := time.LoadLocation("Local")
	birthdayTime, _ := time.ParseInLocation("2006-01-02", "1996-05-15", local)

	user := User{
		Nickname:  nickname,
		Email:     email,
		Password:  password,
		AvatarURL: avatarURL,
		InShort:   "我永远爱猫，天地良心",
		Introduce: "可爱的铲屎官来到人间，上帝说：要有猫，于是便有了猫。铲屎官顿时欢呼雀跃，奔走相告",
		Sex:       " ",
		Birthday:  birthdayTime,
	}

	_, err := O.Insert(&user)
	if err != nil {
		return false
	}

	return true
}

// UserUpdate 修改个人信息
func UserUpdate(userID int, nickname, inShort, sex, birthday, introduce string) bool {
	local, _ := time.LoadLocation("Local")
	birthdayTime, _ := time.ParseInLocation("2006-01-02", birthday, local)

	_, err := O.QueryTable(new(User)).Filter("user_id", userID).Update(orm.Params{
		"nickname":  nickname,
		"in_short":  inShort,
		"sex":       sex,
		"birthday":  birthdayTime,
		"introduce": introduce})
	if err != nil {
		return false
	}

	return true
}

// AvatarURLUpdate 修改头像
func AvatarURLUpdate(userID int, avatarURL string) bool {
	_, err := O.QueryTable(new(User)).Filter("user_id", userID).Update(orm.Params{
		"avatar_url": avatarURL})
	if err != nil {
		return false
	}

	return true
}

// QrImgUpdate 修改打赏二维码
func QrImgUpdate(userID int, qrImg string) bool {
	_, err := O.QueryTable(new(User)).Filter("user_id", userID).Update(orm.Params{
		"qr_img": qrImg})
	if err != nil {
		return false
	}

	return true
}

// GetEmailCount 查询邮箱数量
func GetEmailCount(email string) int64 {
	cnt, _ := O.QueryTable(new(User)).Filter("email", email).Count()

	return cnt
}

// CheckLogin 检查登录
func CheckLogin(email, password string) (bool, int) {
	var user User
	err := O.QueryTable(new(User)).Filter("email", email).One(&user)

	if err == orm.ErrNoRows {
		return false, 0
	}
	if user.Password != helper.Md5(password) {
		return false, 0
	}

	return true, user.UserID
}

// ArticleData 文章数据结构体
type ArticleData struct {
	SumReadNum    int
	SumLikeNum    int
	SumCollectNum int
	SumCommentNum int
}

// AllArticleDataCount 用户所有文章数据统计
func AllArticleDataCount(userID int) ArticleData {
	var data ArticleData
	O.Raw("SELECT sum(read_num) as sum_read_num, sum(like_num) as sum_like_num, sum(collect_num) as sum_collect_num,sum(comment_num) as sum_comment_num FROM article WHERE user_id = ?", userID).QueryRow(&data)
	return data
}
