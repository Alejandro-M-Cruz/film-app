package main

import (
    "film-app/auth"
    "film-app/config"
    "film-app/film"
    "film-app/middleware"
    "github.com/joho/godotenv"
    "github.com/labstack/echo/v4"
    "gorm.io/driver/sqlite"
    "gorm.io/gorm"
)

func main() {
    err := godotenv.Load()
    if err != nil {
        panic(err)
    }

    conf := config.InitConfig()

    db, err := gorm.Open(sqlite.Open(conf.DBPath), new(gorm.Config))
    if err != nil {
        panic(err)
    }

    e := echo.New()

    filmService := film.NewDBService(db)
    filmController := film.NewController(filmService)

    filmRouteGroup := e.Group("/films", auth.AuthMiddleware)
    filmRouteGroup.GET("", filmController.Index)
    filmRouteGroup.GET("/:id", filmController.Show)
    filmRouteGroup.POST("", filmController.Create)
    filmRouteGroup.PUT("/:id", filmController.Update)
    filmRouteGroup.PATCH("/:id", filmController.Update)
    filmRouteGroup.DELETE("/:id", filmController.Delete)

    e.Logger.Fatal(e.Start(conf.AppUrl))
}
