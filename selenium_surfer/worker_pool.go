package main

import (
	"sync"
)

func newWebClientWorkerPool(number_of_workers int, workwg *sync.WaitGroup) WebClientWorkerPool {
	wcwp := WebClientWorkerPool{
		pool:   make(map[int]WebClientWorker),
		jobs:   make(chan string, 100),
		workwg: workwg}
	wcwp.Init(number_of_workers)
	return wcwp
}

func (self WebClientWorkerPool) Init(num int) {
	Ligneous.Debug(`[WebClientPool] Creating WebClients`)
	for i := 0; i < num; i++ {
		self.pool[i] = NewWebClient(self.jobs, self.workwg)
	}

	Ligneous.Debug(`[WebClientPool] Starting WebClients`)
	for i := 0; i < num; i++ {
		self.pool[i].Run()
	}
}

func (self WebClientWorkerPool) Add(message string) {
	self.jobs <- message
}

func (self WebClientWorkerPool) Close() {
	Ligneous.Debug(`[WebClientPool] Closing job channel`)
	//_, ok := <-self.jobs
	if nil != self.jobs {
		close(self.jobs)
		self.jobs = nil
	}
	//close(self.jobs)
}

func (self WebClientWorkerPool) Shutdown() {
	self.Close()
	Ligneous.Debug(`[WebClientPool] Shutting down WebClients`)
	for i := range self.pool {
		self.pool[i].Shutdown()
	}
}
