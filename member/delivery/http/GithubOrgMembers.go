package http

import (
	"github.com/labstack/echo/v4"
	"gorgs/member/usecase"
	"strconv"
)


type githubOrgMembers struct {
	gomUseCase usecase.GithubOrgMembersUseCaseInterface
}

// create new object delivery githubOrgMembers with routes defined
func NewGithubOrgMembers(e *echo.Echo, _gomUseCase usecase.GithubOrgMembersUseCaseInterface) {
	handler := &githubOrgMembers{
		gomUseCase: _gomUseCase,
	}
	e.GET("/:orgs", handler.GetMembers)
}

func (g githubOrgMembers) GetMembers(ctx echo.Context) error {
	orgs := ctx.Param("orgs")

	// if empty or error get query args page, set default to 1
	page, err := strconv.Atoi(ctx.QueryParam("page"))
	if err != nil {
		page = 1
	}

	// if empty or error get query args per_page, set default to 10
	perPage, err := strconv.Atoi(ctx.QueryParam("per_page"))
	if err != nil {
		perPage = 10
	}

	// call use case or logic get members
	results, err := g.gomUseCase.GetMembers(orgs, page, perPage)

	if err != nil {
		return errorResponse(ctx, err, nil)
	}

	return successResponse(ctx, "", results)
}
