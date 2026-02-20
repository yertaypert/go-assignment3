package repository

import (
	"github.com/yertaypert/go-assignment3/internal/repository/_postgres"
	"github.com/yertaypert/go-assignment3/internal/repository/_postgres/users"
	"github.com/yertaypert/go-assignment3/pkg/modules"
)

type UserRepository interface {
	GetUsers() ([]modules.User, error)
}

type Repositories struct {
	UserRepository
}

func NewRepositories(db *_postgres.Dialect) *Repositories {
	return &Repositories{
		UserRepository: users.NewRepository(db),
	}
}
