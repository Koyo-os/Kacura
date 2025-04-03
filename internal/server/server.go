package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/koyo-os/kacura/internal/config"
	"github.com/koyo-os/kacura/pkg/logger"
)

type Server struct {
	server *http.Server
	logger *logger.Logger
}

func Init(cfg *config.Config) *Server {
	return &Server{
		server: &http.Server{
			Addr: fmt.Sprintf("%s:%d", cfg.Host, cfg.Portv),
		}
		logger: logger.Init(),
	}
}

func (s *Server) SetHandler(mux *http.ServeMux) {
	s.server.Handler = mux
}

func (s *Server) Run() error {
	s.logger.Infof("starting kacura server on: %s", s.server.Addr)

	return s.server.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) {
	s.logger.Info("stopping kacura server...Bye!")

	s.server.Shutdown(ctx)
}
