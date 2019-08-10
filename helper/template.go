package helper

import (
	"strconv"
	"strings"
	"time"
)

// 路由对比
func IsActive(x, y string) bool {
	s1 := strings.Trim(x, "/")
	s2 := strings.Trim(y, "/")
	return strings.Compare(s1, s2) == 0
}

// 按空格拆分
func StrSplitForBlankSpace(str string) []string {
	return strings.Fields(str)
}

// 时间间隔
func TimeInterval(t string) string {
	timeLayout := "2006-01-02 15:04:05"
	loc, _ := time.LoadLocation("Local")
	theTime, _ := time.ParseInLocation(timeLayout, t, loc)
	var (
		putC int64 = theTime.Unix()
		nowC int64 = time.Now().Unix()
		i    int64 = nowC - putC
	)

	switch true {
	case i < 60:
		return strconv.FormatInt(i, 10) + "秒前"
	case i < 3600:
		return strconv.FormatInt(i/60, 10) + "分钟前"
	case i < 86400:
		return strconv.FormatInt(i/3600, 10) + "小时前"
	case i < 2592000:
		return strconv.FormatInt(i/86400, 10) + "天前"
	case i < 31536000:
		return strconv.FormatInt(i/2592000, 10) + "个月前"
	default:
		return strconv.FormatInt(i/31536000, 10) + "年前"
	}
}
