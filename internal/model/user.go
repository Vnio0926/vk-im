package model

import (
	"gorm.io/gorm"
	"time"
	"vk-im/util/constant/enmu"
)

type User struct {
	Id          int32     `json:"id" `
	Username    string    `json:"username"`
	Password    string    `json:"password"`
	CreatedAt   time.Time `time_format:"2006-01-02 15:04:05" json:"created_at"`
	Description string    `json:"description"`
	Avatar      string    `json:"avatar"`
	//LastLogin   time.Time   `time_format:"2006-01-02 15:04:05" json:"last_login"`
	Email  string      `json:"email"`
	Sex    enmu.Sex    `json:"sex"`
	Status enmu.Status `json:"status"`
	Token  string      `json:"token"`
}

func (u *User) BeforeCreate(db *gorm.DB) error {
	u.CreatedAt = time.Now()
	return nil
}

func (u *User) BeforeUpdate(db *gorm.DB) error {
	db.Statement.SetColumn("updated_at", time.Now())
	return nil
}
