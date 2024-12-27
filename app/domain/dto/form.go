package dto

import (
	"base-gin/app/domain"
	"base-gin/app/domain/dao"
)

type FormPostReq struct {
	Nama            string `json:"nama" binding:"required,max=48,min=6"`
	GolonganDarah   string `json:"gol_darah" binding:"required,max=48,min=1"`
	MakananKesukaan string `json:"makanan_kesukaan" binding:"required,max=32,min=2"`
	MinumanKesukaan string `json:"minuman_kesukaan" binding:"required,max=32,min=2"`
	Gender          string `json:"gender" binding:"required,oneof=m f"`
}

func (v FormPostReq) ToEntity() dao.FormTable {

	var gender domain.TypeGender
	switch v.Gender {
	case "f":
		gender = domain.GenderFemale
	case "m":
		gender = domain.GenderMale
	}
	return dao.FormTable{
		Nama:            v.Nama,
		GolonganDarah:   v.GolonganDarah,
		MakananKesukaan: v.MakananKesukaan,
		MinumanKesukaan:  v.MinumanKesukaan,
		Gender:          gender,
	}
}
