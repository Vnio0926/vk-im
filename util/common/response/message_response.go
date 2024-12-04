package response

import (
	"time"
)

type MessageResponse struct {
	ID           int32     `json:"id" gorm:"primary"`
	FromUserId   int32     `json:"fromUserId"`
	ToUserId     int32     `json:"toUserId"`
	Content      string    `json:"content"`
	ContentType  string    `json:"contentType"`
	CreatedAt    time.Time `time_format:"2006-01-02 15:04:05" json:"createdAt"`
	FromUsername string    `json:"fromUsername"`
	ToUsername   string    `json:"toUsername"`
	Avatar       string    `json:"avatar"`
	Url          string    `json:"url"`
	Seq          int32     `json:"seq"`
}
