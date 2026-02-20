package users

import (
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/yertaypert/go-assignment3/internal/repository/_postgres"
	"github.com/yertaypert/go-assignment3/pkg/modules"
)

type Repository struct {
	db              *sqlx.DB
	executionTimout time.Duration
}

func NewRepository(db *_postgres.Dialect) *Repository {
	return &Repository{
		db:              db.DB,
		executionTimout: 5 * time.Second,
	}
}

func (r *Repository) GetUsers() ([]modules.User, error) {
	var users []modules.User
	err := r.db.Select(&users, "SELECT id, name FROM users")
	if err != nil {
		return nil, err
	}
	fmt.Println(users)
	return users, nil
}
