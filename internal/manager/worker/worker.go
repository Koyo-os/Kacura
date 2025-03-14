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
	ReqChan chan *http.Request
	counterInc chan struct{}
	counterDec chan struct{}
	mux *sync.Mutex
}

var COUNTER uint

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

	var (
		chreq chan *http.Request
	 	counter chan struct{}
		counterDec chan struct{}
		mux sync.Mutex
	)

	var wg sync.WaitGroup
	return &Worker{
		client: client,
		logger: logger,
		wg: &wg,
		ReqChan: chreq,
		counterInc: counter,
		counterDec: counterDec,
		mux: &mux,
	},nil
}

func (w *Worker) CounterInc() {
	for {
		w.mux.Lock()

		select{
		case <- w.counterInc:
			COUNTER++
			w.mux.Unlock()
		case <- w.counterDec:
			COUNTER--
			w.mux.Unlock()
		}
	}
}

