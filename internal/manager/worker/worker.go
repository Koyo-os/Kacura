package worker

import (
	"net/http"
	"sync"

	"github.com/koyo-os/kacura/internal/config"
	"github.com/koyo-os/kacura/pkg/logger"
	"github.com/wneessen/go-mail"
)

type Worker struct{
	client *mail.Client
	logger *logger.Logger
	wg *sync.WaitGroup
	reqChan chan *http.Request
}

func Init(cfg *config.Config) (*Worker,error) {
	logger := logger.Init()

	logger.Info("start init worker")

	client, err := mail.NewClient(
		cfg.Smpt.Host,
		mail.WithPort(cfg.Smpt.Port),
		mail.WithSMTPAuth(mail.SMTPAuthPlain),
		mail.WithUsername(cfg.Smpt.Username),
		mail.WithPassword(cfg.Smpt.Password),
	)
	if err != nil{
		logger.Errorf("error initialize smpt client: %v",err)
		return nil,err
	}

	var wg sync.WaitGroup
	return &Worker{
		client: client,
		logger: logger,
		wg: &wg,
	},nil
}

func (w *Worker) Start() error {

}