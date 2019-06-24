package test

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"gorgs/member/mocks"
	"gorgs/member/model"
	"gorgs/member/usecase"
	"testing"
)

func TestGetMembers(t *testing.T) {
	mockGithubAdapter := new(mocks.GithubAdapter)

	mockGithubMember := model.GithubMember{
		Login:     "andrewrk",
		ID:        106511,
		AvatarURL: "https://avatars0.githubusercontent.com/u/106511?v=4",
	}

	mockGithubUser := model.GithubUser{
		Login:     "andrewrk",
		AvatarURL: "https://avatars0.githubusercontent.com/u/106511?v=4",
		Followers: 1187,
		Following: 192,
	}

	var mockListGithubMember []model.GithubMember
	mockListGithubMember = append(mockListGithubMember, mockGithubMember)

	t.Run("success", func(t *testing.T) {
		mockGithubAdapter.On("GetMembers", "expressjs", 1, 1).Return(mockListGithubMember, nil).Once()
		mockGithubAdapter.On("DetailUser", "andrewrk").Return(mockGithubUser, nil).Once()

		useCase := usecase.NewGithubOrgMemberUsecase(mockGithubAdapter)
		res, err := useCase.GetMembers("expressjs", 1, 1)

		assert.NoError(t, err)
		assert.NotEmpty(t, res)
		assert.Len(t, res, 1)
		assert.Equal(t, res[0].Followers, 1187)
		assert.Equal(t, res[0].Following, 192)
	})

	t.Run("Error", func(t *testing.T) {
		mockGithubAdapter.On("GetMembers", "expressjs", 1, 1).Return(mockListGithubMember, errors.New("unexpected Error")).Once()
		mockGithubAdapter.On("DetailUser", "andrewrk").Return(mockGithubUser, errors.New("unexpected Error")).Once()

		useCase := usecase.NewGithubOrgMemberUsecase(mockGithubAdapter)
		_, err := useCase.GetMembers("expressjs", 1, 1)

		assert.Error(t, err)
	})

}
