package factory

import (
	accountH "learn-echo/features/accounts/handler"
	accountR "learn-echo/features/accounts/repository"
	accountS "learn-echo/features/accounts/service"
	userC "learn-echo/features/users/controller"
	userR "learn-echo/features/users/repository"
	userS "learn-echo/features/users/service"

	"gorm.io/gorm"
)

type Presenter struct {
	UserPresenter    userC.UserController
	AccountPresenter accountH.AccountHandler
}

func InitFactory(db *gorm.DB) Presenter {
	UserRepository := userR.NewUserRepository()
	UserService := userS.NewUserService(UserRepository, db)
	UserController := userC.NewUserController(UserService)

	AccountRepository := accountR.NewAccountRepository()
	AccountService := accountS.NewAccountService(AccountRepository, db)
	AccountHandler := accountH.NewAccountHandler(AccountService)

	return Presenter{
		UserPresenter:    UserController,
		AccountPresenter: *AccountHandler,
	}
}
