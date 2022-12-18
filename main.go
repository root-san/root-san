package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/root-san/root-san/app/handler"
	"github.com/root-san/root-san/app/repository/impl"
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

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:3000", "https://root3.trap.games"},
		AllowHeaders:     []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods:     []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
		AllowCredentials: true,
	}))
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
		getEnvOrDefault("MARIADB_USERNAME", "root"),
		getEnvOrDefault("MARIADB_PASSWORD", "password"),
		getEnvOrDefault("MARIADB_HOSTNAME", "db"),
		getEnvOrDefault("MARIADB_PORT", "3306"),
		getEnvOrDefault("MARIADB_DATABASE", "root_san"),
	)
}

func getEnvOrDefault(key, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}

	return defaultValue
}
