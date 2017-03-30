package main

import (
	"sync"
)

import "github.com/tebeka/selenium"

type WebClient struct {
	Queue     chan string
	id        int
	workwg    *sync.WaitGroup
	WebDriver selenium.WebDriver
}
