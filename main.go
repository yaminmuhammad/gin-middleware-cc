package main

import (
	"database/sql"
	"fmt"
	"gin/model"
	"gin/repository"
	"gin/usecase"

	_ "github.com/lib/pq"
)

// type Task struct {
// 	ID      string `json:"id"`
// 	Title   string `json:"title"`
// 	Content string `json:"content"`
// }

// var tasks []Task

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

	// Create Task
	payload := model.Task{
		Title:    "CDR",
		Content:  "",
		AuthorId: "4a223243-72e4-4bc1-af27-ee2b94d84142",
	}

	rsv, err := taskUC.RegisterNewTask(payload)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("created:", rsv)

	// Inject DB ke dalam repository
	// taskRepo := repository.NewTaskRepository(db)
	// authorRepo := repository.NewAuthorRepository(db)
	// Inject REPO ke dalam usecase
	// taskUC := usecase.NewTaskUseCase(taskRepo)
	// authorUC := usecase.NewAuthorUseCase(authorRepo)
	// Buat router dan inject UC ke dalam controller

	// authors, err := authorUC.FindAllAuthor()
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }
	// for _, author := range authors {
	// 	fmt.Println("List Author:")
	// 	fmt.Println("ID:", author.ID)
	// 	fmt.Println("Name:", author.Name)
	// 	fmt.Println("Email:", author.Email)
	// 	fmt.Println("CreatedAt:", author.CreatedAt)
	// 	fmt.Println()
	// }

	// tasks, err := taskUC.FindAllTask()
	// if err != nil {
	// 	fmt.Println(err.Error())
	// 	return
	// }
	// for _, task := range tasks {
	// 	fmt.Println("List Task:")
	// 	fmt.Println("ID:", task.ID)
	// 	fmt.Println("Title:", task.Title)
	// 	fmt.Println("Content:", task.Content)
	// 	fmt.Println("CreatedAt:", task.CreatedAt)
	// 	fmt.Println()
	// }

	// r := gin.New()
	// r.Use(LogMiddleware())
	// rg := r.Group("/api/v1")
	// rg.GET("/tasks", getTasks)
	// rg.POST("/tasks", BasicAuthMiddleware(), addTask)
	// rg.PUT("/tasks/:id", BasicAuthMiddleware(), updateTask)
	// rg.DELETE("/tasks/:id", BasicAuthMiddleware(), deleteTask)

	// r.Run(":8080")
}

// func getTasks(c *gin.Context) {
// 	if len(tasks) == 0 {
// 		response := gin.H{
// 			"href":    "/tasks",
// 			"message": "data not found",
// 			"status":  200,
// 		}
// 		c.JSON(http.StatusOK, response)
// 		return
// 	}

// 	response := struct {
// 		Message string `json:"message"`
// 		Data    []Task `json:"data"`
// 		Status  string `json:"status"`
// 	}{
// 		Message: "Success",
// 		Data:    tasks,
// 		Status:  "200",
// 	}

// 	c.JSON(http.StatusOK, response)
// }

// func addTask(c *gin.Context) {
// 	var task Task
// 	if err := c.ShouldBind(&task); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	tasks = append(tasks, task)
// 	c.JSON(http.StatusCreated, task)
// }

// func updateTask(c *gin.Context) {
// 	taskID := c.Param("id")

// 	// Cari task berdasarkan ID
// 	var foundTask *Task
// 	for i := range tasks {
// 		if tasks[i].ID == taskID {
// 			foundTask = &tasks[i]
// 			break
// 		}
// 	}

// 	// Jika task tidak ditemukan
// 	if foundTask == nil {
// 		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
// 		return
// 	}

// 	// Bind data yang baru dari request
// 	var updatedTask Task
// 	if err := c.ShouldBindJSON(&updatedTask); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	// Perbarui task
// 	foundTask.Title = updatedTask.Title
// 	foundTask.Content = updatedTask.Content

// 	c.JSON(http.StatusOK, foundTask)
// }

// func deleteTask(c *gin.Context) {
// 	taskID := c.Param("id")
// 	for i, task := range tasks {
// 		if task.ID == taskID {
// 			tasks = append(tasks[:i], tasks[i+1:]...)
// 			c.JSON(http.StatusCreated, nil)
// 			return
// 		}
// 	}
// 	c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
// }

// func BasicAuth(c *gin.Context) {
// 	authHeader := c.Request.Header.Get("Authorization")

// 	if !strings.Contains(authHeader, "Basic") {
// 		result := gin.H{
// 			"status":  http.StatusForbidden,
// 			"message": "invalid token",
// 			"href":    c.Request.RequestURI,
// 		}
// 		c.JSON(http.StatusForbidden, result)
// 		c.Abort()
// 		return
// 	}

// 	clientSecret := "1jutadolar2024"
// 	clientID := "bikin.dev"
// 	tokenString := strings.Replace(authHeader, "Basic ", "", -1)
// 	myToken := clientID + ":" + clientSecret
// 	myBasicAuth := base64.StdEncoding.EncodeToString([]byte(myToken))
// 	if tokenString != myBasicAuth {
// 		result := gin.H{
// 			"status":  http.StatusUnauthorized,
// 			"message": "invalid authentication",
// 			"href":    c.Request.RequestURI,
// 		}
// 		c.JSON(http.StatusUnauthorized, result)
// 		c.Abort()
// 		return
// 	}
// }

// func BasicAuthMiddleware() gin.HandlerFunc {
// 	return func(context *gin.Context) {
// 		username, password, ok := context.Request.BasicAuth()

// 		// gabungkan
// 		if !ok || username != "bikin.dev" || password != "1jutadolar2024" {
// 			context.AbortWithStatus(http.StatusUnauthorized)
// 			return
// 		}
// 		// supaya bisa lanjut ke request lainnya
// 		context.Next()
// 	}
// }

// Branch:
// 1-native-http
// 2-with-gin
// 3-middleware-basic-auth
// 4-middleware-logger
// 5-clean-code-part-1

// func LogMiddleware() gin.HandlerFunc {
// 	return func(context *gin.Context) {
// 		context.Next()
// 		// Log waktu access
// 		// Method, status code, client, IP , waktu akses (latency)
// 		latency := time.Since(time.Now())
// 		method := context.Request.Method
// 		statusCode := context.Writer.Status()
// 		ipAddress := context.ClientIP()
// 		client := context.Request.UserAgent()
// 		path := context.Request.URL.Path
// 		log.Println(latency, method, statusCode, ipAddress, client, path)
// 	}
// }
