package interfaces

import "raedmajeed/model"

type Repository interface {
	CreateUser(user model.Entity) error
	FindAll(user model.Entity) error
}