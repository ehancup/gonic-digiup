package repository

import (
	"base-gin/app/domain/dao"

	"base-gin/storage"

	"gorm.io/gorm"
)

type FormRepository struct {
	db *gorm.DB
}

func newFormRepo(db *gorm.DB) *FormRepository {
	return &FormRepository{db: db}
}
func (r *FormRepository) Create(newItem *dao.FormTable) error {
	ctx, cancelFunc := storage.NewDBContext()
	defer cancelFunc()
	tx := r.db.WithContext(ctx).Create(&newItem)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}