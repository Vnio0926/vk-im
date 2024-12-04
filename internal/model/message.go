package model

import (
	"gorm.io/gorm"
	"time"
)

type Message struct {
	Id          int32     `json:"id"`
	Content     string    `json:"content"`
	CreatedAt   time.Time `time_format:"2006-01-02 15:04:05" json:"created_at"`
	FromId      int32     `json:"from_id"`
	ToId        int32     `json:"to_id"`
	MessageType int16     `json:"messageType" gorm:"comment:'消息类型：1单聊，2群聊'"`
	ContentType int16     `json:"contentType" gorm:"comment:'消息内容类型：1文字 2.普通文件 3.图片 4.音频 5.视频 6.语音聊天 7.视频聊天'"`
	RoomId      int32     `json:"room_id"`
	Pic         string    `json:"pic" gorm:"type:text;comment:'缩略图"`
	Url         string    `json:"url" gorm:"type:varchar(350);comment:'文件或者图片地址'"`
}

func (m *Message) BeforeCreate(db *gorm.DB) error {
	m.CreatedAt = time.Now()
	return nil
}
