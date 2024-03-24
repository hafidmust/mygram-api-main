package usecase_test

import (
	"context"
	"errors"
	"mygram-api/domain"
	"mygram-api/domain/mocks"
	"testing"
	"time"

	commentUseCase "mygram-api/comment/usecase"

	"github.com/asaskevich/govalidator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestFetch(t *testing.T) {
	now := time.Now()
	mockComment := domain.Comment{
		ID:        "comment-123",
		UserID:    "user-123",
		PhotoID:   "photo-123",
		Message:   "A message",
		CreatedAt: &now,
		UpdatedAt: &now,
	}

	mockComments := make([]domain.Comment, 0)

	mockComments = append(mockComments, mockComment)

	mockCommentRepository := new(mocks.CommentRepository)
	commentUseCase := commentUseCase.NewCommentUseCase(mockCommentRepository)

	t.Run("fetch all comments correctly", func(t *testing.T) {
		mockCommentRepository.On("Fetch", mock.Anything, mock.AnythingOfType("*[]domain.Comment"), mock.AnythingOfType("string")).Return(nil).Once()

		err := commentUseCase.Fetch(context.Background(), &mockComments, mockComment.UserID)

		assert.NoError(t, err)
	})
}

func TestStore(t *testing.T) {
	now := time.Now()
	mockAddedComment := domain.Comment{
		ID:        "comment-123",
		UserID:    "user-123",
		PhotoID:   "photo-123",
		Message:   "A comment",
		CreatedAt: &now,
	}

	mockCommentRepository := new(mocks.CommentRepository)
	commentUseCase := commentUseCase.NewCommentUseCase(mockCommentRepository)

	t.Run("add comment correctly", func(t *testing.T) {
		tempMockAddComment := domain.Comment{
			Message: "A comment",
			PhotoID: "photo-123",
		}

		tempMockAddComment.ID = "comment-123"

		mockCommentRepository.On("Store", mock.Anything, mock.AnythingOfType("*domain.Comment")).Return(nil).Once()

		err := commentUseCase.Store(context.Background(), &tempMockAddComment)

		assert.NoError(t, err)

		_, err = govalidator.ValidateStruct(tempMockAddComment)

		assert.NoError(t, err)
		assert.Equal(t, mockAddedComment.ID, tempMockAddComment.ID)
		assert.Equal(t, mockAddedComment.Message, tempMockAddComment.Message)
		assert.Equal(t, mockAddedComment.PhotoID, tempMockAddComment.PhotoID)
		mockCommentRepository.AssertExpectations(t)
	})

	t.Run("add comment with empty message", func(t *testing.T) {
		tempMockAddComment := domain.Comment{
			Message: "",
			PhotoID: "photo-123",
		}

		tempMockAddComment.ID = "comment-123"

		mockCommentRepository.On("Store", mock.Anything, mock.AnythingOfType("*domain.Comment")).Return(nil).Once()

		err := commentUseCase.Store(context.Background(), &tempMockAddComment)

		assert.NoError(t, err)

		_, err = govalidator.ValidateStruct(tempMockAddComment)

		assert.Error(t, err)
		assert.Equal(t, mockAddedComment.ID, tempMockAddComment.ID)
		assert.NotEqual(t, mockAddedComment.Message, tempMockAddComment.Message)
		assert.Equal(t, mockAddedComment.PhotoID, tempMockAddComment.PhotoID)
		mockCommentRepository.AssertExpectations(t)
	})

	t.Run("add comment with empty photo id", func(t *testing.T) {
		tempMockAddComment := domain.Comment{
			Message: "A comment",
			PhotoID: "",
		}

		tempMockAddComment.ID = "comment-123"

		mockCommentRepository.On("Store", mock.Anything, mock.AnythingOfType("*domain.Comment")).Return(errors.New("fail")).Once()

		err := commentUseCase.Store(context.Background(), &tempMockAddComment)

		assert.Error(t, err)

		_, err = govalidator.ValidateStruct(tempMockAddComment)

		assert.NoError(t, err)
		assert.Equal(t, mockAddedComment.ID, tempMockAddComment.ID)
		assert.Equal(t, mockAddedComment.Message, tempMockAddComment.Message)
		assert.NotEqual(t, mockAddedComment.PhotoID, tempMockAddComment.PhotoID)
		mockCommentRepository.AssertExpectations(t)
	})

	t.Run("add comment with not contain needed property", func(t *testing.T) {
		tempMockAddComment := domain.Comment{
			PhotoID: "photo-123",
		}

		tempMockAddComment.ID = "comment-123"

		mockCommentRepository.On("Store", mock.Anything, mock.AnythingOfType("*domain.Comment")).Return(nil).Once()

		err := commentUseCase.Store(context.Background(), &tempMockAddComment)

		assert.NoError(t, err)

		_, err = govalidator.ValidateStruct(tempMockAddComment)

		assert.Error(t, err)
		assert.Equal(t, mockAddedComment.ID, tempMockAddComment.ID)
		assert.NotEqual(t, mockAddedComment.Message, tempMockAddComment.Message)
		assert.Equal(t, mockAddedComment.PhotoID, tempMockAddComment.PhotoID)
		mockCommentRepository.AssertExpectations(t)
	})
}

