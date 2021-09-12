package model

import (
	"time"
	"unique/jedi/common"

	"gorm.io/gorm"
)

// Basic user infomation
type User struct {
	CreateAt time.Time      `json:"-"`
	UpdateAt time.Time      `json:"-"`
	DeleteAt gorm.DeletedAt `json:"-"`

	UID          string          `json:"uid" gorm:"column:uid;primaryKey"`
	WorkwxUserId string          `json:"-" gorm:"column:workwx_user_id;index"`
	Name         string          `json:"name" gorm:"column:name"`
	Phone        string          `json:"phone" gorm:"column:phone;index"`
	EMail        string          `json:"email" gorm:"column:email;index"`
	Password     string          `json:"-" gorm:"column:password"`
	Role         common.UserRole `json:"role" gorm:"column:role"`
}

func (u *User) TableName() string {
	return "user"
}
