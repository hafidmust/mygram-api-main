package usecase_test

import (
	"context"
	"errors"
	"mygram-api/domain"
	"mygram-api/domain/mocks"
	"testing"
	"time"

	photoUseCase "mygram-api/photo/usecase"

	"github.com/asaskevich/govalidator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFetch(t *testing.T) {
	mockPhoto := domain.Photo{
		ID:       "photo-123",
		Title:    "A Title",
		Caption:  "A caption",
		PhotoUrl: "https://www.example.com/image.jpg",
		UserID:   "user-123",
	}

	mockPhotos := make([]domain.Photo, 0)

	mockPhotos = append(mockPhotos, mockPhoto)

	mockPhotoRepository := new(mocks.PhotoRepository)
	photoUseCase := photoUseCase.NewPhotoUseCase(mockPhotoRepository)

	t.Run("fetch all photos correctly", func(t *testing.T) {
		mockPhotoRepository.On("Fetch", mock.Anything, mock.AnythingOfType("*[]domain.Photo")).Return(nil).Once()

		err := photoUseCase.Fetch(context.Background(), &mockPhotos)

		assert.NoError(t, err)
	})
}

func TestStore(t *testing.T) {
	now := time.Now()
	mockAddedPhoto := domain.Photo{
		ID:        "photo-123",
		Title:     "A Title",
		Caption:   "A caption",
		PhotoUrl:  "https://www.example.com/image.jpg",
		UserID:    "user-123",
		CreatedAt: &now,
	}

	mockPhotoRepository := new(mocks.PhotoRepository)
	photoUseCase := photoUseCase.NewPhotoUseCase(mockPhotoRepository)

	t.Run("add photo correctly", func(t *testing.T) {
		tempMockAddPhoto := domain.Photo{
			Title:    "A Title",
			Caption:  "A caption",
			PhotoUrl: "https://www.example.com/image.jpg",
		}

		tempMockAddPhoto.ID = "photo-123"

		mockPhotoRepository.On("Store", mock.Anything, mock.AnythingOfType("*domain.Photo")).Return(nil).Once()

		err := photoUseCase.Store(context.Background(), &tempMockAddPhoto)

		assert.NoError(t, err)

		_, err = govalidator.ValidateStruct(tempMockAddPhoto)

		assert.NoError(t, err)
		assert.Equal(t, mockAddedPhoto.ID, tempMockAddPhoto.ID)
		assert.Equal(t, mockAddedPhoto.Title, tempMockAddPhoto.Title)
		assert.Equal(t, mockAddedPhoto.Caption, tempMockAddPhoto.Caption)
		assert.Equal(t, mockAddedPhoto.PhotoUrl, tempMockAddPhoto.PhotoUrl)
		mockPhotoRepository.AssertExpectations(t)
	})

	t.Run("add photo with empty title", func(t *testing.T) {
		tempMockAddPhoto := domain.Photo{
			Title:    "",
			Caption:  "A caption",
			PhotoUrl: "https://www.example.com/image.jpg",
		}

		tempMockAddPhoto.ID = "photo-123"

		mockPhotoRepository.On("Store", mock.Anything, mock.AnythingOfType("*domain.Photo")).Return(nil).Once()

		err := photoUseCase.Store(context.Background(), &tempMockAddPhoto)

		assert.NoError(t, err)

		_, err = govalidator.ValidateStruct(tempMockAddPhoto)

		assert.Error(t, err)
		assert.Equal(t, mockAddedPhoto.ID, tempMockAddPhoto.ID)
		assert.NotEqual(t, mockAddedPhoto.Title, tempMockAddPhoto.Title)
		assert.Equal(t, mockAddedPhoto.Caption, tempMockAddPhoto.Caption)
		assert.Equal(t, mockAddedPhoto.PhotoUrl, tempMockAddPhoto.PhotoUrl)
		mockPhotoRepository.AssertExpectations(t)
	})

	t.Run("add photo with empty photo url", func(t *testing.T) {
		tempMockAddPhoto := domain.Photo{
			Title:    "A Title",
			Caption:  "A caption",
			PhotoUrl: "",
		}

		tempMockAddPhoto.ID = "photo-123"

		mockPhotoRepository.On("Store", mock.Anything, mock.AnythingOfType("*domain.Photo")).Return(nil).Once()

		err := photoUseCase.Store(context.Background(), &tempMockAddPhoto)

		assert.NoError(t, err)

		_, err = govalidator.ValidateStruct(tempMockAddPhoto)

		assert.Error(t, err)
		assert.Equal(t, mockAddedPhoto.ID, tempMockAddPhoto.ID)
		assert.Equal(t, mockAddedPhoto.Title, tempMockAddPhoto.Title)
		assert.Equal(t, mockAddedPhoto.Caption, tempMockAddPhoto.Caption)
		assert.NotEqual(t, mockAddedPhoto.PhotoUrl, tempMockAddPhoto.PhotoUrl)
		mockPhotoRepository.AssertExpectations(t)
	})

	t.Run("add photo with not contain needed property", func(t *testing.T) {
		tempMockAddPhoto := domain.Photo{
			Title:   "A Title",
			Caption: "A caption",
		}

		tempMockAddPhoto.ID = "photo-123"

		mockPhotoRepository.On("Store", mock.Anything, mock.AnythingOfType("*domain.Photo")).Return(nil).Once()

		err := photoUseCase.Store(context.Background(), &tempMockAddPhoto)

		assert.NoError(t, err)

		_, err = govalidator.ValidateStruct(tempMockAddPhoto)

		assert.Error(t, err)
		assert.Equal(t, mockAddedPhoto.ID, tempMockAddPhoto.ID)
		assert.Equal(t, mockAddedPhoto.Title, tempMockAddPhoto.Title)
		assert.Equal(t, mockAddedPhoto.Caption, tempMockAddPhoto.Caption)
		assert.NotEqual(t, mockAddedPhoto.PhotoUrl, tempMockAddPhoto.PhotoUrl)
		mockPhotoRepository.AssertExpectations(t)
	})
}

