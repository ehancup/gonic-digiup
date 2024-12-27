package service

import (
	"base-gin/app/domain/dto"
	"base-gin/app/repository"
)

type FormService struct {
	repo *repository.FormRepository
}

func newFormService(formRepo *repository.FormRepository) *FormService {
	return &FormService{repo: formRepo}
}

func (s *FormService) Create(params *dto.FormPostReq) (error) {
	newItem := params.ToEntity()

	err := s.repo.Create(&newItem)
	// if err != nil {
	// 	return nil, err
	// }

	// var resp dto.PublisherCreateResp
	// resp.FromEntity(&newItem)

	return err
}