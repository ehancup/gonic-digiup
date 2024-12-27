package dto

import (
	"base-gin/app/domain"
	"base-gin/app/domain/dao"
	"time"
)

type AuthorsCreateReq struct {
	Fullname     string `json:"fullname" binding:"required,min=4,max=56"`
	Gender       string `json:"gender" binding:"required,oneof=m f"`
	BirthDateStr string `json:"birth_date" binding:"required,datetime=2006-01-02"`
}

func (v AuthorsCreateReq) ToEntity() dao.Authors {
	var gender domain.TypeGender
	switch v.Gender {
	case "f":
		gender = domain.GenderFemale
	case "m":
		gender = domain.GenderMale
	}

	birthDate, _ := time.Parse("2006-01-02", v.BirthDateStr)

	return dao.Authors{
		Fullname:  v.Fullname,
		Gender:    &gender,
		BirthDate: &birthDate,
	}
}

// func (o *AuthorsCreateReq) GetGender() domain.TypeGender {
// 	if o.Gender == "f" {
// 		return domain.GenderFemale
// 	}

// 	return domain.GenderMale
// }

// func (o *PersonUpdateReq) GetBirthDate() (time.Time, error) {
// 	return time.Parse("2006-01-02", o.BirthDateStr)
// }

type AuthorsCreateResp struct {
	ID           int    `json:"id"`
	Fullname     string `json:"fullname"`
	Gender       string `json:"gender"`
	BirthDateStr string `json:"birth_date"`
}

func (v *AuthorsCreateResp) FromEntity(item *dao.Authors) {
	v.ID = int(item.ID)
	v.Fullname = item.Fullname
	v.Gender = string(*item.Gender)
	v.BirthDateStr = item.BirthDate.Format("2006-01-02")
}

type AuthorsDetailResp struct {
	ID           int    `json:"id"`
	Fullname     string `json:"fullname"`
	Gender       string `json:"gender"`
	BirthDateStr string `json:"birth_date"`
}
func (v *AuthorsDetailResp) FromEntity(item *dao.Authors) {
	v.ID = int(item.ID)
	v.Fullname = item.Fullname
	v.Gender = string(*item.Gender)
	v.BirthDateStr = item.BirthDate.Format("2006-01-02")
}

type AuthorsUpdateReq struct {
	ID           uint      `json:"-"`
	Fullname     string    `json:"fullname" binding:"required,min=4,max=56"`
	Gender       string    `json:"gender" binding:"required,oneof=m f"`
	BirthDateStr string    `json:"birth_date" binding:"required,datetime=2006-01-02"`
	BirthDate    time.Time `json:"-"`
}

func (o *AuthorsUpdateReq) GetGender() domain.TypeGender {
	if o.Gender == "f" {
		return domain.GenderFemale
	}

	return domain.GenderMale
}

func (o *AuthorsUpdateReq) GetBirthDate() (time.Time, error) {
	return time.Parse("2006-01-02", o.BirthDateStr)
}