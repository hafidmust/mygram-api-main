package repository

import (
	"context"
	"fmt"
	"mygram-api/domain"
	"time"

	gonanoid "github.com/matoous/go-nanoid/v2"
	"gorm.io/gorm"
)

type photoRepository struct {
	db *gorm.DB
}

func NewPhotoRepository(db *gorm.DB) *photoRepository {
	return &photoRepository{db}
}

func (photoRepository *photoRepository) Fetch(ctx context.Context, photos *[]domain.Photo) (err error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)

	defer cancel()

	if err = photoRepository.db.WithContext(ctx).Preload("User", func(db *gorm.DB) *gorm.DB {
		return db.Select("id", "username", "email")
	}).Find(&photos).Error; err != nil {
		return err
	}

	return
}

func (photoRepository *photoRepository) Store(ctx context.Context, photo *domain.Photo) (err error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)

	defer cancel()

	ID, _ := gonanoid.New(16)

	photo.ID = fmt.Sprintf("photo-%s", ID)

	if err := photoRepository.db.WithContext(ctx).Create(&photo).Error; err != nil {
		return err
	}

	return
}

func (photoRepository *photoRepository) GetByID(ctx context.Context, photo *domain.Photo, id string) (err error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)

	defer cancel()

	if err = photoRepository.db.WithContext(ctx).First(&photo, &id).Error; err != nil {
		return err
	}

	return
}

func (photoRepository *photoRepository) Update(ctx context.Context, photo domain.Photo, id string) (p domain.Photo, err error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)

	defer cancel()

	p = domain.Photo{}

	if err = photoRepository.db.WithContext(ctx).First(&p, &id).Error; err != nil {
		return p, err
	}

	if err = photoRepository.db.WithContext(ctx).Model(&p).Updates(photo).Error; err != nil {
		return p, err
	}

	return p, nil
}

func (photoRepository *photoRepository) Delete(ctx context.Context, id string) (err error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)

	defer cancel()

	if err = photoRepository.db.WithContext(ctx).First(&domain.Photo{}, &id).Error; err != nil {
		return err
	}

	if err = photoRepository.db.WithContext(ctx).Delete(&domain.Photo{}, &id).Error; err != nil {
		return err
	}

	return
}
