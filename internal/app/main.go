package app

import (
	"context"
	"fmt"
	"time"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/yertaypert/go-assignment3/internal/repository"
	"github.com/yertaypert/go-assignment3/internal/repository/_postgres"
	"github.com/yertaypert/go-assignment3/pkg/modules"
)

func Run() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	dbConfig := initPostgreConfig()

	_postgre := _postgres.NewPGXDialect(ctx, dbConfig)

	repositories := repository.NewRepositories(_postgre)

	users, err := repositories.GetUsers()
	if err != nil {
		fmt.Printf("Error fetching users: %v\n", err)
		return
	}
	fmt.Printf("Users: %v\n", users)
}

func initPostgreConfig() *modules.PostgreConfig {
	return &modules.PostgreConfig{
		Host:        "localhost",
		Port:        "5432",
		Username:    "postgres",
		Password:    "starPsql1221!",
		DBName:      "mydb",
		SSLMode:     "disable",
		ExecTimeout: 5 * time.Second,
	}
}
