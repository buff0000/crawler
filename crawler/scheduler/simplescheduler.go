package scheduler

import "com.buff/Crawler/crawler/crawler/engine"

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) Submit ( request engine.Request) {
	go func() {
		s.workerChan <- request
	}()
	//s.workerChan <- request
}

func (s *SimpleScheduler) ConfigureMasterWorkerChan (c chan engine.Request){
	s.workerChan = c
}
