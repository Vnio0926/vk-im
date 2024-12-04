package service

import (
	"errors"
	"gorm.io/gorm"
	"vk-im/internal/db"
	"vk-im/internal/model"
	"vk-im/util/common/request"
	"vk-im/util/common/response"
	"vk-im/util/protocol"
)

const (
	NULL_ID            int32 = 0
	MESSAGE_TYPE_USER        = 1 //单聊1
	MESSAGE_TYPE_GROUP       = 2 //群聊2
)

type messageService struct{}

var MessageService = new(messageService)

func (m *messageService) GetMessage(message request.MessageRequest) ([]response.MessageResponse, error) {
	_db := db.GetDb()
	_db.AutoMigrate(&model.Message{})

	if message.MessageType == MESSAGE_TYPE_USER {
		var queryUser *model.User
		_db.First(&queryUser, "id = ?", message.FromId)
		if queryUser.Id == NULL_ID {
			return nil, errors.New("用户不存在")
		}

		var friendUser *model.User
		_db.First(&friendUser, "id = ?", message.ToId)
		if friendUser.Id == NULL_ID {
			return nil, errors.New("用户不存在")
		}

		var messages []response.MessageResponse
		_db.Raw("messages.content,messages.content_type,messages.created_at,messages.from_id,messages.id,messages.to_id,"+
			"messages.url,users.avatar,users.username").
			Joins("left join users on messages.from_id = users.id").
			Where("messages.from_id = ? and messages.to_id = ? or messages.from_id = ? and messages.to_id = ?",
				message.FromId, message.ToId, message.ToId, message.FromId).
			Order("messages.created_at asc").
			Scan(&messages)
		return messages, nil
	}

	if message.MessageType == MESSAGE_TYPE_GROUP {
		message, err := fetchGroupMessage(_db, message.Uuid)
		if err != nil {
			return nil, err
		}
		return message, nil
	}

	return nil, errors.New("暂不支持")

}

func fetchGroupMessage(db *gorm.DB, toUuid string) ([]response.MessageResponse, error) {
	var group *model.Group
	db.First(&group, "uuid = ?", toUuid)
	if group.Id <= NULL_ID {
		return nil, errors.New("群不存在")
	}

	var messages []response.MessageResponse
	db.Raw("messages.content,messages.content_type,messages.created_at,messages.from_id,messages.id,messages.to_id,"+
		"messages.url,users.avatar,users.username").
		Joins("left join users on messages.from_id = users.id").
		Where("messages.to_id = ?", group.Id).
		Order("messages.created_at asc").
		Scan(&messages)
	return messages, nil
}

func (m *messageService) SaveMessage(message protocol.Message) {
	_db := db.GetDb()
	var fromUser *model.User
	_db.Find(&fromUser, "id = ?", message.From)
	if fromUser.Id == NULL_ID {
		return
	}
	var toUserId int32 = 0
	if message.MsgType == MESSAGE_TYPE_USER {
		var toUser *model.User
		_db.Find(&toUser, "id = ?", message.To)
		if toUser.Id == NULL_ID {
			return
		}
		toUserId = toUser.Id
	}

	if message.MsgType == MESSAGE_TYPE_GROUP {
		var group *model.Group
		_db.Find(&group, "uuid = ?", message.To)
		if group.Id == NULL_ID {
			return
		}
		toUserId = group.Id
	}
	saveMessage := model.Message{
		FromId:      fromUser.Id,
		ToId:        toUserId,
		Content:     message.Content,
		ContentType: int16(message.ContentType),
		MessageType: int16(message.MsgType),
		Url:         message.Url,
	}
	_db.Save(&saveMessage)
}
