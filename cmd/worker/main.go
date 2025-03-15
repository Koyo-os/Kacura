package main

import (
	"net/http"

	"github.com/koyo-os/kacura/internal/config"
	"github.com/koyo-os/kacura/internal/handler"
	"github.com/koyo-os/kacura/pkg/logger"
)

func main() {
	logger := logger.Init()

	cfg,err := config.Init()
	if err != nil{
		logger.Error(err)
		return
	}

	handler,err := handler.Init(cfg)
	if err != nil{
		logger.Error(err)
		return
	}

	http.HandleFunc("/info", handler.Counter)
	http.HandleFunc("/", handler.MainHandler)

	http.ListenAndServe(":8080", nil)
}