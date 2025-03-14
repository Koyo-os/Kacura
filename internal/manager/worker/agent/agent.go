package agent

import (
	"io"
	"net/http"
	"sync"

	"github.com/bytedance/sonic"
	"github.com/koyo-os/kacura/internal/config"
	"github.com/koyo-os/kacura/internal/models"
	"github.com/koyo-os/kacura/pkg/logger"
	"github.com/wneessen/go-mail"
)

type Agent struct{
	Wg *sync.WaitGroup
	Logger *logger.Logger
	Client *mail.Client
	Cfg *config.Config
}

func RunAgent(a *Agent, r *http.Request) {
	var msg models.MailMsg

	if r.Method != http.MethodPost{
		return
	}

	body,err := io.ReadAll(r.Body)
	if err != nil{
		a.Logger.Errorf("cant get body: %v",err)
		a.Wg.Done()
		return
	}

	if err = sonic.Unmarshal(body, &msg);err != nil{
		a.Logger.Errorf("cant unmarshal msg: %v",err)
		a.Wg.Done()
		return
	}

	message := mail.NewMsg()
	if err = message.From(msg.From);err != nil{
		a.Wg.Done()
		return
	}

	if err = message.To(msg.To);err != nil{
		a.Wg.Done()
		return
	}

	message.Subject(msg.Subject)
	message.SetBodyString(mail.TypeTextPlain, msg.Body)

	if err = a.Client.Send(message);err != nil{
		a.Logger.Errorf("cant send email: %v",err)
		a.Wg.Done()
		return
	}

	a.Wg.Done()
}