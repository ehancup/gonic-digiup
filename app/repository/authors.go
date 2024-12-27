package repository

import (
	"base-gin/app/domain/dao"
	"base-gin/app/domain/dto"
	"base-gin/exception"
	"base-gin/storage"
	"errors"

	"gorm.io/gorm"
)

type AuthorsRepository struct {
	db *gorm.DB
}

func newAuthorsRepo(db *gorm.DB) *AuthorsRepository {
	return &AuthorsRepository{db: db}
}

func (r *AuthorsRepository) Create(newItem *dao.Authors) error {
	ctx, cancelFunc := storage.NewDBContext()
	defer cancelFunc()
	tx := r.db.WithContext(ctx).Create(&newItem)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}

func (r *AuthorsRepository) GetByID(id uint) (*dao.Authors, error) {
	ctx, cancelFunc := storage.NewDBContext()
	defer cancelFunc()

	var item dao.Authors
	tx := r.db.WithContext(ctx).First(&item, id)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, exception.ErrUserNotFound
		}

		return nil, tx.Error
	}

	return &item, nil
}
func (r *AuthorsRepository) List() ([]dao.Authors, error) {
	ctx, cancelFunc := storage.NewDBContext()
	defer cancelFunc()

	var items []dao.Authors
	tx := r.db.WithContext(ctx).Find(&items)
	if tx.Error != nil {
		if errors.Is(tx.Error, gorm.ErrRecordNotFound) {
			return nil, exception.ErrUserNotFound
		}

		return nil, tx.Error
	}

	return items, nil
}

func (r *AuthorsRepository) Update(params *dto.AuthorsUpdateReq) error {
	ctx, cancelFunc := storage.NewDBContext()
	defer cancelFunc()

	tx := r.db.WithContext(ctx).Model(&dao.Authors{}).
		Where("id = ?", params.ID).
		Updates(map[string]interface{}{
			"fullname":   params.Fullname,
			"gender":     params.GetGender(),
			"birth_date": params.BirthDate,
		})

	return tx.Error
}

func (r *AuthorsRepository) Delete(id uint) error {
	ctx, cancelFunc := storage.NewDBContext()
	defer cancelFunc()

	tx := r.db.WithContext(ctx).Delete(&dao.Authors{}, id)

	return tx.Error
}