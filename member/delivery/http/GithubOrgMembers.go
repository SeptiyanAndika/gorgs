package http

import (
	"github.com/labstack/echo/v4"
	"gorgs/member/usecase"
	"strconv"
)

type GithubOrgMembers struct {
	gomUsecase usecase.GithubOrgMembersUseCaseInterface
}

func NewGithubOrgMembers(e *echo.Echo, _gomUsecase usecase.GithubOrgMembersUseCaseInterface) {
	handler := &GithubOrgMembers{
		gomUsecase: _gomUsecase,
	}
	e.GET("/:orgs", handler.GetMembers)
}

func (g GithubOrgMembers) GetMembers(ctx echo.Context) error {
	orgs := ctx.Param("orgs")

	page, err := strconv.Atoi(ctx.QueryParam("page"))
	if err != nil {
		page = 1
	}

	perPage, err := strconv.Atoi(ctx.QueryParam("per_page"))
	if err != nil {
		perPage = 10
	}

	results, err := g.gomUsecase.GetMembers(orgs, page, perPage)

	if err != nil {
		return errorResponse(ctx, err, nil)
	}

	return successResponse(ctx, "", results)
}
