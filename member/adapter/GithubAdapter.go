package adapter

import (
	"errors"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"gorgs/member/model"
)

// Interface for github adapter for abstraction
type GithubAdapterInterface interface {
	GetMembers(orgs string, page, perPage int) ([]model.GithubMember, error)
	DetailUser(login string) (model.GithubUser, error)
}

// New object for github adapter
func NewGithubAdapter() GithubAdapterInterface {
	return &githubAdapter{}
}

// struct or class github adapter
type githubAdapter struct {
	HttpRequest
}

// get members from github api with organisation name param
func (g githubAdapter) GetMembers(orgs string, page, perPage int) ([]model.GithubMember, error) {
	var members []model.GithubMember
	var genError model.GithubGeneralError

	url := fmt.Sprintf("https://api.github.com/orgs/%s/members?page=%d&per_page=%d", orgs, page, perPage)
	response, err := g.HttpRequest.Get(url, nil, nil)
	if err != nil {
		return members, err
	}

	// map response body to general error, if have error message return error
	mapstructure.Decode(response.Body, &genError)
	if genError.Message != "" {
		return members, errors.New(genError.Message)
	}

	// map response to struct members
	err = g.HttpRequest.DecodeBody(response.Body, &members)
	if err != nil {
		return members, err
	}

	return members, nil
}

// get detail user from github api witt login id params
func (g githubAdapter) DetailUser(login string) (model.GithubUser, error) {
	var user model.GithubUser
	var genError model.GithubGeneralError

	url := fmt.Sprintf("https://api.github.com/users/%s", login)
	response, err := g.HttpRequest.Get(url, nil, nil)
	if err != nil {
		return user, err
	}

	// map response body to general error, if have error message return error
	mapstructure.Decode(response.Body, &genError)
	if genError.Message != "" {
		return user, errors.New(genError.Message)
	}

	// map response to struct user
	err = g.HttpRequest.DecodeBody(response.Body, &user)
	if err != nil {
		return user, err
	}

	return user, nil
}
