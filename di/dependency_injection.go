package di

import (
	handlers "raedmajeed/controllers"
	"raedmajeed/db"
	"raedmajeed/repository"
	"raedmajeed/service"
)

func InitApi() handlers.HttpControllers {
	
	DB := db.ConnnectToDB()
	repository := repository.NewRepositoryImpl(DB)
	service := service.NewServiceImpl(repository)
	controller := handlers.NewHandlerImpl(service)

	return controller

}