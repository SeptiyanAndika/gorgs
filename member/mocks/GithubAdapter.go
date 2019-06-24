package mocks

import (
	"github.com/stretchr/testify/mock"
	"gorgs/member/model"
)

type GithubAdapter struct {
	mock.Mock
}

func (_mock *GithubAdapter) GetMembers(orgs string, page, perPage int) ([]model.GithubMember, error) {
	ret := _mock.Called(orgs, page, perPage)
	members := ret.Get(0).([]model.GithubMember)
	err := ret.Error(1)
	return members, err
}

func (_mock *GithubAdapter) DetailUser(login string) (model.GithubUser, error) {
	ret := _mock.Called(login)
	user := ret.Get(0).(model.GithubUser)
	err := ret.Error(1)
	return user, err
}
