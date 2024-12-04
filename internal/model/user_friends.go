package model

import (
	"gorm.io/gorm"
	"time"
)

type UserFriends struct {
	Id         int32     `json:"id"`
	UserId     int32     `json:"user_id"`
	FriendId   int32     `json:"friend_id"`
	FriendName string    `json:"friend_name"`
	CreatedAt  time.Time `time_format:"2006-01-02 15:04:05" json:"created_at"`
	UpdatedAt  time.Time `time_format:"2006-01-02 15:04:05" json:"updated_at"`
}

func (u *UserFriends) BeforeCreate(db *gorm.DB) error {
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
	return nil
}
func (u *UserFriends) BeforeUpdate(db *gorm.DB) error {
	u.UpdatedAt = time.Now()
	return nil
}
