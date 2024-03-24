package usecase_test

import (
	"context"
	"errors"
	"mygram-api/domain"
	"mygram-api/domain/mocks"
	"testing"
	"time"

	socialMediaUseCase "mygram-api/socialmedia/usecase"

	"github.com/asaskevich/govalidator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFetch(t *testing.T) {
	now := time.Now()
	mockSocialMedia := domain.SocialMedia{
		ID:             "socialmedia-123",
		Name:           "Example",
		SocialMediaUrl: "https://www.example.com/johndoe",
		UserID:         "user-123",
		CreatedAt:      &now,
		UpdatedAt:      &now,
	}

	mockSocialMedias := make([]domain.SocialMedia, 0)

	mockSocialMedias = append(mockSocialMedias, mockSocialMedia)

	mockSocialMediaRepository := new(mocks.SocialMediaRepository)
	socialMediaUseCase := socialMediaUseCase.NewSocialMediaUseCase(mockSocialMediaRepository)

	t.Run("fetch all social media correctly", func(t *testing.T) {
		mockSocialMediaRepository.On("Fetch", mock.Anything, mock.AnythingOfType("*[]domain.SocialMedia"), mock.AnythingOfType("string")).Return(nil).Once()

		err := socialMediaUseCase.Fetch(context.Background(), &mockSocialMedias, mockSocialMedia.UserID)

		assert.NoError(t, err)
	})
}

func TestStore(t *testing.T) {
	now := time.Now()
	mockAddedSocialMedia := domain.SocialMedia{
		ID:             "socialmedia-123",
		Name:           "Example",
		SocialMediaUrl: "https://www.example.com/johndoe",
		UserID:         "user-123",
		CreatedAt:      &now,
	}

	mockSocialMediaRepository := new(mocks.SocialMediaRepository)
	socialMediaUseCase := socialMediaUseCase.NewSocialMediaUseCase(mockSocialMediaRepository)

	t.Run("add social media correctly", func(t *testing.T) {
		tempMockAddSocialMedia := domain.SocialMedia{
			Name:           "Example",
			SocialMediaUrl: "https://www.example.com/johndoe",
		}

		tempMockAddSocialMedia.ID = "socialmedia-123"

		mockSocialMediaRepository.On("Store", mock.Anything, mock.AnythingOfType("*domain.SocialMedia")).Return(nil).Once()

		err := socialMediaUseCase.Store(context.Background(), &tempMockAddSocialMedia)

		assert.NoError(t, err)

		_, err = govalidator.ValidateStruct(tempMockAddSocialMedia)

		assert.NoError(t, err)
		assert.Equal(t, mockAddedSocialMedia.ID, tempMockAddSocialMedia.ID)
		assert.Equal(t, mockAddedSocialMedia.Name, tempMockAddSocialMedia.Name)
		assert.Equal(t, mockAddedSocialMedia.SocialMediaUrl, tempMockAddSocialMedia.SocialMediaUrl)
		mockSocialMediaRepository.AssertExpectations(t)
	})

	t.Run("add social media with empty name", func(t *testing.T) {
		tempMockAddSocialMedia := domain.SocialMedia{
			Name:           "",
			SocialMediaUrl: "https://www.example.com/johndoe",
		}

		tempMockAddSocialMedia.ID = "socialmedia-123"

		mockSocialMediaRepository.On("Store", mock.Anything, mock.AnythingOfType("*domain.SocialMedia")).Return(nil).Once()

		err := socialMediaUseCase.Store(context.Background(), &tempMockAddSocialMedia)

		assert.NoError(t, err)

		_, err = govalidator.ValidateStruct(tempMockAddSocialMedia)

		assert.Error(t, err)
		assert.Equal(t, mockAddedSocialMedia.ID, tempMockAddSocialMedia.ID)
		assert.NotEqual(t, mockAddedSocialMedia.Name, tempMockAddSocialMedia.Name)
		assert.Equal(t, mockAddedSocialMedia.SocialMediaUrl, tempMockAddSocialMedia.SocialMediaUrl)
		mockSocialMediaRepository.AssertExpectations(t)
	})

	t.Run("add social media with empty social media url", func(t *testing.T) {
		tempMockAddSocialMedia := domain.SocialMedia{
			Name:           "Example",
			SocialMediaUrl: "",
		}

		tempMockAddSocialMedia.ID = "socialmedia-123"

		mockSocialMediaRepository.On("Store", mock.Anything, mock.AnythingOfType("*domain.SocialMedia")).Return(errors.New("fail")).Once()

		err := socialMediaUseCase.Store(context.Background(), &tempMockAddSocialMedia)

		assert.Error(t, err)

		_, err = govalidator.ValidateStruct(tempMockAddSocialMedia)

		assert.Error(t, err)
		assert.Equal(t, mockAddedSocialMedia.ID, tempMockAddSocialMedia.ID)
		assert.Equal(t, mockAddedSocialMedia.Name, tempMockAddSocialMedia.Name)
		assert.NotEqual(t, mockAddedSocialMedia.SocialMediaUrl, tempMockAddSocialMedia.SocialMediaUrl)
		mockSocialMediaRepository.AssertExpectations(t)
	})

	t.Run("add social media with not contain needed property", func(t *testing.T) {
		tempMockAddSocialMedia := domain.SocialMedia{
			Name: "Example",
		}

		tempMockAddSocialMedia.ID = "socialmedia-123"

		mockSocialMediaRepository.On("Store", mock.Anything, mock.AnythingOfType("*domain.SocialMedia")).Return(nil).Once()

		err := socialMediaUseCase.Store(context.Background(), &tempMockAddSocialMedia)

		assert.NoError(t, err)

		_, err = govalidator.ValidateStruct(tempMockAddSocialMedia)

		assert.Error(t, err)
		assert.Equal(t, mockAddedSocialMedia.ID, tempMockAddSocialMedia.ID)
		assert.Equal(t, mockAddedSocialMedia.Name, tempMockAddSocialMedia.Name)
		assert.NotEqual(t, mockAddedSocialMedia.SocialMediaUrl, tempMockAddSocialMedia.SocialMediaUrl)
		mockSocialMediaRepository.AssertExpectations(t)
	})
}

