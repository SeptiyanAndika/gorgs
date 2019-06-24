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
	e.Use(middleware.RequestID(), middlewareHttp.RoutingLog)
	e.Any("", func(context echo.Context) error {
		return context.JSON(http.StatusOK, "member api")
	})
	e.Any("health", func(context echo.Context) error {

		return context.JSON(http.StatusOK, map[string]interface{}{
			"time": time.Now().Unix(),
		})
	})

	ga:=adapter.NewGithubAdapter()
	gomUseCase := usecase.NewGithubAdapter(ga)
	gomUseCase = logging.NewGithubOrgMembersLog(gomUseCase)

	deliveryHttp.NewGithubOrgMembers(e, gomUseCase)
	e.Logger.Fatal(e.Start(":8082"))
}
