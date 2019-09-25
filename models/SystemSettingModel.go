package models

// SystemSettingRead 设置
func SystemSettingRead(key string) SystemSetting {
	var systemSetting SystemSetting
	O.QueryTable(new(SystemSetting)).Filter("key", key).One(&systemSetting)
	return systemSetting
}
