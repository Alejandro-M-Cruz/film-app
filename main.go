package main

import (
    "film-app/auth"
    "film-app/config"
    "film-app/film"
    "film-app/user"
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

    config.InitConfig()

    db, err := gorm.Open(sqlite.Open(config.Env.DBPath), &gorm.Config{TranslateError: true})
    if err != nil {
        panic(err)
    }

    userService := user.NewDBService(db)
    authService := auth.NewJWTService(userService)
    authController := auth.NewController(authService, userService)

    e := echo.New()

    e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
        return auth.UseCustomContext(next, authService)
    })

    authRouteGroup := e.Group("/auth")
    authRouteGroup.POST("/register", authController.Register, auth.VerifyGuest)
    authRouteGroup.POST("/login", authController.LogIn, auth.VerifyGuest)

    filmService := film.NewDBService(db)
    filmController := film.NewController(filmService)

    filmRouteGroup := e.Group("/films", auth.VerifyAuthenticated)
    filmRouteGroup.GET("", filmController.Index)
    filmRouteGroup.GET("/:id", filmController.Show)
    filmRouteGroup.POST("", filmController.Create)
    filmRouteGroup.PUT("/:id", filmController.Update)
    filmRouteGroup.PATCH("/:id", filmController.Update)
    filmRouteGroup.DELETE("/:id", filmController.Delete)

    e.Logger.Fatal(e.Start(config.Env.AppUrl))
}