func TestGetBy(t *testing.T) {
	var mockPhoto *domain.Photo

	now := time.Now()

	mockPhoto = &domain.Photo{
		ID:        "photo-123",
		Title:     "A Title",
		Caption:   "A caption",
		PhotoUrl:  "https://www.example.com/image.jpg",
		UserID:    "user-123",
		CreatedAt: &now,
	}

	mockPhotoRepository := new(mocks.PhotoRepository)
	photoUseCase := photoUseCase.NewPhotoUseCase(mockPhotoRepository)

	t.Run("get by id correctly", func(t *testing.T) {
		mockPhotoID := "photo-123"

		mockPhotoRepository.On("GetByID", mock.Anything, mock.AnythingOfType("*domain.Photo"), mock.AnythingOfType("string")).Return(nil).Once()

		err := photoUseCase.GetByID(context.Background(), mockPhoto, mockPhotoID)

		assert.NoError(t, err)
		assert.Equal(t, mockPhoto.ID, mockPhotoID)
		mockPhotoRepository.AssertExpectations(t)
	})

	t.Run("get by id with not found photo", func(t *testing.T) {
		mockPhotoID := "photo-234"

		mockPhotoRepository.On("GetByID", mock.Anything, mock.AnythingOfType("*domain.Photo"), mock.AnythingOfType("string")).Return(nil).Once()

		err := photoUseCase.GetByID(context.Background(), mockPhoto, mockPhotoID)

		assert.NoError(t, err)
		assert.NotEqual(t, mockPhoto.ID, mockPhotoID)
		mockPhotoRepository.AssertExpectations(t)
	})
}