func TestGetBy(t *testing.T) {
	var mockSocialMedia *domain.SocialMedia

	now := time.Now()

	mockSocialMedia = &domain.SocialMedia{
		ID:             "socialmedia-123",
		Name:           "Example",
		SocialMediaUrl: "https://www.example.com/johndoe",
		UserID:         "user-123",
		CreatedAt:      &now,
		UpdatedAt:      &now,
	}

	mockSocialMediaRepository := new(mocks.SocialMediaRepository)
	socialMediaUseCase := socialMediaUseCase.NewSocialMediaUseCase(mockSocialMediaRepository)

	t.Run("get by id correctly", func(t *testing.T) {
		mockSocialMediaID := "socialmedia-123"

		mockSocialMediaRepository.On("GetByID", mock.Anything, mock.AnythingOfType("*domain.SocialMedia"), mock.AnythingOfType("string")).Return(nil).Once()

		err := socialMediaUseCase.GetByID(context.Background(), mockSocialMedia, mockSocialMediaID)

		assert.NoError(t, err)
		assert.Equal(t, mockSocialMedia.ID, mockSocialMediaID)
		mockSocialMediaRepository.AssertExpectations(t)
	})

	t.Run("get by id with not found social media", func(t *testing.T) {
		mockSocialMediaID := "socialmedia-234"

		mockSocialMediaRepository.On("GetByID", mock.Anything, mock.AnythingOfType("*domain.SocialMedia"), mock.AnythingOfType("string")).Return(nil).Once()

		err := socialMediaUseCase.GetByID(context.Background(), mockSocialMedia, mockSocialMediaID)

		assert.NoError(t, err)
		assert.NotEqual(t, mockSocialMedia.ID, mockSocialMediaID)
		mockSocialMediaRepository.AssertExpectations(t)
	})
}

