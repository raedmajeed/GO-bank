package service

import (
	"encoding/json"
	"fmt"
	"net/http"
	"raedmajeed/model"
	"raedmajeed/repository"
	interfaces "raedmajeed/repository/interfaces"
	interfaceservice "raedmajeed/service/interfaces"
)

type Userserviceimpl struct {
	interfaces.Repository
}

// RegisterUser implements interfaces.Userservice.
func (s *Userserviceimpl) RegisterUser(w http.ResponseWriter, r *http.Request) {
	user := model.Entity{}
	json.NewDecoder(r.Body).Decode(&user)
  fmt.Println("reahced Reginser user =====>>>.")
	s.Create(user)
}

func Test1(w http.ResponseWriter, r *http.Request) {
	fmt.Println("reaffasdfasfas")
}

// FindAllUser implements userservice.
func (*Userserviceimpl) FindAllUser(w http.ResponseWriter, r *http.Request) {
	panic("unimplemented")
}

// RegisterUser implements userservice.
func RegisterUser(w http.ResponseWriter, r *http.Request) {
	s := repository.RepositoryImpl{}
	user := model.Entity{}
	json.NewDecoder(r.Body).Decode(&user)
  fmt.Println("reahced Reginser user =====>>>.")
	s.Create(user)
	// return nil````
}

func NewServiceImpl(repository interfaces.Repository) interfaceservice.Userservice {
	return &Userserviceimpl{
		repository,
	}
}
