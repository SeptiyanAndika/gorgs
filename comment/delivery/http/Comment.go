package http

import (
	"github.com/labstack/echo/v4"
	"gorgs/comment/model"
	"gorgs/comment/usecase"
	"strconv"
)

type Comment struct {
	commUseCase usecase.CommentUseCaseInterface
}

func NewComment(e *echo.Echo, _commUseCase usecase.CommentUseCaseInterface) {
	handler := &Comment{
		commUseCase: _commUseCase,
	}
	e.GET("/:orgs", handler.GetByOrgName)
	e.GET("/:orgs/:id", handler.GetOneByOrgNameAndId)
	e.POST("/:orgs", handler.Create)
	e.PUT("/:orgs/:id", handler.Update)
	e.DELETE("/:orgs/:id", handler.Delete)
}

func (co *Comment) GetByOrgName(ctx echo.Context) error {
	orgs := ctx.Param("orgs")

	results, err := co.commUseCase.GetByOrgName(orgs)

	if err != nil {
		return errorResponse(ctx, err, nil)
	}

	return successResponse(ctx, "", results)
}

func (co *Comment) GetOneByOrgNameAndId(ctx echo.Context) error {
	orgs := ctx.Param("orgs")

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return errorResponse(ctx, err, nil)
	}

	results, err := co.commUseCase.GetOneByOrgNameAndId(orgs,uint(id))

	if err != nil {
		return errorResponse(ctx, err, nil)
	}

	return successResponse(ctx, "", results)
}
func (co *Comment) Create(ctx echo.Context) error {
	orgs := ctx.Param("orgs")

	var request model.CommentRequest
	err := ctx.Bind(&request)
	if err != nil {
		return errorResponse(ctx, err, nil)
	}
	results, err := co.commUseCase.Create(orgs, request)
	if err != nil {
		return errorResponse(ctx, err, nil)
	}
	return successResponse(ctx, "", results)
}

func (co *Comment) Update(ctx echo.Context) error {
	orgs := ctx.Param("orgs")

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return errorResponse(ctx, err, nil)
	}

	var request model.CommentRequest
	err = ctx.Bind(&request)

	if err != nil {
		return errorResponse(ctx, err, nil)
	}

	results, err := co.commUseCase.Update(orgs, uint(id), request)

	if err != nil {
		return errorResponse(ctx, err, nil)
	}

	return successResponse(ctx, "", results)
}

func (co *Comment) Delete(ctx echo.Context) error {
	orgs := ctx.Param("orgs")

	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return errorResponse(ctx, err, nil)
	}

	err = co.commUseCase.Delete(orgs, uint(id))

	if err != nil {
		return errorResponse(ctx, err, nil)
	}

	return successResponse(ctx, "success delete", nil)
}
