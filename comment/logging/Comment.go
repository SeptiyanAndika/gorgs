package logging

import (
	"github.com/sirupsen/logrus"
	"gorgs/comment/model"
	"gorgs/comment/usecase"
	"time"
)

type commentLog struct {
	name string
	u    usecase.CommentUseCaseInterface
}

func NewCommentLog(_u usecase.CommentUseCaseInterface) usecase.CommentUseCaseInterface {
	return &commentLog{
		name: "Comment UseCase",
		u:    _u,
	}
}

func (c *commentLog) Create(orgName string, request model.CommentRequest) (data model.Comment, err error) {

	defer func(begin time.Time) {
		logrus.WithFields(logrus.Fields{
			"method":  "Create",
			"orgName": orgName,
			"request": request,
			"took":    time.Since(begin),
			"err":     err,
		}).Info(c.name)
	}(time.Now())

	return c.u.Create(orgName, request)
}

func (c *commentLog) GetByOrgName(orgName string) (data []model.Comment, err error) {

	defer func(begin time.Time) {
		logrus.WithFields(logrus.Fields{
			"method": "GetAll",
			"took":   time.Since(begin),
			"err":    err,
		}).Info(c.name)
	}(time.Now())

	return c.u.GetByOrgName(orgName)
}

func (c *commentLog) GetOneByOrgNameAndId(orgName string, ID uint) (data model.Comment, err error) {

	defer func(begin time.Time) {
		logrus.WithFields(logrus.Fields{
			"method":  "GetOneById",
			"orgName": orgName,
			"id":      ID,
			"took":    time.Since(begin),
			"err":     err,
		}).Info(c.name)
	}(time.Now())

	return c.u.GetOneByOrgNameAndId(orgName, ID)
}

func (c *commentLog) Update(orgName string, ID uint, newData model.CommentRequest) (data model.Comment, err error) {

	defer func(begin time.Time) {
		logrus.WithFields(logrus.Fields{
			"method":  "Update",
			"orgName": orgName,
			"id":      ID,
			"took":    time.Since(begin),
			"err":     err,
		}).Info(c.name)
	}(time.Now())

	return c.u.Update(orgName, ID, newData)
}

func (c *commentLog) Delete(orgName string, ID uint) (err error) {

	defer func(begin time.Time) {
		logrus.WithFields(logrus.Fields{
			"method":  "Delete",
			"orgName": orgName,
			"id":      ID,
			"took":    time.Since(begin),
			"err":     err,
		}).Info(c.name)
	}(time.Now())

	return c.u.Delete(orgName, ID)
}
