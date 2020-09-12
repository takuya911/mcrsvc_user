package api

import (
	"log"

	"github.com/labstack/echo"
	_userHandler "github.com/takuya911/mcrsvc_user/handler"
	_infra "github.com/takuya911/mcrsvc_user/infra"
	_userRepo "github.com/takuya911/mcrsvc_user/repository"
	_userUsecase "github.com/takuya911/mcrsvc_user/usecase"
)

// Run function
func Run() {
	// mysql connect
	db, err := _infra.NewGormDB()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()

	e := echo.New()
	// user services
	userRepo := _userRepo.NewUserRepository(db)
	userUsecase := _userUsecase.NewUserUsecase(userRepo)
	_userHandler.NewUserHandler(e, userUsecase)

	log.Fatal(e.Start(":8080"))
}
