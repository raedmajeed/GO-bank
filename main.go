package main

import (
	"raedmajeed/di"
)

func main() {
	
	h := di.InitApi()
	// http.ListenAndServe(":8010", nil)
	// https := controllers.HandlerStruct{}
	h.Handlers()

	h.Start()

	
}