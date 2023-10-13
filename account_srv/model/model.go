package model

import "gorm.io/gorm"

type Account struct {
	gorm.Model
	Mobile   string `gorm:"index:idx_mobile;unique;varchar(255);not null"`
	Password string `gorm:"type:varchar(64);not null"`
	NickName string `gorm:"type:varchar(32)"`
	Gender   string `gorm:"type:varchar(10);default:'male';comment:'male-男性, female-女性'"`
	Role     int    `gorm:"type:tinyint unsigned;default:1;comment:'1-普通用户, 2-管理员'"`
	Salt     string `gorm:"type:varchar(16)"`
}