func TestGetBy(t *testing.T) {
	var mockComment *domain.Comment

	now := time.Now()

	mockComment = &domain.Comment{
		ID:        "comment-123",
		UserID:    "user-123",
		PhotoID:   "photo-123",
		Message:   "A message",
		CreatedAt: &now,
		UpdatedAt: &now,
	}

	mockCommentRepository := new(mocks.CommentRepository)
	commentUseCase := commentUseCase.NewCommentUseCase(mockCommentRepository)

	t.Run("get by id correctly", func(t *testing.T) {
		mockCommentID := "comment-123"

		mockCommentRepository.On("GetByID", mock.Anything, mock.AnythingOfType("*domain.Comment"), mock.AnythingOfType("string")).Return(nil).Once()

		err := commentUseCase.GetByID(context.Background(), mockComment, mockCommentID)

		assert.NoError(t, err)
		assert.Equal(t, mockComment.ID, mockCommentID)
		mockCommentRepository.AssertExpectations(t)
	})

	t.Run("get by id with not found comment", func(t *testing.T) {
		mockCommentID := "comment-234"

		mockCommentRepository.On("GetByID", mock.Anything, mock.AnythingOfType("*domain.Comment"), mock.AnythingOfType("string")).Return(nil).Once()

		err := commentUseCase.GetByID(context.Background(), mockComment, mockCommentID)

		assert.NoError(t, err)
		assert.NotEqual(t, mockComment.ID, mockCommentID)
		mockCommentRepository.AssertExpectations(t)
	})
}

