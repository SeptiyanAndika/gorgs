package usecase

import (
	"gorgs/member/adapter"
	"gorgs/member/model"
	"sort"
)

// interface GithubOrgMembers for abstractions use case or logic function
type GithubOrgMembersUseCaseInterface interface {
	GetMembers(orgs string, page, perPage int) ([]model.GithubUser, error)
}

// New object for githubOrgMembersUseCase
func NewGithubOrgMemberUsecase(_adapter adapter.GithubAdapterInterface) GithubOrgMembersUseCaseInterface {
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

	// call get members adapter
	members, err := g.adapter.GetMembers(orgs, page, perPage)
	if err != nil {
		return users, err
	}

	// async get detail user
	for _, val := range members {
		go g.chanDetailUser(val.Login, ch)
	}

	// using chain to waiting all prosess async done.
	for range members {
		user := <-ch
		// append result from async to array of user
		users = append(users, user)
	}

	// sort array by followers
	sort.Sort(model.GithubUserByFollowers(users))

	return users, nil

}

// wrapper chat for adapter detail user
func (g *githubOrgMembersUseCase) chanDetailUser(login string, ch chan<- model.GithubUser) {
	user, _ := g.adapter.DetailUser(login)
	ch <- user
}
