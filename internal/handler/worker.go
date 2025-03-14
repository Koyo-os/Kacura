package handler

import (
	"net/http"

	"github.com/koyo-os/kacura/internal/manager/worker"
	"github.com/koyo-os/kacura/pkg/logger"
)

type WorkerHandler struct{
	worker *worker.Worker
	logger *logger.Logger
}

func (worker *WorkerHandler) MainHandler(w http.ResponseWriter, r *http.Request) {
	
}