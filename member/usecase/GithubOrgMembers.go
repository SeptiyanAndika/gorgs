package usecase

import (
	"gorgs/member/adapter"
	"gorgs/member/model"
	"sort"
)

type GithubOrgMembersUseCaseInterface interface {
	GetMembers(orgs string, page, perPage int) ([]model.GithubUser, error)
}

func NewGithubAdapter(_adapter adapter.GithubAdapterInterface) GithubOrgMembersUseCaseInterface {
	return &githubOrgMembersUseCase{
		adapter: _adapter,
	}
}

type githubOrgMembersUseCase struct {
	adapter adapter.GithubAdapterInterface
}

func (g *githubOrgMembersUseCase) GetMembers(orgs string, page, perPage int) ([]model.GithubUser, error) {
	var users []model.GithubUser
	ch := make(chan model.GithubUser)

	members, err := g.adapter.GetMembers(orgs, page, perPage)
	if err != nil {
		return users, err
	}

	for _, val := range members {
		go g.chanDetailUser(val.Login, ch)
	}

	for range members {
		user := <-ch
		users = append(users, user)
	}

	sort.Sort(model.GithubUserByFollowers(users))

	return users, nil

}

func (g *githubOrgMembersUseCase) chanDetailUser(login string, ch chan<- model.GithubUser) {
	user, _ := g.adapter.DetailUser(login)
	ch <- user
}
