package repository

import (
	"errors"

	"github.com/gusgins/meli-backend/model"
)

type Repository interface {
	FindMutant(model.Registry) (bool, error)
	StoreRegistry(model.Registry) error
	GetStats() (model.Stats, error)
	Close() error
}

// ErrRegistryNotFound is returned by Find when the registry was not
// found in repository
var ErrRegistryNotFound = errors.New("repository: registry not found")

// ErrStatsNotFound is returned by GetStats when stats were not found
// in repository
var ErrStatsNotFound = errors.New("repository: stats not found")
