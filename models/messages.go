package models

import "time"

type Message struct {
	ID        int                 `json:"id" gorm:"primary_key:auto_increment"`
	Message   string              `json:"message" gorm:"type: varchar(255)"`
	UserID    int                 `json:"user_id"`
	User      UserMessageResponse `json:"user"`
	CreatedAt time.Time           `json:"-"`
	UpdatedAt time.Time           `json:"-"`
}

type MessageUserResponse struct {
	Message string `json:"message"`
}

func (MessageUserResponse) TableName() string {
	return "messages"
}
