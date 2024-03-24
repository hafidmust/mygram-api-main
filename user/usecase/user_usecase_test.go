package usecase_test

import (
	"context"
	"errors"
	"mygram-api/domain"
	"mygram-api/domain/mocks"
	"mygram-api/helpers"
	"testing"
	"time"

	userUseCase "mygram-api/user/usecase"

	"github.com/asaskevich/govalidator"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestRegister(t *testing.T) {
	mockRegisteredUser := domain.User{
		ID:       "user-123",
		Age:      8,
		Email:    "johndoe@example.com",
		Password: "secret",
		Username: "johndoe",
	}

	mockUserRepository := new(mocks.UserRepository)
	userUseCase := userUseCase.NewUserUseCase(mockUserRepository)

	t.Run("register user correctly", func(t *testing.T) {
		tempMockRegisterUser := domain.User{
			Age:      8,
			Email:    "johndoe@example.com",
			Password: "secret",
			Username: "johndoe",
		}

		tempMockRegisterUser.ID = "user-123"

		mockUserRepository.On("Register", mock.Anything, mock.AnythingOfType("*domain.User")).Return(nil).Once()

		err := userUseCase.Register(context.Background(), &tempMockRegisterUser)

		assert.NoError(t, err)

		_, err = govalidator.ValidateStruct(tempMockRegisterUser)

		assert.NoError(t, err)
		assert.Equal(t, mockRegisteredUser.ID, tempMockRegisterUser.ID)
		assert.Equal(t, mockRegisteredUser.Age, tempMockRegisterUser.Age)
		assert.Equal(t, mockRegisteredUser.Email, tempMockRegisterUser.Email)
		assert.Equal(t, mockRegisteredUser.Password, tempMockRegisterUser.Password)
		assert.Equal(t, mockRegisteredUser.Username, tempMockRegisterUser.Username)
		mockUserRepository.AssertExpectations(t)
	})

	t.Run("register user with empty email", func(t *testing.T) {
		tempMockRegisterUser := domain.User{
			Age:      8,
			Email:    "",
			Password: "secret",
			Username: "johndoe",
		}

		tempMockRegisterUser.ID = "user-123"

		mockUserRepository.On("Register", mock.Anything, mock.AnythingOfType("*domain.User")).Return(nil).Once()

		err := userUseCase.Register(context.Background(), &tempMockRegisterUser)

		assert.NoError(t, err)

		_, err = govalidator.ValidateStruct(tempMockRegisterUser)

		assert.Error(t, err)
		assert.Equal(t, mockRegisteredUser.ID, tempMockRegisterUser.ID)
		assert.Equal(t, mockRegisteredUser.Age, tempMockRegisterUser.Age)
		assert.NotEqual(t, mockRegisteredUser.Email, tempMockRegisterUser.Email)
		assert.Equal(t, mockRegisteredUser.Password, tempMockRegisterUser.Password)
		assert.Equal(t, mockRegisteredUser.Username, tempMockRegisterUser.Username)
		mockUserRepository.AssertExpectations(t)
	})

	t.Run("register user with empty password", func(t *testing.T) {
		tempMockRegisterUser := domain.User{
			Age:      8,
			Email:    "johndoe@example.com",
			Password: "",
			Username: "johndoe",
		}

		tempMockRegisterUser.ID = "user-123"

		mockUserRepository.On("Register", mock.Anything, mock.AnythingOfType("*domain.User")).Return(nil).Once()

		err := userUseCase.Register(context.Background(), &tempMockRegisterUser)

		assert.NoError(t, err)

		_, err = govalidator.ValidateStruct(tempMockRegisterUser)

		assert.Error(t, err)
		assert.Equal(t, mockRegisteredUser.ID, tempMockRegisterUser.ID)
		assert.Equal(t, mockRegisteredUser.Age, tempMockRegisterUser.Age)
		assert.Equal(t, mockRegisteredUser.Email, tempMockRegisterUser.Email)
		assert.NotEqual(t, mockRegisteredUser.Password, tempMockRegisterUser.Password)
		assert.Equal(t, mockRegisteredUser.Username, tempMockRegisterUser.Username)
		mockUserRepository.AssertExpectations(t)
	})

	t.Run("register user with empty username", func(t *testing.T) {
		tempMockRegisterUser := domain.User{
			Age:      8,
			Email:    "johndoe@example.com",
			Password: "secret",
			Username: "",
		}

		tempMockRegisterUser.ID = "user-123"

		mockUserRepository.On("Register", mock.Anything, mock.AnythingOfType("*domain.User")).Return(nil).Once()

		err := userUseCase.Register(context.Background(), &tempMockRegisterUser)

		assert.NoError(t, err)

		_, err = govalidator.ValidateStruct(tempMockRegisterUser)

		assert.Error(t, err)
		assert.Equal(t, mockRegisteredUser.ID, tempMockRegisterUser.ID)
		assert.Equal(t, mockRegisteredUser.Age, tempMockRegisterUser.Age)
		assert.Equal(t, mockRegisteredUser.Email, tempMockRegisterUser.Email)
		assert.Equal(t, mockRegisteredUser.Password, tempMockRegisterUser.Password)
		assert.NotEqual(t, mockRegisteredUser.Username, tempMockRegisterUser.Username)
		mockUserRepository.AssertExpectations(t)
	})

	t.Run("register user with invalid email format", func(t *testing.T) {
		tempMockRegisterUser := domain.User{
			Age:      8,
			Email:    "johndoe",
			Password: "secret",
			Username: "johndoe",
		}

		tempMockRegisterUser.ID = "user-123"

		mockUserRepository.On("Register", mock.Anything, mock.AnythingOfType("*domain.User")).Return(nil).Once()

		err := userUseCase.Register(context.Background(), &tempMockRegisterUser)

		assert.NoError(t, err)

		_, err = govalidator.ValidateStruct(tempMockRegisterUser)

		assert.Error(t, err)
		assert.Equal(t, mockRegisteredUser.ID, tempMockRegisterUser.ID)
		assert.Equal(t, mockRegisteredUser.Age, tempMockRegisterUser.Age)
		assert.NotEqual(t, mockRegisteredUser.Email, tempMockRegisterUser.Email)
		assert.Equal(t, mockRegisteredUser.Password, tempMockRegisterUser.Password)
		assert.Equal(t, mockRegisteredUser.Username, tempMockRegisterUser.Username)
		mockUserRepository.AssertExpectations(t)
	})

	t.Run("register user with password under limit character", func(t *testing.T) {
		tempMockRegisterUser := domain.User{
			Age:      8,
			Email:    "johndoe@example.com",
			Password: "scrt",
			Username: "johndoe",
		}

		tempMockRegisterUser.ID = "user-123"

		mockUserRepository.On("Register", mock.Anything, mock.AnythingOfType("*domain.User")).Return(nil).Once()

		err := userUseCase.Register(context.Background(), &tempMockRegisterUser)

		assert.NoError(t, err)

		_, err = govalidator.ValidateStruct(tempMockRegisterUser)

		assert.Error(t, err)
		assert.Equal(t, mockRegisteredUser.ID, tempMockRegisterUser.ID)
		assert.Equal(t, mockRegisteredUser.Age, tempMockRegisterUser.Age)
		assert.Equal(t, mockRegisteredUser.Email, tempMockRegisterUser.Email)
		assert.NotEqual(t, mockRegisteredUser.Password, tempMockRegisterUser.Password)
		assert.Equal(t, mockRegisteredUser.Username, tempMockRegisterUser.Username)
		mockUserRepository.AssertExpectations(t)
	})

	t.Run("register user with age under limit number", func(t *testing.T) {
		tempMockRegisterUser := domain.User{
			Age:      7,
			Email:    "johndoe@example.com",
			Password: "secret",
			Username: "johndoe",
		}

		tempMockRegisterUser.ID = "user-123"

		mockUserRepository.On("Register", mock.Anything, mock.AnythingOfType("*domain.User")).Return(nil).Once()

		err := userUseCase.Register(context.Background(), &tempMockRegisterUser)

		assert.NoError(t, err)

		_, err = govalidator.ValidateStruct(tempMockRegisterUser)

		assert.Error(t, err)
		assert.Equal(t, mockRegisteredUser.ID, tempMockRegisterUser.ID)
		assert.NotEqual(t, mockRegisteredUser.Age, tempMockRegisterUser.Age)
		assert.Equal(t, mockRegisteredUser.Email, tempMockRegisterUser.Email)
		assert.Equal(t, mockRegisteredUser.Password, tempMockRegisterUser.Password)
		assert.Equal(t, mockRegisteredUser.Username, tempMockRegisterUser.Username)
		mockUserRepository.AssertExpectations(t)
	})

	t.Run("register user with not contain needed property", func(t *testing.T) {
		tempMockRegisterUser := domain.User{
			Age:   8,
			Email: "johndoe@example.com",
		}

		tempMockRegisterUser.ID = "user-123"

		mockUserRepository.On("Register", mock.Anything, mock.AnythingOfType("*domain.User")).Return(nil).Once()

		err := userUseCase.Register(context.Background(), &tempMockRegisterUser)

		assert.NoError(t, err)

		_, err = govalidator.ValidateStruct(tempMockRegisterUser)

		assert.Error(t, err)
		assert.Equal(t, mockRegisteredUser.ID, tempMockRegisterUser.ID)
		assert.Equal(t, mockRegisteredUser.Age, tempMockRegisterUser.Age)
		assert.Equal(t, mockRegisteredUser.Email, tempMockRegisterUser.Email)
		assert.NotEqual(t, mockRegisteredUser.Password, tempMockRegisterUser.Password)
		assert.NotEqual(t, mockRegisteredUser.Username, tempMockRegisterUser.Username)

		mockUserRepository.AssertExpectations(t)
	})
}

func TestLogin(t *testing.T) {
	mockRegisteredUser := domain.User{
		ID:       "user-123",
		Age:      8,
		Email:    "johndoe@example.com",
		Password: helpers.Hash("secret"),
		Username: "johndoe",
	}

	mockUserRepository := new(mocks.UserRepository)
	userUseCase := userUseCase.NewUserUseCase(mockUserRepository)

	t.Run("login user correctly", func(t *testing.T) {
		tempMockLoginUser := domain.User{
			Email:    "johndoe@example.com",
			Password: "secret",
		}

		mockUserRepository.On("Login", mock.Anything, mock.AnythingOfType("*domain.User")).Return(nil).Once()

		err := userUseCase.Login(context.Background(), &tempMockLoginUser)

		assert.NoError(t, err)

		assert.Equal(t, tempMockLoginUser.Email, mockRegisteredUser.Email)

		isValid := helpers.Compare([]byte(mockRegisteredUser.Password), []byte(tempMockLoginUser.Password))

		assert.True(t, isValid)
		mockUserRepository.AssertExpectations(t)
	})

	t.Run("login user with not registered email", func(t *testing.T) {
		tempMockLoginUser := domain.User{
			Email:    "lorem@example.com",
			Password: "secret",
		}

		mockUserRepository.On("Login", mock.Anything, mock.AnythingOfType("*domain.User")).Return(nil).Once()

		err := userUseCase.Login(context.Background(), &tempMockLoginUser)

		assert.NoError(t, err)

		assert.NotEqual(t, tempMockLoginUser.Email, mockRegisteredUser.Email)

		isValid := helpers.Compare([]byte(mockRegisteredUser.Password), []byte(tempMockLoginUser.Password))

		assert.True(t, isValid)
		mockUserRepository.AssertExpectations(t)
	})

	t.Run("login user with invalid password", func(t *testing.T) {
		tempMockLoginUser := domain.User{
			Email:    "johndoe@example.com",
			Password: "scrt",
		}

		mockUserRepository.On("Login", mock.Anything, mock.AnythingOfType("*domain.User")).Return(nil).Once()

		err := userUseCase.Login(context.Background(), &tempMockLoginUser)

		assert.NoError(t, err)

		assert.Equal(t, tempMockLoginUser.Email, mockRegisteredUser.Email)

		isValid := helpers.Compare([]byte(mockRegisteredUser.Password), []byte(tempMockLoginUser.Password))

		assert.False(t, isValid)
		mockUserRepository.AssertExpectations(t)
	})

	t.Run("login user with empty email", func(t *testing.T) {
		tempMockLoginUser := domain.User{
			Email:    "",
			Password: "secret",
		}

		mockUserRepository.On("Login", mock.Anything, mock.AnythingOfType("*domain.User")).Return(nil).Once()

		err := userUseCase.Login(context.Background(), &tempMockLoginUser)

		assert.NoError(t, err)

		assert.NotEqual(t, tempMockLoginUser.Email, mockRegisteredUser.Email)

		isValid := helpers.Compare([]byte(mockRegisteredUser.Password), []byte(tempMockLoginUser.Password))

		assert.True(t, isValid)
		mockUserRepository.AssertExpectations(t)
	})

	t.Run("login user with empty password", func(t *testing.T) {
		tempMockLoginUser := domain.User{
			Email:    "johndoe@example.com",
			Password: "",
		}

		mockUserRepository.On("Login", mock.Anything, mock.AnythingOfType("*domain.User")).Return(nil).Once()

		err := userUseCase.Login(context.Background(), &tempMockLoginUser)

		assert.NoError(t, err)

		assert.Equal(t, tempMockLoginUser.Email, mockRegisteredUser.Email)

		isValid := helpers.Compare([]byte(mockRegisteredUser.Password), []byte(tempMockLoginUser.Password))

		assert.False(t, isValid)
		mockUserRepository.AssertExpectations(t)
	})
}

func TestUpdate(t *testing.T) {
	now := time.Now()
	mockUpdatedUser := domain.User{
		ID:        "user-123",
		Email:     "newjohndoe@example.com",
		Username:  "newjohndoe",
		Age:       8,
		Password:  "secret",
		UpdatedAt: &now,
	}

	mockUserRepository := new(mocks.UserRepository)
	userUseCase := userUseCase.NewUserUseCase(mockUserRepository)

	t.Run("update user correctly", func(t *testing.T) {
		tempMockUpdateUser := domain.User{
			Email:    "newjohndoe@example.com",
			Username: "newjohndoe",
		}

		mockUserRepository.On("Update", mock.Anything, mock.AnythingOfType("domain.User")).Return(mockUpdatedUser, nil).Once()

		user, err := userUseCase.Update(context.Background(), tempMockUpdateUser)

		assert.NoError(t, err)

		tempMockUpdatedUser := domain.User{
			ID:        "user-123",
			Email:     tempMockUpdateUser.Email,
			Username:  tempMockUpdateUser.Username,
			Age:       8,
			Password:  "secret",
			UpdatedAt: &now,
		}

		_, err = govalidator.ValidateStruct(tempMockUpdatedUser)

		assert.NoError(t, err)
		assert.Equal(t, user, tempMockUpdatedUser)
		assert.Equal(t, mockUpdatedUser.Email, tempMockUpdateUser.Email)
		assert.Equal(t, mockUpdatedUser.Username, tempMockUpdatedUser.Username)
		mockUserRepository.AssertExpectations(t)
	})

	t.Run("update user with empty email", func(t *testing.T) {
		tempMockUpdateUser := domain.User{
			Email:    "",
			Username: "newjohndoe",
		}

		mockUserRepository.On("Update", mock.Anything, mock.AnythingOfType("domain.User")).Return(mockUpdatedUser, nil).Once()

		user, err := userUseCase.Update(context.Background(), tempMockUpdateUser)

		assert.NoError(t, err)

		tempMockUpdatedUser := domain.User{
			ID:        "user-123",
			Email:     tempMockUpdateUser.Email,
			Username:  tempMockUpdateUser.Username,
			Age:       8,
			Password:  "secret",
			UpdatedAt: &now,
		}

		_, err = govalidator.ValidateStruct(tempMockUpdatedUser)

		assert.Error(t, err)
		assert.NotEqual(t, user, tempMockUpdatedUser)
		assert.NotEqual(t, mockUpdatedUser.Email, tempMockUpdatedUser.Email)
		assert.Equal(t, mockUpdatedUser.Username, tempMockUpdateUser.Username)
		mockUserRepository.AssertExpectations(t)
	})

	t.Run("update user with empty username", func(t *testing.T) {
		tempMockUpdateUser := domain.User{
			Email:    "newjohndoe@example.com",
			Username: "",
		}

		mockUserRepository.On("Update", mock.Anything, mock.AnythingOfType("domain.User")).Return(mockUpdatedUser, nil).Once()

		user, err := userUseCase.Update(context.Background(), tempMockUpdateUser)

		assert.NoError(t, err)

		tempMockUpdatedUser := domain.User{
			ID:        "user-123",
			Email:     tempMockUpdateUser.Email,
			Username:  tempMockUpdateUser.Username,
			Age:       8,
			Password:  "secret",
			UpdatedAt: &now,
		}

		_, err = govalidator.ValidateStruct(tempMockUpdatedUser)

		assert.Error(t, err)
		assert.NotEqual(t, user, tempMockUpdatedUser)
		assert.Equal(t, mockUpdatedUser.Email, tempMockUpdatedUser.Email)
		assert.NotEqual(t, mockUpdatedUser.Username, tempMockUpdatedUser.Username)
		mockUserRepository.AssertExpectations(t)
	})

	t.Run("update user with empty username and email", func(t *testing.T) {
		tempMockUpdateUser := domain.User{
			Email:    "",
			Username: "",
		}

		mockUserRepository.On("Update", mock.Anything, mock.AnythingOfType("domain.User")).Return(mockUpdatedUser, nil).Once()

		user, err := userUseCase.Update(context.Background(), tempMockUpdateUser)

		assert.NoError(t, err)

		tempMockUpdatedUser := domain.User{
			ID:        "user-123",
			Email:     tempMockUpdateUser.Email,
			Username:  tempMockUpdateUser.Username,
			Age:       8,
			Password:  "secret",
			UpdatedAt: &now,
		}

		assert.NoError(t, err)

		_, err = govalidator.ValidateStruct(tempMockUpdatedUser)

		assert.Error(t, err)
		assert.NotEqual(t, user, tempMockUpdatedUser)
		assert.NotEqual(t, mockUpdatedUser.Email, tempMockUpdatedUser.Email)
		assert.NotEqual(t, mockUpdatedUser.Username, tempMockUpdatedUser.Username)
		mockUserRepository.AssertExpectations(t)
	})

	t.Run("update user with invalid email format", func(t *testing.T) {
		tempMockUpdateUser := domain.User{
			Email:    "newjohndoe",
			Username: "newjohndoe",
		}

		mockUserRepository.On("Update", mock.Anything, mock.AnythingOfType("domain.User")).Return(mockUpdatedUser, nil).Once()

		user, err := userUseCase.Update(context.Background(), tempMockUpdateUser)

		assert.NoError(t, err)

		tempMockUpdatedUser := domain.User{
			ID:        "user-123",
			Email:     tempMockUpdateUser.Email,
			Username:  tempMockUpdateUser.Username,
			Age:       8,
			Password:  "secret",
			UpdatedAt: &now,
		}

		assert.NoError(t, err)

		_, err = govalidator.ValidateStruct(tempMockUpdatedUser)

		assert.Error(t, err)
		assert.NotEqual(t, user, tempMockUpdatedUser)
		assert.NotEqual(t, mockUpdatedUser.Email, tempMockUpdatedUser.Email)
		assert.Equal(t, mockUpdatedUser.Username, tempMockUpdatedUser.Username)
		mockUserRepository.AssertExpectations(t)
	})
}

func TestDelete(t *testing.T) {
	mockUser := domain.User{
		ID:       "user-123",
		Age:      8,
		Email:    "johndoe@example.com",
		Password: "secret",
		Username: "johndoe",
	}

	mockUserRepository := new(mocks.UserRepository)
	userUseCase := userUseCase.NewUserUseCase(mockUserRepository)

	t.Run("delete user correctly", func(t *testing.T) {
		mockUserRepository.On("Delete", mock.Anything, mock.AnythingOfType("string")).Return(nil).Once()

		err := userUseCase.Delete(context.Background(), mockUser.ID)

		assert.NoError(t, err)
		mockUserRepository.AssertExpectations(t)
	})

	t.Run("delete user with not found user", func(t *testing.T) {
		mockUserRepository.On("Delete", mock.Anything, mock.AnythingOfType("string")).Return(errors.New("fail")).Once()

		err := userUseCase.Delete(context.Background(), "user-234")

		assert.Error(t, err)
		mockUserRepository.AssertExpectations(t)
	})
}
