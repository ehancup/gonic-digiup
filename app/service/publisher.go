package service

import (
	"base-gin/app/domain/dto"
	"base-gin/app/repository"
	"base-gin/exception"
)

type PublisherService struct {
	repo *repository.PublisherRepository
}

func newPublisherService(publisherRepo *repository.PublisherRepository) *PublisherService {
	return &PublisherService{repo: publisherRepo}
}

func (s *PublisherService) Create(params *dto.PublisherCreateReq) (*dto.PublisherCreateResp, error) {
	newItem := params.ToEntity()

	err := s.repo.Create(&newItem)
	if err != nil {
		return nil, err
	}

	var resp dto.PublisherCreateResp
	resp.FromEntity(&newItem)

	return &resp, nil
}

func (s *PublisherService) GetList() ([]dto.PublisherDetailResp, error) {
	var resp []dto.PublisherDetailResp

	items, err := s.repo.List()
	if err != nil {
		return nil, err
	}
	if len(items) < 1 {
		return nil, exception.ErrUserNotFound
	}

	for _, item := range items {
		var t dto.PublisherDetailResp
		t.FromEntity(&item)

		resp = append(resp, t)
	}

	return resp, nil
}

func (s *PublisherService) Update(params *dto.PublisherUpdateReq) error {
	if params.ID <= 0 {
		return exception.ErrUserNotFound
	}

	// birthDate, err := params.GetBirthDate()
	// if err != nil {
	// 	exception.LogError(err, "PersonService.Update")
	// 	return exception.ErrDateParsing
	// }
	// params.BirthDate = birthDate

	return s.repo.Update(params)
}
func (s *PublisherService) Delete(id uint) error {
	if id <= 0 {
		return exception.ErrUserNotFound
	}

	// birthDate, err := params.GetBirthDate()
	// if err != nil {
	// 	exception.LogError(err, "PersonService.Update")
	// 	return exception.ErrDateParsing
	// }
	// params.BirthDate = birthDate

	return s.repo.Delete(id)
}



func (s *PublisherService) GetByID(id uint) (dto.PublisherDetailResp, error) {
	var resp dto.PublisherDetailResp

	item, err := s.repo.GetByID(id)
	if err != nil {
		return resp, err
	}
	if item == nil {
		return resp, exception.ErrUserNotFound
	}

	resp.FromEntity(item)

	return resp, nil
}