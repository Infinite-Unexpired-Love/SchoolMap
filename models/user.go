package models

import (
	"time"
)

type User struct {
	ID         uint32     `gorm:"primaryKey"`
	Username   string     `gorm:"type:varchar(60) not null"`
	Mobile     string     `gorm:"index:idx_mobile;unique;type:varchar(11);not null"`
	Password   string     `gorm:"type:varchar(100);not null"`
	Role       int        `gorm:"default:1;type:int comment '1表示普通管理员，2表示超级管理员'"`
	Expiration *time.Time `gorm:"type:datetime"`
}

func (u *User) GetID() uint {
	return uint(u.ID)
}

func (u *User) GetTitle() string {
	return u.Mobile
}

func (u *User) Update(target interface{}) {
	updateStructFields(u, target)
}
