package model

import (
	"gorm.io/gorm"
	"time"
)

type Group struct {
	Id             int32     `json:"id"`
	GroupName      string    `json:"group_name"`
	UserId         int32     `json:"user_id"`
	CreatorId      int32     `json:"creator_id"`
	GroupAvatar    string    `json:"group_avatar"`
	CurrentMembers int32     `json:"current_members"`
	MaxMembers     int32     `json:"max_members"`
	Notice         string    `json:"notice"`
	CreatedAt      time.Time `time_format:"2006-01-02 15:04:05" json:"created_at"`
	UpdatedAt      time.Time `time_format:"2006-01-02 15:04:05" json:"updated_at"`
}

func (g *Group) BeforeCreate(db *gorm.DB) error {
	g.CreatedAt = time.Now()
	g.UpdatedAt = time.Now()
	return nil
}
