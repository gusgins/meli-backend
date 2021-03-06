package storage

import "github.com/gusgins/meli-backend/model"

type Storage interface {
	Find(model.Registry) (bool, error)
	Store(model.Registry) error
}