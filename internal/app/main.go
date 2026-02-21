package app

import (
	"context"
	"net/http"
	"time"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/yertaypert/go-assignment3/internal/handler"
	"github.com/yertaypert/go-assignment3/internal/repository"
	"github.com/yertaypert/go-assignment3/internal/repository/_postgres"
	"github.com/yertaypert/go-assignment3/internal/usecase"
	"github.com/yertaypert/go-assignment3/pkg/modules"
)

func Run() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Config & DB connection
	dbConfig := initPostgreConfig()
	_postgre := _postgres.NewPGXDialect(ctx, dbConfig)

	// Repository Layer
	repositories := repository.NewRepositories(_postgre)

	// Usecase Layer
	userUsecase := usecase.NewUserUsecase(repositories.UserRepository)

	// Handler Layer
	userHandler := handler.NewUserHandler(userUsecase)

	// Setup routes
	http.HandleFunc("/users", userHandler.GetUsers)
	http.HandleFunc("/users/get", userHandler.GetUserByID) // /users/get?id=1
	http.HandleFunc("/users/create", userHandler.CreateUser)
	http.HandleFunc("/users/update", userHandler.UpdateUser)
	http.HandleFunc("/users/delete", userHandler.DeleteUser)

	// Start server
	println("Server starting on :8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}

	//// Get all Users
	//users, err := repositories.GetUsers()
	//if err != nil {
	//	fmt.Printf("Error fetching users: %v\n", err)
	//	return
	//}
	//fmt.Printf("Users: %v\n", users)

	//// Create a New User
	//newUser := &modules.User{Name: "Jane Smith",
	//	Age:   24,
	//	Email: "jane@example.com"}
	//id, err := repositories.CreateUser(newUser)
	//if err != nil {
	//	fmt.Printf("%v\n", err)
	//}
	//fmt.Printf("Created user: %v\n", id)
	//
	//// Create another New User
	//newUser = &modules.User{Name: "John Doe",
	//	Age:   26,
	//	Email: "joe@example.com"}
	//id, err = repositories.CreateUser(newUser)
	//if err != nil {
	//	fmt.Printf("%v\n", err)
	//}
	//fmt.Printf("Created user: %v\n", id)
	//
	//// Get all Users
	//users, err = repositories.GetUsers()
	//if err != nil {
	//	fmt.Printf("Error fetching users: %v\n", err)
	//	return
	//}
	//fmt.Printf("Users: %v\n", users)
	//
	//// Update User
	//ExistingUserToUpdate := &modules.User{
	//	ID:    2,
	//	Name:  "Jane Updated",
	//	Age:   30,
	//	Email: "jane_new@example.com",
	//}
	//err = repositories.UpdateUser(ExistingUserToUpdate)
	//if err != nil {
	//	fmt.Printf("Error updating user: %v\n", err)
	//} else {
	//	fmt.Printf("Updated user: %v\n", ExistingUserToUpdate)
	//}
	//
	//// Delete User
	//rowsAffected, err := repositories.DeleteUser(2)
	//if err != nil {
	//	fmt.Printf("Error deleting users: %v\n", err)
	//}
	//fmt.Printf("Rows affected: %d\n", rowsAffected)
	//
	//// Get all Users
	//users, err = repositories.GetUsers()
	//if err != nil {
	//	fmt.Printf("Error fetching users: %v\n", err)
	//	return
	//}
	//fmt.Printf("Users: %v\n", users)
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
