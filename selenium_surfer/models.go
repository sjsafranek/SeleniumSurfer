package main

import (
	"sync"
)

import "github.com/tebeka/selenium"

type WebClientWorker struct {
	Queue     chan string
	id        int
	workwg    *sync.WaitGroup
	WebDriver selenium.WebDriver
	Uuid      string
}

type WebClientWorkerPool struct {
	pool   map[int]WebClientWorker
	jobs   chan string
	workwg *sync.WaitGroup
}

type Task struct {
	Url         string  `json:"url"`
	Search      string  `json:"search"`
	Probability float32 `json:"probability"`
}
