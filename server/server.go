package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/rhiadc/infra-provision-orch/config"
	"github.com/rhiadc/infra-provision-orch/domain"
)

type Server struct {
	service *domain.Service
}

func NewServer(service *domain.Service, envs *config.Environments) *Server {

	serverCtx, serverStopCtx := context.WithCancel(context.Background())
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	server := &Server{service: service}
	r := chi.NewRouter()
	server.router(r)
	s := &http.Server{Addr: fmt.Sprintf("0.0.0.0:%s", envs.APIPort), Handler: r}

	go func() {
		<-sig
		shutdownCtx, cancel := context.WithTimeout(serverCtx, 30*time.Second)
		defer cancel()

		go func() {
			<-shutdownCtx.Done()
			if shutdownCtx.Err() == context.DeadlineExceeded {
				log.Fatal("graceful shutdown timed out.. forcing exit.")
			}
		}()

		// Trigger graceful shutdown
		err := s.Shutdown(shutdownCtx)
		if err != nil {
			log.Fatal(err)
		}
		serverStopCtx()

	}()

	go func() {
		log.Printf("Server started on %s", envs.APIPort)
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("Error starting server: %v", err)
		}
	}()

	<-serverCtx.Done()
	log.Println("Server stopped gracefully.")
	return server
}
