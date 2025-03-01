package main

import (
	"film-app/auth"
	"film-app/config"
	"film-app/film"
	"film-app/user"
	"film-app/validation"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

	userRepository := user.NewDBRepository(db)
	authService := auth.NewJWTService(userRepository)
	authController := auth.NewController(authService, userRepository)
	filmRepository := film.NewDBRepository(db)
	filmPolicy := film.NewPolicy()
	filmController := film.NewController(filmRepository, filmPolicy)
	appContextMiddleware := auth.NewAppContextMiddleware(authService)

	e := echo.New()
	e.Validator = validation.NewStructValidator(validator.New(validator.WithRequiredStructEnabled()))

	e.Use(appContextMiddleware.UseAppContext)
	if config.Env.AppDebug {
		e.Use(middleware.Logger())
	}

	authRoutes := e.Group("/auth")
	authRoutes.POST("/register", authController.Register, auth.VerifyGuest)
	authRoutes.POST("/login", authController.LogIn, auth.VerifyGuest)

	filmRoutes := e.Group("/films", auth.VerifyAuthenticated)
	filmRoutes.GET("", filmController.Index)
	filmRoutes.GET("/:id", filmController.Show)
	filmRoutes.POST("", filmController.Create)
	filmRoutes.PATCH("/:id", filmController.Update)
	filmRoutes.DELETE("/:id", filmController.Delete)

	e.Logger.Fatal(e.Start(config.Env.AppUrl))
}
