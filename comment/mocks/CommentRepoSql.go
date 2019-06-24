package mocks

import (
	"github.com/stretchr/testify/mock"
	"gorgs/comment/model"
)

type CommentRepoSql struct {
	mock.Mock
}

func (_mock *CommentRepoSql) Create(input model.Comment) (model.Comment, error) {
	ret := _mock.Called(input)
	data := ret.Get(0).(model.Comment)

	data.ID = 1
	data.OrgName = input.OrgName
	data.Comment = input.Comment

	err := ret.Error(1)
	if err != nil {
		return model.Comment{}, err
	}
	return data, err
}

func (_mock *CommentRepoSql) GetByOrgName(orgName string) ([]model.Comment, error) {
	ret := _mock.Called(orgName)
	data := ret.Get(0).([]model.Comment)
	err := ret.Error(1)
	return data, err
}

func (_mock *CommentRepoSql) GetOneByOrgNameAndId(orgName string, ID uint) (model.Comment, error) {
	ret := _mock.Called(orgName, ID)
	data := ret.Get(0).(model.Comment)
	err := ret.Error(1)
	return data, err
}

func (_mock *CommentRepoSql) Update(orgName string, ID uint, newData model.Comment) (model.Comment, error) {
	ret := _mock.Called(orgName, ID, newData)
	data := ret.Get(0).(model.Comment)
	data.ID = 1
	data.OrgName = orgName
	data.Comment = newData.Comment

	err := ret.Error(1)
	if err != nil {
		return model.Comment{}, err
	}
	return data, err
}

func (_mock *CommentRepoSql) Delete(orgName string, ID uint) error {
	ret := _mock.Called(orgName, ID)
	err := ret.Error(0)
	return err
}
