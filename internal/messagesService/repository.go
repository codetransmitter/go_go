package messagesService

import (
	"gorm.io/gorm"
)

type MessageRepository interface {
	CreateMessage(message Message) (Message, error)

	GetAllMessages() ([]Message, error)

	UpdateMessageByID(id uint, message Message) (Message, error)

	DeleteMessageByID(id uint) error
}

type messageRepository struct {
	db *gorm.DB
}

func NewMessageRepository(db *gorm.DB) *messageRepository {
	return &messageRepository{db: db}
}

func (r *messageRepository) CreateMessage(message Message) (Message, error) {

	result := r.db.Create(&message)
	if result.Error != nil {
		return Message{}, result.Error
	}
	return message, nil
}
func (r *messageRepository) GetAllMessages() ([]Message, error) {

	var messages []Message
	err := r.db.Find(&messages).Error
	return messages, err
}
func (r *messageRepository) UpdateMessageByID(id uint, message Message) (Message, error) {
	//fmt.Println("MSG_REP", message.ID)
	result := r.db.Save(&message)
	if result.Error != nil {
		return Message{}, result.Error
	}
	return message, nil
}
func (r *messageRepository) DeleteMessageByID(id uint) error {
	var message Message
	message.ID = id
	result := r.db.Delete(&message)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
