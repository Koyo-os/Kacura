package server

import (
	"net/http"

	"github.com/koyo-os/kacura/internal/config"
	"github.com/koyo-os/kacura/pkg/logger"
)

type Server struct {
	server *http.Server
	logger *logger.Logger
}

func Init(cfg *config.Config) *Server {

}
