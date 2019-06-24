package http

import (
	"github.com/labstack/echo/v4"
	"gorgs/comment/model"
	"gorgs/comment/usecase"
	"strconv"
)


type comment struct {
	commUseCase usecase.CommentUseCaseInterface
}

// create new object delivery comment with routes defined
func NewComment(e *echo.Echo, _commUseCase usecase.CommentUseCaseInterface) {
	handler := &comment{
		commUseCase: _commUseCase,
	}
	e.GET("/:orgs", handler.GetByOrgName)
	e.GET("/:orgs/:id", handler.GetOneByOrgNameAndId)
	e.POST("/:orgs", handler.Create)
	e.PUT("/:orgs/:id", handler.Update)
	e.DELETE("/:orgs/:id", handler.Delete)
}

func (co *comment) GetByOrgName(ctx echo.Context) error {
	orgs := ctx.Param("orgs")

	// call use case or logic get GetByOrgName
	results, err := co.commUseCase.GetByOrgName(orgs)
	if err != nil {
		return errorResponse(ctx, err, nil)
	}

	return successResponse(ctx, "", results)
}

func (co *comment) GetOneByOrgNameAndId(ctx echo.Context) error {
	orgs := ctx.Param("orgs")

	// if err convert id to int return error response
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return errorResponse(ctx, err, nil)
	}

	// call use case or logic get GetOneByOrgNameAndId
	results, err := co.commUseCase.GetOneByOrgNameAndId(orgs,uint(id))
	if err != nil {
		return errorResponse(ctx, err, nil)
	}

	return successResponse(ctx, "", results)
}
func (co *comment) Create(ctx echo.Context) error {
	orgs := ctx.Param("orgs")

	// binding json request to struct
	var request model.CommentRequest
	err := ctx.Bind(&request)
	if err != nil {
		return errorResponse(ctx, err, nil)
	}

	// call use case or logic get Create
	results, err := co.commUseCase.Create(orgs, request)
	if err != nil {
		return errorResponse(ctx, err, nil)
	}

	return successResponse(ctx, "", results)
}

func (co *comment) Update(ctx echo.Context) error {
	orgs := ctx.Param("orgs")

	// if err convert id to int return error response
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return errorResponse(ctx, err, nil)
	}

	//binding json request to struct
	var request model.CommentRequest
	err = ctx.Bind(&request)
	if err != nil {
		return errorResponse(ctx, err, nil)
	}

	// call use case or logic get Update
	results, err := co.commUseCase.Update(orgs, uint(id), request)
	if err != nil {
		return errorResponse(ctx, err, nil)
	}

	return successResponse(ctx, "", results)
}

func (co *comment) Delete(ctx echo.Context) error {
	orgs := ctx.Param("orgs")

	// if err convert id to int return error response
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return errorResponse(ctx, err, nil)
	}

	// call use case or logic get Delete
	err = co.commUseCase.Delete(orgs, uint(id))
	if err != nil {
		return errorResponse(ctx, err, nil)
	}

	return successResponse(ctx, "success delete", nil)
}