func TestUpdate(t *testing.T) {
	now := time.Now()
	mockUpdatedSocialMedia := domain.SocialMedia{
		ID:             "socialmedia-123",
		Name:           "New Example",
		SocialMediaUrl: "https://www.newexample.com/johndoe",
		UserID:         "user-123",
		UpdatedAt:      &now,
	}

	mockSocialMediaRepository := new(mocks.SocialMediaRepository)
	socialMediaUseCase := socialMediaUseCase.NewSocialMediaUseCase(mockSocialMediaRepository)

	t.Run("update social media correctly", func(t *testing.T) {
		tempMockSocialMediaID := "socialmedia-123"
		tempMockUpdateSocialMedia := domain.SocialMedia{
			Name:           "New Example",
			SocialMediaUrl: "https://www.newexample.com/johndoe",
		}

		mockSocialMediaRepository.On("Update", mock.Anything, mock.AnythingOfType("domain.SocialMedia"), mock.AnythingOfType("string")).Return(mockUpdatedSocialMedia, nil).Once()

		socialmedia, err := socialMediaUseCase.Update(context.Background(), tempMockUpdateSocialMedia, tempMockSocialMediaID)

		assert.NoError(t, err)

		tempMockUpdatedSocialMedia := domain.SocialMedia{
			ID:             tempMockSocialMediaID,
			Name:           tempMockUpdateSocialMedia.Name,
			SocialMediaUrl: tempMockUpdateSocialMedia.SocialMediaUrl,
			UserID:         "user-123",
			UpdatedAt:      &now,
		}

		_, err = govalidator.ValidateStruct(tempMockUpdateSocialMedia)

		assert.NoError(t, err)
		assert.Equal(t, socialmedia, tempMockUpdatedSocialMedia)
		assert.Equal(t, mockUpdatedSocialMedia.ID, tempMockUpdatedSocialMedia.ID)
		assert.Equal(t, mockUpdatedSocialMedia.Name, tempMockUpdatedSocialMedia.Name)
		assert.Equal(t, mockUpdatedSocialMedia.SocialMediaUrl, tempMockUpdatedSocialMedia.SocialMediaUrl)
		assert.Equal(t, mockUpdatedSocialMedia.UserID, tempMockUpdatedSocialMedia.UserID)
		assert.Equal(t, mockUpdatedSocialMedia.UpdatedAt, tempMockUpdatedSocialMedia.UpdatedAt)
		mockSocialMediaRepository.AssertExpectations(t)
	})

	t.Run("update social media with empty name", func(t *testing.T) {
		tempMockSocialMediaID := "socialmedia-123"
		tempMockUpdateSocialMedia := domain.SocialMedia{
			Name:           "",
			SocialMediaUrl: "https://www.newexample.com/johndoe",
		}

		mockSocialMediaRepository.On("Update", mock.Anything, mock.AnythingOfType("domain.SocialMedia"), mock.AnythingOfType("string")).Return(mockUpdatedSocialMedia, nil).Once()

		socialmedia, err := socialMediaUseCase.Update(context.Background(), tempMockUpdateSocialMedia, tempMockSocialMediaID)

		assert.NoError(t, err)

		tempMockUpdatedSocialMedia := domain.SocialMedia{
			ID:             tempMockSocialMediaID,
			Name:           tempMockUpdateSocialMedia.Name,
			SocialMediaUrl: tempMockUpdateSocialMedia.SocialMediaUrl,
			UserID:         "user-123",
			UpdatedAt:      &now,
		}

		_, err = govalidator.ValidateStruct(tempMockUpdateSocialMedia)

		assert.Error(t, err)
		assert.NotEqual(t, socialmedia, tempMockUpdatedSocialMedia)
		assert.Equal(t, mockUpdatedSocialMedia.ID, tempMockUpdatedSocialMedia.ID)
		assert.NotEqual(t, mockUpdatedSocialMedia.Name, tempMockUpdatedSocialMedia.Name)
		assert.Equal(t, mockUpdatedSocialMedia.SocialMediaUrl, tempMockUpdatedSocialMedia.SocialMediaUrl)
		assert.Equal(t, mockUpdatedSocialMedia.UserID, tempMockUpdatedSocialMedia.UserID)
		assert.Equal(t, mockUpdatedSocialMedia.UpdatedAt, tempMockUpdatedSocialMedia.UpdatedAt)
		mockSocialMediaRepository.AssertExpectations(t)
	})

	t.Run("update social media with empty social media url", func(t *testing.T) {
		tempMockSocialMediaID := "socialmedia-123"
		tempMockUpdateSocialMedia := domain.SocialMedia{
			Name:           "New Example",
			SocialMediaUrl: "",
		}

		mockSocialMediaRepository.On("Update", mock.Anything, mock.AnythingOfType("domain.SocialMedia"), mock.AnythingOfType("string")).Return(mockUpdatedSocialMedia, nil).Once()

		socialmedia, err := socialMediaUseCase.Update(context.Background(), tempMockUpdateSocialMedia, tempMockSocialMediaID)

		assert.NoError(t, err)

		tempMockUpdatedSocialMedia := domain.SocialMedia{
			ID:             tempMockSocialMediaID,
			Name:           tempMockUpdateSocialMedia.Name,
			SocialMediaUrl: tempMockUpdateSocialMedia.Name,
			UserID:         "user-123",
			UpdatedAt:      &now,
		}

		_, err = govalidator.ValidateStruct(tempMockUpdateSocialMedia)

		assert.Error(t, err)
		assert.NotEqual(t, socialmedia, tempMockUpdatedSocialMedia)
		assert.Equal(t, mockUpdatedSocialMedia.ID, tempMockUpdatedSocialMedia.ID)
		assert.Equal(t, mockUpdatedSocialMedia.Name, tempMockUpdatedSocialMedia.Name)
		assert.NotEqual(t, mockUpdatedSocialMedia.SocialMediaUrl, tempMockUpdatedSocialMedia.SocialMediaUrl)
		assert.Equal(t, mockUpdatedSocialMedia.UserID, tempMockUpdatedSocialMedia.UserID)
		assert.Equal(t, mockUpdatedSocialMedia.UpdatedAt, tempMockUpdatedSocialMedia.UpdatedAt)
		mockSocialMediaRepository.AssertExpectations(t)
	})

	t.Run("update social media with not contain property", func(t *testing.T) {
		tempMockSocialMediaID := "socialmedia-123"
		tempMockUpdateSocialMedia := domain.SocialMedia{
			Name: "New Example",
		}

		mockSocialMediaRepository.On("Update", mock.Anything, mock.AnythingOfType("domain.SocialMedia"), mock.AnythingOfType("string")).Return(mockUpdatedSocialMedia, nil).Once()

		socialmedia, err := socialMediaUseCase.Update(context.Background(), tempMockUpdateSocialMedia, tempMockSocialMediaID)

		assert.NoError(t, err)

		tempMockUpdatedSocialMedia := domain.SocialMedia{
			ID:             tempMockSocialMediaID,
			Name:           tempMockUpdateSocialMedia.Name,
			SocialMediaUrl: tempMockUpdateSocialMedia.Name,
			UserID:         "user-123",
			UpdatedAt:      &now,
		}

		_, err = govalidator.ValidateStruct(tempMockUpdateSocialMedia)

		assert.Error(t, err)
		assert.NotEqual(t, socialmedia, tempMockUpdatedSocialMedia)
		assert.Equal(t, mockUpdatedSocialMedia.ID, tempMockUpdatedSocialMedia.ID)
		assert.Equal(t, mockUpdatedSocialMedia.Name, tempMockUpdatedSocialMedia.Name)
		assert.NotEqual(t, mockUpdatedSocialMedia.SocialMediaUrl, tempMockUpdatedSocialMedia.SocialMediaUrl)
		assert.Equal(t, mockUpdatedSocialMedia.UserID, tempMockUpdatedSocialMedia.UserID)
		assert.Equal(t, mockUpdatedSocialMedia.UpdatedAt, tempMockUpdatedSocialMedia.UpdatedAt)
		mockSocialMediaRepository.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	now := time.Now()
	mockSocialMedia := domain.SocialMedia{
		ID:             "socialmedia-123",
		Name:           "Example",
		SocialMediaUrl: "https://www.example.com/johndoe",
		UserID:         "user-123",
		CreatedAt:      &now,
		UpdatedAt:      &now,
	}

	mockSocialMediaRepository := new(mocks.SocialMediaRepository)
	socialMediaUseCase := socialMediaUseCase.NewSocialMediaUseCase(mockSocialMediaRepository)

	t.Run("delete social media correctly", func(t *testing.T) {
		mockSocialMediaRepository.On("Delete", mock.Anything, mock.AnythingOfType("string")).Return(nil).Once()

		err := socialMediaUseCase.Delete(context.Background(), mockSocialMedia.ID)

		assert.NoError(t, err)
		mockSocialMediaRepository.AssertExpectations(t)
	})

	t.Run("delete social media with not found social media", func(t *testing.T) {
		mockSocialMediaRepository.On("Delete", mock.Anything, mock.AnythingOfType("string")).Return(errors.New("fail")).Once()

		err := socialMediaUseCase.Delete(context.Background(), "socialmedia-234")

		assert.Error(t, err)
		mockSocialMediaRepository.AssertExpectations(t)
	})
}
