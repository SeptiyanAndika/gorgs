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

	// get config from env
	dbHost := os.Getenv("DB_HOST")
	dbPort, err := strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		dbPort = 5423
	}
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")

	// init mysql
	dbSql.Init(dbHost, dbPort, dbUser, dbPass, dbName,true)
	defer dbSql.GetInstance().Close()

	// auto migrate schema table
	autoMigrate(dbSql.GetInstance())

	e := echo.New()

	// add general middleware
	e.Use(middleware.RequestID(), middlewareHttp.RoutingLog)

	// root endpoint
	e.Any("", func(context echo.Context) error {
		return context.JSON(http.StatusOK, "comment services")
	})

	// health endpoint give response time server and status db
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

	// initialise repo
	commentRepo := repoSql.NewCommentRepo(dbSql.GetInstance())

	// initialise use case
	commentUseCase := usecase.NewCommentUseCase(commentRepo)

	// wrap use case with logging
	commentUseCase = logging.NewCommentLog(commentUseCase)

	//attach use case or logic to delivery
	deliveryHttp.NewComment(e, commentUseCase)

	e.Logger.Fatal(e.Start(":8081"))

}

func autoMigrate(db *gorm.DB) {
	db.AutoMigrate(model.Comment{})
}
