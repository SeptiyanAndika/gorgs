package test

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"gorgs/comment/mocks"
	"gorgs/comment/model"
	"gorgs/comment/usecase"
	"testing"
	"time"
)

func TestGetByOrgName(t *testing.T) {
	mockCommentRepo := new(mocks.CommentRepoSql)

	mockComment1 := model.Comment{
		ID:        1,
		OrgName:   "expressjs",
		Comment:   "qazwsx123",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	mockComment2 := model.Comment{
		ID:        2,
		OrgName:   "expressjs",
		Comment:   "abcdef",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	var mockListComment []model.Comment
	mockListComment = append(mockListComment, mockComment1)
	mockListComment = append(mockListComment, mockComment2)

	t.Run("success", func(t *testing.T) {
		mockCommentRepo.On("GetByOrgName", "expressjs").Return(mockListComment, nil).Once()

		useCase := usecase.NewCommentUseCase(mockCommentRepo)
		res, err := useCase.GetByOrgName("expressjs")

		assert.NoError(t, err)
		assert.NotEmpty(t, res)
		assert.Len(t, res, 2)
	})

	t.Run("error", func(t *testing.T) {
		mockCommentRepo.On("GetByOrgName", "expressjs").Return([]model.Comment{}, errors.New("unexpected Error")).Once()
		useCase := usecase.NewCommentUseCase(mockCommentRepo)
		res, err := useCase.GetByOrgName("expressjs")
		assert.Error(t, err)
		assert.Empty(t, res)
	})

}

func TestGetOneByOrgNameAndId(t *testing.T) {
	mockCommentRepo := new(mocks.CommentRepoSql)

	mockComment := model.Comment{
		ID:        1,
		OrgName:   "expressjs",
		Comment:   "qazwsx123",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	t.Run("success", func(t *testing.T) {
		mockCommentRepo.On("GetOneByOrgNameAndId", "expressjs", uint(1)).Return(mockComment, nil).Once()

		useCase := usecase.NewCommentUseCase(mockCommentRepo)
		res, err := useCase.GetOneByOrgNameAndId("expressjs", 1)

		assert.NoError(t, err)
		assert.NotEmpty(t, res)
	})

	t.Run("error", func(t *testing.T) {
		mockCommentRepo.On("GetOneByOrgNameAndId", "expressjs", uint(1)).Return(model.Comment{}, errors.New("unexpected Error")).Once()
		useCase := usecase.NewCommentUseCase(mockCommentRepo)
		res, err := useCase.GetOneByOrgNameAndId("expressjs", 1)
		assert.Error(t, err)
		assert.Empty(t, res)
	})

}

func TestCreate(t *testing.T) {
	mockCommentRepo := new(mocks.CommentRepoSql)

	mockCommentResult := model.Comment{
		OrgName: "expressjs",
		Comment: "qazwsx123",
	}
	mockCommentRequest := model.CommentRequest{
		Comment: "qazwsx123",
	}

	t.Run("success", func(t *testing.T) {
		mockCommentRepo.On("Create", mockCommentResult).Return(mockCommentResult, nil).Once()
		useCase := usecase.NewCommentUseCase(mockCommentRepo)
		res, err := useCase.Create("expressjs", mockCommentRequest)

		assert.NoError(t, err)
		assert.NotEmpty(t, res)
		assert.Equal(t, res.ID, uint(1))
		assert.Equal(t, res.OrgName,"expressjs")
		assert.Equal(t, mockCommentRequest.Comment, res.Comment)
	})

	t.Run("error", func(t *testing.T) {
		mockCommentRepo.On("Create", mockCommentResult).Return(model.Comment{}, errors.New("unexpected Error")).Once()
		useCase := usecase.NewCommentUseCase(mockCommentRepo)
		res, err := useCase.Create("expressjs", mockCommentRequest)
		assert.Error(t, err)
		assert.Empty(t, res)
	})

}

func TestUpdate(t *testing.T) {
	mockCommentRepo := new(mocks.CommentRepoSql)

	mockCommentResult := model.Comment{
		Comment: "qazwsx123",
	}
	mockCommentRequest := model.CommentRequest{
		Comment: "qazwsx123",
	}

	t.Run("success", func(t *testing.T) {
		mockCommentRepo.On("Update", "expressjs", uint(1), mockCommentResult).Return(mockCommentResult, nil).Once()
		useCase := usecase.NewCommentUseCase(mockCommentRepo)
		res, err := useCase.Update("expressjs", 1, mockCommentRequest)

		assert.NoError(t, err)
		assert.NotEmpty(t, res)
		assert.Equal(t, res.ID, uint(1))
		assert.Equal(t, mockCommentRequest.Comment, res.Comment)
		assert.Equal(t, res.OrgName,"expressjs")
	})

	t.Run("error", func(t *testing.T) {
		mockCommentRepo.On("Update", "expressjs", uint(1), mockCommentResult).Return(model.Comment{}, errors.New("unexpected Error")).Once()
		useCase := usecase.NewCommentUseCase(mockCommentRepo)
		res, err := useCase.Update("expressjs", 1, mockCommentRequest)
		assert.Error(t, err)
		assert.Empty(t, res)
	})
}

func TestDelete(t *testing.T) {
	mockCommentRepo := new(mocks.CommentRepoSql)

	t.Run("success", func(t *testing.T) {
		mockCommentRepo.On("Delete", "expressjs", uint(1)).Return(nil).Once()
		useCase := usecase.NewCommentUseCase(mockCommentRepo)
		err := useCase.Delete("expressjs", 1)

		assert.NoError(t, err)
	})

	t.Run("error", func(t *testing.T) {
		mockCommentRepo.On("Delete", "expressjs", uint(1)).Return(errors.New("unexpected Error")).Once()
		useCase := usecase.NewCommentUseCase(mockCommentRepo)
		err := useCase.Delete("expressjs", 1)
		assert.Error(t, err)
	})
}