func TestUpdate(t *testing.T) {
	now := time.Now()
	mockUpdatedComment := domain.Photo{
		ID:        "photo-123",
		Title:     "A Title",
		Caption:   "A caption",
		PhotoUrl:  "https://www.example.com/image.jpg",
		UserID:    "user-123",
		UpdatedAt: &now,
	}

	mockCommentRepository := new(mocks.CommentRepository)
	commentUseCase := commentUseCase.NewCommentUseCase(mockCommentRepository)

	t.Run("update comment correctly", func(t *testing.T) {
		tempMockCommentID := "comment-123"
		tempMockUpdateComment := domain.Comment{
			Message: "A new comment",
		}

		mockCommentRepository.On("Update", mock.Anything, mock.AnythingOfType("domain.Comment"), mock.AnythingOfType("string")).Return(mockUpdatedComment, nil).Once()

		comment, err := commentUseCase.Update(context.Background(), tempMockUpdateComment, tempMockCommentID)

		assert.NoError(t, err)

		tempMockUpdatedComment := domain.Photo{
			ID:        "photo-123",
			Title:     "A Title",
			Caption:   "A caption",
			PhotoUrl:  "https://www.example.com/image.jpg",
			UserID:    "user-123",
			UpdatedAt: &now,
		}

		_, err = govalidator.ValidateStruct(tempMockUpdateComment)

		assert.NoError(t, err)
		assert.Equal(t, comment, tempMockUpdatedComment)
		assert.Equal(t, mockUpdatedComment.ID, tempMockUpdatedComment.ID)
		assert.Equal(t, mockUpdatedComment.Title, tempMockUpdatedComment.Title)
		assert.Equal(t, mockUpdatedComment.Caption, tempMockUpdatedComment.Caption)
		assert.Equal(t, mockUpdatedComment.PhotoUrl, tempMockUpdatedComment.PhotoUrl)
		assert.Equal(t, mockUpdatedComment.UserID, tempMockUpdatedComment.UserID)
		assert.Equal(t, mockUpdatedComment.UpdatedAt, tempMockUpdatedComment.UpdatedAt)
		mockCommentRepository.AssertExpectations(t)
	})

	t.Run("update comment with empty message", func(t *testing.T) {
		tempMockCommentID := "Comment-123"
		tempMockUpdateComment := domain.Comment{
			Message: "",
		}

		mockCommentRepository.On("Update", mock.Anything, mock.AnythingOfType("domain.Comment"), mock.AnythingOfType("string")).Return(mockUpdatedComment, nil).Once()

		comment, err := commentUseCase.Update(context.Background(), tempMockUpdateComment, tempMockCommentID)

		assert.NoError(t, err)

		tempMockUpdatedComment := domain.Photo{
			ID:        "photo-123",
			Title:     "A Title",
			Caption:   "A caption",
			PhotoUrl:  "https://www.example.com/image.jpg",
			UserID:    "user-123",
			UpdatedAt: &now,
		}

		_, err = govalidator.ValidateStruct(tempMockUpdateComment)

		assert.Error(t, err)
		assert.Equal(t, comment, tempMockUpdatedComment)
		assert.Equal(t, mockUpdatedComment.ID, tempMockUpdatedComment.ID)
		assert.Equal(t, mockUpdatedComment.Title, tempMockUpdatedComment.Title)
		assert.Equal(t, mockUpdatedComment.Caption, tempMockUpdatedComment.Caption)
		assert.Equal(t, mockUpdatedComment.PhotoUrl, tempMockUpdatedComment.PhotoUrl)
		assert.Equal(t, mockUpdatedComment.UserID, tempMockUpdatedComment.UserID)
		assert.Equal(t, mockUpdatedComment.UpdatedAt, tempMockUpdatedComment.UpdatedAt)
		mockCommentRepository.AssertExpectations(t)
	})

	t.Run("update comment with not contain property", func(t *testing.T) {
		tempMockCommentID := "comment-123"
		tempMockUpdateComment := domain.Comment{}

		mockCommentRepository.On("Update", mock.Anything, mock.AnythingOfType("domain.Comment"), mock.AnythingOfType("string")).Return(mockUpdatedComment, nil).Once()

		comment, err := commentUseCase.Update(context.Background(), tempMockUpdateComment, tempMockCommentID)

		assert.NoError(t, err)

		tempMockUpdatedComment := domain.Photo{
			ID:        "photo-123",
			Title:     "A Title",
			Caption:   "A caption",
			PhotoUrl:  "https://www.example.com/image.jpg",
			UserID:    "user-123",
			UpdatedAt: &now,
		}

		_, err = govalidator.ValidateStruct(tempMockUpdateComment)

		assert.Error(t, err)
		assert.Equal(t, comment, tempMockUpdatedComment)
		assert.Equal(t, mockUpdatedComment.ID, tempMockUpdatedComment.ID)
		assert.Equal(t, mockUpdatedComment.Title, tempMockUpdatedComment.Title)
		assert.Equal(t, mockUpdatedComment.Caption, tempMockUpdatedComment.Caption)
		assert.Equal(t, mockUpdatedComment.PhotoUrl, tempMockUpdatedComment.PhotoUrl)
		assert.Equal(t, mockUpdatedComment.UserID, tempMockUpdatedComment.UserID)
		assert.Equal(t, mockUpdatedComment.UpdatedAt, tempMockUpdatedComment.UpdatedAt)
		mockCommentRepository.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	now := time.Now()
	mockComment := domain.Comment{
		ID:        "comment-123",
		UserID:    "user-123",
		PhotoID:   "photo-123",
		Message:   "A message",
		CreatedAt: &now,
		UpdatedAt: &now,
	}

	mockCommentRepository := new(mocks.CommentRepository)
	commentUseCase := commentUseCase.NewCommentUseCase(mockCommentRepository)

	t.Run("delete comment correctly", func(t *testing.T) {
		mockCommentRepository.On("Delete", mock.Anything, mock.AnythingOfType("string")).Return(nil).Once()

		err := commentUseCase.Delete(context.Background(), mockComment.ID)

		assert.NoError(t, err)
		mockCommentRepository.AssertExpectations(t)
	})

	t.Run("delete comment with not found Comment", func(t *testing.T) {
		mockCommentRepository.On("Delete", mock.Anything, mock.AnythingOfType("string")).Return(errors.New("fail")).Once()

		err := commentUseCase.Delete(context.Background(), "comment-234")

		assert.Error(t, err)
		mockCommentRepository.AssertExpectations(t)
	})
}
