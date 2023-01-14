package model

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	gorm.Model
	Phone    string    `gorm:"type:varchar(20);not null"`  //手机号
	Password string    `gorm:"type:varchar(50);not null"`  //密码
	UserName string    `gorm:"type:varchar(50);not null"`  //用户名
	Email    string    `gorm:"type:varchar(100);not null"` //邮箱
	NickName string    `gorm:"type:varchar(50)"`           //用户昵称
	Avatar   string    `gorm:"type:varchar(255);not null"` //头像
	Gender   int       `gorm:"not null"`                   //性别
	Birthday time.Time //生日
	Address  string    `gorm:"type:varchar(255);not null"` //地址
	Status   int       `gorm:"not null"`                   //用户状态(是否激活，是否禁用等)
	Role     int       `gorm:"not null"`                   //角色
}

func (t User) TableName() string {
	return "user"
}

type LoginRecord struct {
	gorm.Model
	UserId    uint      `gorm:"not null"` //用户表ID
	LoginTime time.Time `gorm:"not null"` //最后登陆时间
}

//func (t LoginRecord) TableName() string {
//
//}
