package api

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgxpool"
	"gitub.com/matheus-hrm/curiously/service/user"
)

type APIServer struct {
	addr string
	db  *pgxpool.Pool
	router *gin.Engine
}

func New(addr string, db *pgxpool.Pool) *APIServer {
	return &APIServer{
		addr: addr, 
		db: db,
		router: gin.Default(),
	}
}

func (s *APIServer) Router() *gin.Engine {
	return s.router
} 

func (s *APIServer) SetupRoutes() {
	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(s.router)
}

func (s *APIServer) Run() error {
	s.SetupRoutes()

	srv := &http.Server{
		Addr: s.addr,
		Handler: s.router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutting down server...")
	
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %s\n", err)
	}

	log.Fatalf("Server exiting")
	return nil
}