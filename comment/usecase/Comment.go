package usecase

import (
	validator "github.com/asaskevich/govalidator"

	"gorgs/comment/model"
	"gorgs/comment/repository/sql"
)

type CommentUseCaseInterface interface {
	Create(orgName string, request model.CommentRequest) (model.Comment, error)
	GetByOrgName(orgName string) ([]model.Comment, error)
	GetOneByOrgNameAndId(orgName string, ID uint) (model.Comment, error)
	Update(orgName string, ID uint, request model.CommentRequest) (model.Comment, error)
	Delete(orgName string, ID uint) error
}
type commentUseCase struct {
	repo sql.CommentRepoInterface
}

func NewCommentUseCase(_repo sql.CommentRepoInterface) CommentUseCaseInterface {
	return &commentUseCase{repo: _repo}
}
func (c *commentUseCase) Create(orgName string, request model.CommentRequest) (model.Comment, error) {
	_, err := validator.ValidateStruct(request)
	if err != nil {
		return model.Comment{}, err
	}
	input := model.Comment{
		OrgName: orgName,
		Comment: request.Comment,
	}
	return c.repo.Create(input)
}

func (c *commentUseCase) GetByOrgName(orgName string) ([]model.Comment, error) {
	return c.repo.GetByOrgName(orgName)
}

func (c *commentUseCase) GetOneByOrgNameAndId(orgName string, ID uint) (model.Comment, error) {
	return c.repo.GetOneByOrgNameAndId(orgName, ID)
}

func (c *commentUseCase) Update(orgName string, ID uint, request model.CommentRequest) (model.Comment, error) {
	_, err := validator.ValidateStruct(request)
	if err != nil {
		return model.Comment{}, err
	}

	input := model.Comment{
		Comment: request.Comment,
	}
	return c.repo.Update(orgName, ID, input)
}

func (c *commentUseCase) Delete(orgName string, ID uint) error {
	return c.repo.Delete(orgName, ID)
}
