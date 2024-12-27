package dao

import (
	"base-gin/app/domain"

	"gorm.io/gorm"
)

type FormTable struct {
	gorm.Model
	GolonganDarah   string            `gorm:"size:255;not null;table:golongan_darah;"`
	MakananKesukaan string            `gorm:"size:255;not null;table:makanan_kesukaan;"`
	MinumanKesukaan string            `gorm:"size:255;not null;table:minuman_kesukaan;"`
	Nama            string            `gorm:"size:255;not null;table:nama;"`
	Gender          domain.TypeGender `gorm:"type:enum('f','m');table:gender;"`
}

func (FormTable) TableName() string {
	return "form"
}

