package interfaces

import "raedmajeed/model"

type Repository interface {
	Create(user model.Entity) error
	FindAll(user model.Entity) error
}