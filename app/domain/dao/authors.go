package dao

import (
	"base-gin/app/domain"
	"time"

	"gorm.io/gorm"
)

type Authors struct {
	gorm.Model
	Fullname  string             `gorm:"size:56;not null;"`
	Gender    *domain.TypeGender `gorm:"type:enum('f','m');"`
	BirthDate *time.Time
	Book      []Book 
}

func (Authors) TableName() string {
	return "authors"
}
