package service

import (
	"base-gin/app/domain/dto"
	"base-gin/app/repository"
	"base-gin/exception"
)

type AuthorsService struct {
	repo *repository.AuthorsRepository
}

func newAuthorsService(authorsRepo *repository.AuthorsRepository) *AuthorsService {
	return &AuthorsService{repo: authorsRepo}
}

func (s *AuthorsService) Create(params *dto.AuthorsCreateReq) (*dto.AuthorsCreateResp, error) {
	newItem := params.ToEntity()

	err := s.repo.Create(&newItem)
	if err != nil {
		return nil, err
	}

	var resp dto.AuthorsCreateResp
	resp.FromEntity(&newItem)

	return &resp, nil
}

func (s *AuthorsService) GetByID(id uint) (dto.AuthorsDetailResp, error) {
	var resp dto.AuthorsDetailResp

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

func (s *AuthorsService) GetList() ([]dto.AuthorsDetailResp, error) {
	var resp []dto.AuthorsDetailResp

	items, err := s.repo.List()
	if err != nil {
		return nil, err
	}
	if len(items) < 1 {
		return nil, exception.ErrUserNotFound
	}

	for _, item := range items {
		var t dto.AuthorsDetailResp
		t.FromEntity(&item)

		resp = append(resp, t)
	}

	return resp, nil
}

func (s *AuthorsService) Update(params *dto.AuthorsUpdateReq) error {
	if params.ID <= 0 {
		return exception.ErrUserNotFound
	}

	birthDate, err := params.GetBirthDate()
	if err != nil {
		exception.LogError(err, "AuthorsService.Update")
		return exception.ErrDateParsing
	}
	params.BirthDate = birthDate

	return s.repo.Update(params)
}

func (s *AuthorsService) Delete(id uint) error {
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

