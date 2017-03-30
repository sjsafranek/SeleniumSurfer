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
	for i := 0; i < num; i++ {
		self.pool[i] = NewWebClient(self.jobs, self.workwg)
	}

	for i := 0; i < num; i++ {
		self.pool[i].Run()
	}
}

func (self WebClientWorkerPool) Add(message string) {
	self.jobs <- message
}

func (self WebClientWorkerPool) Close() {
	close(self.jobs)
}
