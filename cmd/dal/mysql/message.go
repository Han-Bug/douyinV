package mysql

import (
	"tiktok/biz/model"
)

func GetMessages(fromUID, toUID int) ([]model.Message, error) {
	var msgs []model.Message
	if err := DB.Where("(from_user_id = ? and to_user_id = ?) or (from_user_id = ? and to_user_id = ?)", fromUID, toUID, toUID, fromUID).Find(&msgs).Error; err != nil {
		return nil, err
	}
	return msgs, nil
}

func NewMessage(fromUID, toUID int, content string) error {
	msg := model.Message{
		FromUserId: fromUID,
		ToUserId:   toUID,
		Content:    content,
	}
	return DB.Create(&msg).Error
}
