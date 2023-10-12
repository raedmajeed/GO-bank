package controllers

import (
	"fmt"
	"net/http"
	"raedmajeed/service/interfaces"
)

type HandlerStruct struct {
	interfaces.Userservice
}

func (s *HandlerStruct) Handlers() {
	fmt.Println("in handler function")
	// http.HandleFunc("/register", s.RegisterUser)
	http.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("here")
		s.Userservice.RegisterUser(w, r)
		fmt.Println("here")
		w.WriteHeader(200)
})
	// http.HandleFunc("/show", us.FindAllUser())
}

func (h *HandlerStruct) Start() {
	fmt.Println("Starting Server")
	http.ListenAndServe(":8010", nil)
}

func NewHandlerImpl(us interfaces.Userservice) HttpControllers {
	return &HandlerStruct{
		us,
	}
}
