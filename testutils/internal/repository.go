package testutils

import (
	"context"
)

type Entity struct {
	ID   string
	Name string
	Time string
}

//go:generate mockery --name Repository --output mocks --with-expecter --filename repository.go --structname Repository
type Repository interface {
	GetByID(ctx context.Context, entityID string, entity interface{}) error
	Save(ctx context.Context, entityID string, entity interface{}) error
}
