package agent

import (
	"io"
	"net/http"
	"sync"

	"github.com/koyo-os/kacura/internal/models"
	"github.com/koyo-os/kacura/pkg/logger"
)

func RunAgent(logger *logger.Logger, wg *sync.WaitGroup, r *http.Request) {
	var msg models.MailMsg

	body,err := io.ReadAll(r.Body)
	if err != nil{
		
	}
}