package delivery

import (
	"database/sql"
	"fmt"
	"gin/config"
	"gin/delivery/controller"
	"gin/repository"
	"gin/usecase"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

type Server struct {
	authorUc usecase.AuthorUseCase
	taskUc   usecase.TaskUseCase
	engine   *gin.Engine
	host     string
}

func (s *Server) initRoute() {
	rg := s.engine.Group(config.ApiGroup)
	controller.NewAuthorController(s.authorUc, rg).Route()
	controller.NewTaskController(s.taskUc, rg).Route()
}

func (s *Server) Run() {
	s.initRoute()
	if err := s.engine.Run(s.host); err != nil {
		panic(fmt.Errorf("server not running on host %s, becauce error %v", s.host, err.Error()))
	}
}

func NewServer() *Server {
	cfg, _ := config.NewConfig()
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.Name)
	db, err := sql.Open(cfg.Driver, dsn)
	if err != nil {
		panic("connection error")
	}
	// Inject DB ke -> repository
	authorRepo := repository.NewAuthorRepository(db)
	taskRepo := repository.NewTaskRepository(db)
	// Inject REPO ke -> useCase
	authUC := usecase.NewAuthorUseCase(authorRepo)
	taskUC := usecase.NewTaskUseCase(taskRepo, authUC)
	engine := gin.Default()
	host := fmt.Sprintf(":%s", cfg.ApiPort)
	return &Server{
		authorUc: authUC,
		taskUc:   taskUC,
		engine:   engine,
		host:     host,
	}
}
