// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package config

import (
	"go-checkin/controllers"
	"go-checkin/repository"
	"go-checkin/service"
	"gorm.io/gorm"
)

// Injectors from di.go:

func InjectUserController(db *gorm.DB) controllers.UserController {
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userController := controllers.NewUserController(userService)
	return userController
}

func InjectAuthController(db *gorm.DB) controllers.AuthController {
	userRepository := repository.NewUserRepository(db)
	authService := service.NewAuthService(userRepository)
	authController := controllers.NewAuthController(authService)
	return authController
}

func InjectMenuController(db *gorm.DB) controllers.MenuController {
	menuRepository := repository.NewMenuRepository(db)
	menuService := service.NewMenuService(menuRepository)
	menuController := controllers.NewMenuController(menuService)
	return menuController
}

func InjectConfigController(db *gorm.DB) controllers.ConfigController {
	repository := repository.NewConfigRepository(db)
	service := service.NewConfigService(repository)
	controller := controllers.NewConfigController(service)
	return controller
}

func InjectAttendanceController(db *gorm.DB) controllers.AttendanceController {
	repository := repository.NewAttendanceRepository(db)
	service := service.NewAttendanceService(repository)
	controller := controllers.NewAttendanceController(service)
	return controller
}

