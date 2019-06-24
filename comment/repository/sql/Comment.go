package sql

import (
	"github.com/jinzhu/gorm"
	"gorgs/comment/model"
)

// interface repository for abstractions comment repo function
type CommentRepoInterface interface {
	Create(input model.Comment) (model.Comment, error)
	GetByOrgName(orgName string) ([]model.Comment, error)
	GetOneByOrgNameAndId(orgName string, ID uint) (model.Comment, error)
	Update(orgName string, ID uint, newData model.Comment) (model.Comment, error)
	Delete(orgName string, ID uint) error
}

type commentRepo struct {
	db *gorm.DB
}

func NewCommentRepo(_db *gorm.DB) CommentRepoInterface {
	return &commentRepo{db: _db}
}

func (c *commentRepo) Create(input model.Comment) (model.Comment, error) {
	err := c.db.Create(&input).Error
	if err != nil {
		return model.Comment{}, err
	}
	return input, err
}

func (c *commentRepo) GetByOrgName(orgName string) ([]model.Comment, error) {
	var comments []model.Comment
	err := c.db.Where("org_name = ?", orgName).Find(&comments).Error
	if err != nil {
		return comments, err
	}
	return comments, err
}

func (c *commentRepo) GetOneByOrgNameAndId(orgName string, ID uint) (model.Comment, error) {
	var comment model.Comment
	err := c.db.Where("id = ? and org_name = ? ", ID, orgName).First(&comment).Error
	if err != nil {
		return comment, err
	}
	return comment, err
}

func (c *commentRepo) Update(orgName string, ID uint, newData model.Comment) (model.Comment, error) {
	var comment model.Comment
	comment.ID = ID
	comment.OrgName = orgName

	err := c.db.Model(&comment).Update(newData).Error
	if err != nil {
		return comment, err
	}
	return comment, err
}

func (c *commentRepo) Delete(orgName string, ID uint) error {
	var comment model.Comment
	comment.ID = ID
	err := c.db.Delete(comment).Error
	if err != nil {
		return err
	}
	return nil
}
