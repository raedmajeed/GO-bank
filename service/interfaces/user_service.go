package interfaces

import (
	"net/http"
)

type Userservice interface {
	RegisterUser(w http.ResponseWriter, r *http.Request)
	FindAllUser(w http.ResponseWriter, r *http.Request)
}