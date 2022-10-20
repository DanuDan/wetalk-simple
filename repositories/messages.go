package repositories

import (
	"wetalk/models"

	"gorm.io/gorm"
)

type MessageRepository interface {
	FindMessages(UserID int) ([]models.Message, error)
	GetMessage(ID int) (models.Message, error)
	CreateMessage(message models.Message) (models.Message, error)
	DeleteMessage(message models.Message) (models.Message, error)
}

func RepositoryMessage(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindMessages(UserID int) ([]models.Message, error) {
	var messages []models.Message
	err := r.db.Preload("User").Find(&messages, "user_id = ?", UserID).Error

	return messages, err
}

func (r *repository) GetMessage(ID int) (models.Message, error) {
	var message models.Message
	err := r.db.Preload("User").First(&message, ID).Error

	return message, err
}

func (r *repository) CreateMessage(message models.Message) (models.Message, error) {
	err := r.db.Preload("User").Create(&message).Error

	return message, err
}

func (r *repository) DeleteMessage(message models.Message) (models.Message, error) {
	err := r.db.Preload("User").Delete(&message).Error

	return message, err
}
