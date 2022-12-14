package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/root-san/root-san/app/handler"
	"github.com/root-san/root-san/app/repository/impl"
	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/root-san/root-san/gen/api"
)

func main() {
	db, err := sqlx.Open("mysql", getDSN())
	if err != nil {
		log.Fatal(err)
	}

	for i := 0; i < 10; i++ {
		if err := db.DB.Ping(); err == nil {
			break
		} else if i == 9 {
			log.Fatal(err)
		}

		time.Sleep(time.Second * time.Duration(i+1))
	}

	e := echo.New()

	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	r := impl.NewRepository(db)

	server := handler.Server{
		Repo: r,
	}

	api.RegisterHandlers(e, &server)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%v", 8080)))
}

func getDSN() string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?parseTime=True&loc=Asia%%2FTokyo&charset=utf8mb4",
		getEnvOrDefault("DB_USER", "root"),
		getEnvOrDefault("DB_PASS", "password"),
		getEnvOrDefault("DB_HOST", "db"),
		getEnvOrDefault("DB_PORT", "3306"),
		getEnvOrDefault("DB_NAME", "root_san"),
	)
}

func getEnvOrDefault(key, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return defaultValue
}