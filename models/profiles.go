package models

import "time"

type Profile struct {
	ID        int                 `json:"id" gorm:"primary_key:auto_increment"`
	Image     string              `json:"image" gorm:"type: varchar(255)"`
	UserID    int                 `json:"user_id"`
	User      UserProfileResponse `json:"user"`
	CreatedAt time.Time           `json:"-"`
	UpdatedAt time.Time           `json:"-"`
}

type ProfileUserResponse struct {
	Image string `json:"image"`
}

func (ProfileUserResponse) TableName() string {
	return "profiles"
}
