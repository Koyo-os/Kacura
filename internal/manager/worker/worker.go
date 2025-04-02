package worker

import (
	"net/http"
	"sync"

	"github.com/koyo-os/kacura/internal/config"
	"github.com/koyo-os/kacura/internal/manager/worker/agent"
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
	stopChan chan struct{}
	mux *sync.Mutex
	counter uint
	cfg *config.Config
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

	var (
		chreq chan *http.Request
	 	counter chan struct{}
	 	stopChan chan struct{}
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
		stopChan: stopChan,
		mux: &mux,
		cfg: cfg,
	},nil
}

func (w *Worker) GetCount() uint {
	return w.counter
}

func (w *Worker) Run() error {
	go w.CounterInc()
	
	for {		
		w.counterInc <- struct{}{}
		
		req := <- w.ReqChan
		agent.RunAgent(
			&agent.Agent{
				Wg: w.wg,
				Logger: w.logger,
				Client: w.client,
				Cfg: w.cfg,
			}
		)

		w.counterDec <- struct{}{}
	}
}

func (w *Worker) CounterInc() {
	for {
		w.mux.Lock()

		select{
		case <- w.counterInc:
			w.counter++
			w.mux.Unlock()
		case <- w.counterDec:
			w.counter--
			w.mux.Unlock()
		}
	}
}

