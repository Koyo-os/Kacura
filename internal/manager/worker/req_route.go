package worker

import "github.com/koyo-os/kacura/internal/manager/worker/agent"

func (w *Worker) Route() {
	for {
		r := <- w.reqChan
		agent.RunAgent(
			&agent.Agent{
				Logger: w.logger,
				Wg: w.wg,
				Client: w.client,
			},
			r,
		)
	}
}