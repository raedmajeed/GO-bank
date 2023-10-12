package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"raedmajeed/model"
	interfaces "raedmajeed/repository/interfaces"
	interfaceservice "raedmajeed/service/interfaces"
)

type Userserviceimpl struct {
	interfaces.Repository
}

// RegisterUser implements interfaces.Userservice.
func (s *Userserviceimpl) RegisterUser(w http.ResponseWriter, r *http.Request) {
	user := &model.Entity{}
	json.NewDecoder(r.Body).Decode(user)
  fmt.Println("reahced Reginser user =====>>>.")
	fmt.Println(*user)
	s.CreateUser(*user)
}

// FindAllUser implements userservice.
func (s *Userserviceimpl) FindAllUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("raed")
}


func NewServiceImpl(repository interfaces.Repository) interfaceservice.Userservice {
	return &Userserviceimpl{
		repository,
	}
}