func TestUpdate(t *testing.T) {
	now := time.Now()
	mockUpdatedPhoto := domain.Photo{
		ID:        "photo-123",
		Title:     "A New Title",
		Caption:   "A new caption",
		PhotoUrl:  "https://www.example.com/new-image.jpg",
		UserID:    "user-123",
		UpdatedAt: &now,
	}

	mockPhotoRepository := new(mocks.PhotoRepository)
	photoUseCase := photoUseCase.NewPhotoUseCase(mockPhotoRepository)

	t.Run("update photo correctly", func(t *testing.T) {
		tempMockPhotoID := "photo-123"
		tempMockUpdatePhoto := domain.Photo{
			Title:    "A New Title",
			Caption:  "A new caption",
			PhotoUrl: "https://www.example.com/new-image.jpg",
		}

		mockPhotoRepository.On("Update", mock.Anything, mock.AnythingOfType("domain.Photo"), mock.AnythingOfType("string")).Return(mockUpdatedPhoto, nil).Once()

		photo, err := photoUseCase.Update(context.Background(), tempMockUpdatePhoto, tempMockPhotoID)

		assert.NoError(t, err)

		tempMockUpdatedPhoto := domain.Photo{
			ID:        tempMockPhotoID,
			Title:     tempMockUpdatePhoto.Title,
			Caption:   tempMockUpdatePhoto.Caption,
			PhotoUrl:  tempMockUpdatePhoto.PhotoUrl,
			UserID:    "user-123",
			UpdatedAt: &now,
		}

		_, err = govalidator.ValidateStruct(tempMockUpdatedPhoto)

		assert.NoError(t, err)
		assert.Equal(t, photo, tempMockUpdatedPhoto)
		assert.Equal(t, mockUpdatedPhoto.Title, tempMockUpdatePhoto.Title)
		assert.Equal(t, mockUpdatedPhoto.Caption, tempMockUpdatePhoto.Caption)
		assert.Equal(t, mockUpdatedPhoto.PhotoUrl, tempMockUpdatedPhoto.PhotoUrl)
		mockPhotoRepository.AssertExpectations(t)
	})

	t.Run("update photo with empty title", func(t *testing.T) {
		tempMockPhotoID := "photo-123"
		tempMockUpdatePhoto := domain.Photo{
			Title:    "",
			Caption:  "A new caption",
			PhotoUrl: "https://www.example.com/new-image.jpg",
		}

		mockPhotoRepository.On("Update", mock.Anything, mock.AnythingOfType("domain.Photo"), mock.AnythingOfType("string")).Return(mockUpdatedPhoto, nil).Once()

		photo, err := photoUseCase.Update(context.Background(), tempMockUpdatePhoto, tempMockPhotoID)

		assert.NoError(t, err)

		tempMockUpdatedPhoto := domain.Photo{
			ID:        tempMockPhotoID,
			Title:     tempMockUpdatePhoto.Title,
			Caption:   tempMockUpdatePhoto.Caption,
			PhotoUrl:  tempMockUpdatePhoto.PhotoUrl,
			UserID:    "user-123",
			UpdatedAt: &now,
		}

		_, err = govalidator.ValidateStruct(tempMockUpdatedPhoto)

		assert.Error(t, err)
		assert.NotEqual(t, photo, tempMockUpdatedPhoto)
		assert.NotEqual(t, mockUpdatedPhoto.Title, tempMockUpdatePhoto.Title)
		assert.Equal(t, mockUpdatedPhoto.Caption, tempMockUpdatePhoto.Caption)
		assert.Equal(t, mockUpdatedPhoto.PhotoUrl, tempMockUpdatedPhoto.PhotoUrl)
		mockPhotoRepository.AssertExpectations(t)
	})

	t.Run("update photo with empty photo url", func(t *testing.T) {
		tempMockPhotoID := "photo-123"
		tempMockUpdatePhoto := domain.Photo{
			Title:    "A New Title",
			Caption:  "A new caption",
			PhotoUrl: "",
		}

		mockPhotoRepository.On("Update", mock.Anything, mock.AnythingOfType("domain.Photo"), mock.AnythingOfType("string")).Return(mockUpdatedPhoto, nil).Once()

		photo, err := photoUseCase.Update(context.Background(), tempMockUpdatePhoto, tempMockPhotoID)

		assert.NoError(t, err)

		tempMockUpdatedPhoto := domain.Photo{
			ID:        tempMockPhotoID,
			Title:     tempMockUpdatePhoto.Title,
			Caption:   tempMockUpdatePhoto.Caption,
			PhotoUrl:  tempMockUpdatePhoto.PhotoUrl,
			UserID:    "user-123",
			UpdatedAt: &now,
		}

		_, err = govalidator.ValidateStruct(tempMockUpdatedPhoto)

		assert.Error(t, err)
		assert.NotEqual(t, photo, tempMockUpdatedPhoto)
		assert.Equal(t, mockUpdatedPhoto.Title, tempMockUpdatePhoto.Title)
		assert.Equal(t, mockUpdatedPhoto.Caption, tempMockUpdatePhoto.Caption)
		assert.NotEqual(t, mockUpdatedPhoto.PhotoUrl, tempMockUpdatedPhoto.PhotoUrl)
		mockPhotoRepository.AssertExpectations(t)
	})

	t.Run("update photo with empty title and photo url", func(t *testing.T) {
		tempMockPhotoID := "photo-123"
		tempMockUpdatePhoto := domain.Photo{
			Title:    "",
			Caption:  "A new caption",
			PhotoUrl: "",
		}

		mockPhotoRepository.On("Update", mock.Anything, mock.AnythingOfType("domain.Photo"), mock.AnythingOfType("string")).Return(mockUpdatedPhoto, nil).Once()

		photo, err := photoUseCase.Update(context.Background(), tempMockUpdatePhoto, tempMockPhotoID)

		assert.NoError(t, err)

		tempMockUpdatedPhoto := domain.Photo{
			ID:        tempMockPhotoID,
			Title:     tempMockUpdatePhoto.Title,
			Caption:   tempMockUpdatePhoto.Caption,
			PhotoUrl:  tempMockUpdatePhoto.PhotoUrl,
			UserID:    "user-123",
			UpdatedAt: &now,
		}

		_, err = govalidator.ValidateStruct(tempMockUpdatedPhoto)

		assert.Error(t, err)
		assert.NotEqual(t, photo, tempMockUpdatedPhoto)
		assert.NotEqual(t, mockUpdatedPhoto.Title, tempMockUpdatePhoto.Title)
		assert.Equal(t, mockUpdatedPhoto.Caption, tempMockUpdatePhoto.Caption)
		assert.NotEqual(t, mockUpdatedPhoto.PhotoUrl, tempMockUpdatedPhoto.PhotoUrl)
		mockPhotoRepository.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	mockPhoto := domain.Photo{
		ID:       "photo-123",
		Title:    "A Title",
		Caption:  "A caption",
		PhotoUrl: "https://www.example.com/image.jpg",
		UserID:   "user-123",
	}

	mockPhotoRepository := new(mocks.PhotoRepository)
	photoUseCase := photoUseCase.NewPhotoUseCase(mockPhotoRepository)

	t.Run("delete photo correctly", func(t *testing.T) {
		mockPhotoRepository.On("Delete", mock.Anything, mock.AnythingOfType("string")).Return(nil).Once()

		err := photoUseCase.Delete(context.Background(), mockPhoto.ID)

		assert.NoError(t, err)
		mockPhotoRepository.AssertExpectations(t)
	})

	t.Run("delete photo with not found photo", func(t *testing.T) {
		mockPhotoRepository.On("Delete", mock.Anything, mock.AnythingOfType("string")).Return(errors.New("fail")).Once()

		err := photoUseCase.Delete(context.Background(), "photo-234")

		assert.Error(t, err)
		mockPhotoRepository.AssertExpectations(t)
	})
}
