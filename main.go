package main

import (
	"database/sql"
	"fmt"
	"gin/delivery/controller"
	"gin/repository"
	"gin/usecase"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", "localhost", "5432", "postgres", "stanners2020", "task_management_db")
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic("connection error")
	}

	// Inject DB ke -> repository
	authorRepo := repository.NewAuthorRepository(db)
	taskRepo := repository.NewTaskRepository(db)
	// Inject REPO ke -> useCase
	authUC := usecase.NewAuthorUseCase(authorRepo)
	taskUC := usecase.NewTaskUseCase(taskRepo, authUC)

	r := gin.Default()
	rg := r.Group("/api/v1")

	// Inject UseCase ke -> controller
	controller.NewAuthorController(authUC, rg).Route()
	controller.NewTaskController(taskUC, rg).Route()

	err = r.Run(":8080")
	if err != nil {
		panic("failed to run server")
	}

	// Create Task
	// payload := model.Task{
	// 	Title:    "Golang",
	// 	Content:  "Belajar Golang with GIN",
	// 	AuthorId: "a184e442-1adc-4f61-b87e-a0f16d6cb9cd",
	// }

	// rsv, err := taskUC.RegisterNewTask(payload)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }

	// fmt.Println("created:", rsv)

}

// Branch:
// 1-native-http
// 2-with-gin
// 3-middleware-basic-auth
// 4-middleware-logger
// 5-clean-code-part-1
