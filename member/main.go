package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gorgs/member/adapter"
	deliveryHttp "gorgs/member/delivery/http"
	"gorgs/member/logging"
	middlewareHttp "gorgs/member/middleware/http"
	"gorgs/member/usecase"
	"net/http"
	"time"
)

func main() {
	e := echo.New()
	// add general middleware
	e.Use(middleware.RequestID(), middlewareHttp.RoutingLog)

	// root endpoint
	e.Any("", func(context echo.Context) error {
		return context.JSON(http.StatusOK, "member services")
	})

	// health endpoint give response time server
	e.Any("health", func(context echo.Context) error {

		return context.JSON(http.StatusOK, map[string]interface{}{
			"time": time.Now().Unix(),
		})
	})

	// initialise adapter
	ga:=adapter.NewGithubAdapter()

	// initialise use case
	gomUseCase := usecase.NewGithubOrgMemberUsecase(ga)

	// wrap use case with logging
	gomUseCase = logging.NewGithubOrgMembersLog(gomUseCase)

	//attach use case or logic to delivery
	deliveryHttp.NewGithubOrgMembers(e, gomUseCase)

	e.Logger.Fatal(e.Start(":8082"))
}
