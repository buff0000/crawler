package scheduler

import (
	"com.buff/Crawler/crawler/crawler/engine"
)

type QueuedScheduler struct {
	requestChan chan engine.Request
	workerChan  chan chan engine.Request
}

func (s QueuedScheduler) Submit(r engine.Request) {
	s.requestChan <- r
}

func (s *QueuedScheduler) WorkerChan() chan engine.Request {
	return make (chan engine.Request)
}

func (s *QueuedScheduler) WorkerReady(w chan engine.Request) {
	s.workerChan <- w
}

func (s *QueuedScheduler) Run() {
	s.workerChan = make(chan chan engine.Request)
	s.requestChan = make(chan engine.Request)
	go func() {
		var reqestQ []engine.Request
		var workerQ []chan engine.Request
		for {
			var activeRequest engine.Request
			var activeWorker chan engine.Request
			if len(reqestQ) > 0 && len(workerQ) > 0 {
				activeRequest = reqestQ[0]
				activeWorker = workerQ[0]
			}
			select {
			case r := <-s.requestChan:
				reqestQ = append(reqestQ, r)
			case w := <-s.workerChan:
				workerQ = append(workerQ, w)
				case activeWorker <- activeRequest:
					workerQ = workerQ[1:]
					reqestQ = reqestQ[1:]

			}
		}
	}()
}
