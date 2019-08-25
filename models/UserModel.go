package models

import (
	"fmt"
	"maomaogogo/helper"
	"time"

	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func UserRead(uid int) User {
	user := User{Uid: uid}
	O.Read(&user)

	return user
}

// 添加用户
func UserInsert(nickname, email, password string) bool {
	var systemSetting SystemSetting
	O.QueryTable("system_setting").Filter("key", "default_avatar_url").One(&systemSetting, "value")
	avatarUrl := systemSetting.Value

	local, _ := time.LoadLocation("Local")
	birthdayTime, _ := time.ParseInLocation("2006-01-02", "1996-05-15", local)

	fmt.Println(birthdayTime)
	user := User{
		Nickname:  nickname,
		Email:     email,
		Password:  password,
		AvatarUrl: avatarUrl,
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

// 修改个人信息
func UserUpdate(uid int, nickname, in_short, sex, birthday, introduce string) bool {
	local, _ := time.LoadLocation("Local")
	birthdayTime, _ := time.ParseInLocation("2006-01-02", birthday, local)

	_, err := O.QueryTable(new(User)).Filter("uid", uid).Update(orm.Params{
		"nickname":  nickname,
		"in_short":  in_short,
		"sex":       sex,
		"birthday":  birthdayTime,
		"introduce": introduce})
	if err != nil {
		return false
	}

	return true
}

// 修改头像
func AvatarUrlUpdate(uid int, avatar_url string) bool {
	_, err := O.QueryTable(new(User)).Filter("uid", uid).Update(orm.Params{
		"avatar_url": avatar_url})
	if err != nil {
		return false
	}

	return true
}

// 修改打赏二维码
func QrImgUpdate(uid int, qr_img string) bool {
	_, err := O.QueryTable(new(User)).Filter("uid", uid).Update(orm.Params{
		"qr_img": qr_img})
	if err != nil {
		return false
	}

	return true
}

// 查询邮箱数量
func GetEmailCount(email string) int64 {
	cnt, _ := O.QueryTable(new(User)).Filter("email", email).Count()

	return cnt
}

// 检查登录
func CheckLogin(email, password string) (bool, int) {
	var user User
	err := O.QueryTable(new(User)).Filter("email", email).One(&user)
	if err == orm.ErrNoRows {
		return false, 0
	}
	if user.Password != helper.Md5(password) {
		return false, 0
	}

	return true, user.Uid
}
