package adapter

import (
	"errors"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"gorgs/member/model"
)

type GithubAdapterInterface interface {
	GetMembers(orgs string, page, perPage int) ([]model.GithubMember, error)
	DetailUser(login string) (model.GithubUser, error)
}

func NewGithubAdapter() GithubAdapterInterface {
	return &githubAdapter{}
}

type githubAdapter struct {
	HttpRequest
}

func (g githubAdapter) GetMembers(orgs string, page, perPage int) ([]model.GithubMember, error) {
	var members []model.GithubMember
	var genError model.GithubGeneralError

	url := fmt.Sprintf("https://api.github.com/orgs/%s/members?page=%d&per_page=%d", orgs, page, perPage)
	response, err := g.HttpRequest.Get(url, nil, nil)
	if err != nil {
		return members, err
	}
	mapstructure.Decode(response.Body, &genError)
	if genError.Message != "" {
		return members, errors.New(genError.Message)
	}

	err = g.HttpRequest.DecodeBody(response.Body, &members)
	if err != nil {
		return members, err
	}

	return members, nil
}

func (g githubAdapter) DetailUser(login string) (model.GithubUser, error) {
	var user model.GithubUser
	var genError model.GithubGeneralError

	url := fmt.Sprintf("https://api.github.com/users/%s", login)
	response, err := g.HttpRequest.Get(url, nil, nil)
	if err != nil {
		return user, err
	}
	mapstructure.Decode(response.Body, &genError)
	if genError.Message != "" {
		return user, errors.New(genError.Message)
	}

	err = g.HttpRequest.DecodeBody(response.Body, &user)
	if err != nil {
		return user, err
	}

	return user, nil
}
