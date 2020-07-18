package http

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

//Server iniciar aplicação Go no localhost
type Server struct {
	server *http.Server
}

//New custumiza ambiente da aplicação Go
func New(port string, handler http.Handler) *Server {
	return &Server{
		server: &http.Server{
			Addr:         ":" + port,
			Handler:      handler,
			ReadTimeout:  5 * time.Second,
			WriteTimeout: 55 * time.Second,
		},
	}
}

//ListenAndServe apresenta informações se foi possível estabelecer a aplicação Go em localhost.
func (s *Server) ListenAndServe() {
	go func() {
		fmt.Printf("Server started at %s!\n", s.server.Addr)
		if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			fmt.Printf("Error: %s\n", err)
		}
	}()
}

//Shutdown derruba a aplicação
func (s *Server) Shutdown() {
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()
	s.server.Shutdown(ctx)
	//fmt.Println("server down!")
}
