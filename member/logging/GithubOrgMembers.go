package logging

import (
	"github.com/sirupsen/logrus"
	"gorgs/member/model"

	"gorgs/member/usecase"
	"time"
)

// wrapper logging for use case or logic

type githubOrgMembersLog struct {
	name string
	u    usecase.GithubOrgMembersUseCaseInterface
}

func NewGithubOrgMembersLog(_u usecase.GithubOrgMembersUseCaseInterface) usecase.GithubOrgMembersUseCaseInterface {
	return &githubOrgMembersLog{
		name: "Member UseCase",
		u:    _u,
	}
}

func (g *githubOrgMembersLog) GetMembers(orgs string, page, perPage int) (data []model.GithubUser, err error) {

	defer func(begin time.Time) {
		logrus.WithFields(logrus.Fields{
			"method":  "GetMembers",
			"orgs":    orgs,
			"page":    page,
			"perPage": perPage,
			"took":    time.Since(begin),
			"err":     err,
		}).Info(g.name)
	}(time.Now())

	return g.u.GetMembers(orgs, page, perPage)
}
