package factory

import (
	userController "learn-echo/features/users/controller"
	userRepository "learn-echo/features/users/repository"
	userService "learn-echo/features/users/service"

	"gorm.io/gorm"
)

type Presenter struct {
	UserPresenter userController.UserController
}

func InitFactory(db *gorm.DB) Presenter {
	UserRepository := userRepository.NewUserRepository()
	UserService := userService.NewUserService(UserRepository, db)
	UserController := userController.NewUserController(UserService)

	return Presenter{
		UserPresenter: UserController,
	}
}
