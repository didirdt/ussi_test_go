package models

import "gorm.io/gorm"

type Note struct {
	gorm.Model
	Title   string `gorm:"type:varchar(255);not null" json:"title"`
	Content string `gorm:"type:text" json:"content"`
	UserID  uint   `gorm:"not null;index" json:"user_id"`
	User    User   `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;" json:"-"`
}

type CreateNoteRequest struct {
	Title   string `json:"title" binding:"required,max=255"`
	Content string `json:"content"`
}

type UpdateNoteRequest struct {
	Title   string `json:"title" binding:"max=255"`
	Content string `json:"content"`
}
