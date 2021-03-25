package main

import (
	"fmt"
	"github.com/pkg/errors"
	"log"
	"os"
	"time"

	"github.com/jeroldleslie/go-coding-challenge-trivia/internal/numbers_api"
	"github.com/jeroldleslie/go-coding-challenge-trivia/internal/utils"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

const (
	version           = "v1"
	defaultPort       = "8000"
	dbConnectionRetry = 3
)

func main() {
	if err := func() error {
		log.SetFlags(log.Llongfile | log.LstdFlags)
		dbConnectionStr := os.Getenv("APP_DB_CONNECTION_STR")
		log.Printf("Database Connection String = %s",dbConnectionStr)


		db, err := utils.ConnectDBWithInfiniteRetry(dbConnectionStr, time.Second*5)
		if err != nil {
			return errors.Wrap(err, "Database connection error")
		}
		defer db.Close()

		api := numbers_api.NewAPIHandler(&numbers_api.API{
			DBMapper: &numbers_api.DBMapper{
				DB: db,
			},
		})
		/*api := &numbers_api.APIHandler{
			API: &numbers_api.API{
				DBMapper: &numbers_api.DBMapper{
					DB: db,
				},
			},
		}*/

		serve(api)
		return nil
	}(); err != nil {
		log.Printf("%+v", err)
		os.Exit(1)
	}
}

func serve(a *numbers_api.APIHandler) {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())

	g := e.Group("api")
	g.GET(fmt.Sprintf("/%s/ping", version), a.Ping)
	g.GET(fmt.Sprintf("/%s/trivia", version), a.GetDetails)

	e.Logger.Fatal(e.Start(":" + utils.ParseApiPortFromEnv(defaultPort)))
}


