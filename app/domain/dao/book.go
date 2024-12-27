package dao

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title         string    `gorm:"size:56;not null;"`
	Subtitle      *string   `gorm:"size:46;"`
	PublisherID   uint      `gorm:"not null;"`
	BookPublisher Publisher `gorm:"foreignKey:PublisherID;references:ID;"`
	AuthorsID     uint      `gorm:"not null;"`
	BookAuthor    Authors   `gorm:"foreignKey:AuthorsID;references:ID;"`
}
