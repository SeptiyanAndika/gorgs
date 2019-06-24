package main

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	dbSql "gorgs/comment/database/sql"
	deliveryHttp "gorgs/comment/delivery/http"
	"gorgs/comment/logging"
	middlewareHttp "gorgs/comment/middleware/http"
	"gorgs/comment/model"
	repoSql "gorgs/comment/repository/sql"
	"gorgs/comment/usecase"
	"net/http"
	"os"
	"strconv"
	"time"
)

func main() {

	dbHost := os.Getenv("DB_HOST")
	dbPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		dbPort = 5423
	}
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	dbSql.Init(dbHost, dbPort, dbUser, dbPass, dbName,true)
	defer dbSql.GetInstance().Close()

	autoMigrate(dbSql.GetInstance())

	e := echo.New()
	e.Use(middleware.RequestID(), middlewareHttp.RoutingLog)
	e.Any("", func(context echo.Context) error {
		return context.JSON(http.StatusOK, "comment api")
	})
	e.Any("health", func(context echo.Context) error {
		dbConnect := true
		dbMessage := "ok"
		errDB := dbSql.GetInstance().DB().Ping()
		if errDB != nil {
			dbConnect = false
			dbMessage = errDB.Error()
		}
		return context.JSON(http.StatusOK, map[string]interface{}{
			"time": time.Now().Unix(),
			"db": map[string]interface{}{
				"connect": dbConnect,
				"message": dbMessage,
			},
		})
	})

	commentRepo := repoSql.NewCommentRepo(dbSql.GetInstance())
	commentUseCase := usecase.NewCommentUseCase(commentRepo)
	commentUseCase = logging.NewCommentLog(commentUseCase)

	deliveryHttp.NewComment(e, commentUseCase)

	e.Logger.Fatal(e.Start(":8081"))

}

func autoMigrate(db *gorm.DB) {
	db.AutoMigrate(model.Comment{})
}
