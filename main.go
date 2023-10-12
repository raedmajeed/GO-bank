package main

import (
	"raedmajeed/controllers"
	"raedmajeed/di"
)

func main() {
	
	di.InitApi()
	// http.ListenAndServe(":8010", nil)
	https := controllers.HandlerStruct{}
	https.Handlers()

	https.Start()

	
}