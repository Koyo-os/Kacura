package handler

import (
	"fmt"
	"net/http"

	"github.com/koyo-os/kacura/internal/config"
	"github.com/koyo-os/kacura/internal/manager/worker"
	"github.com/koyo-os/kacura/pkg/logger"
)

type WorkerHandler struct{
	logger *logger.Logger
	worker *worker.Worker
}

func (work *WorkerHandler) MainHandler(w http.ResponseWriter, r *http.Request) {
	work.worker.ReqChan <- r
}

func (work *WorkerHandler) Counter(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprint(w, worker.COUNTER)
	if err != nil{
		work.logger.Errorf("cant print counter: %v", err)
		return
	}
}

func Init(cfg *config.Config) (*WorkerHandler,error) {
	logger := logger.Init()

	worker, err := worker.Init(cfg)
	if err != nil{
		logger.Errorf("cant get worker: %v",err)
		return nil,err
	}

	return &WorkerHandler{
		worker: worker,
		logger: logger,
	},nil
}