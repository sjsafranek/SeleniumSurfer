package main

import (
	//"fmt"
	"sync"
	"time"
)

import "github.com/tebeka/selenium"

const (
	// DEFAULT_SELENIUM_SERVER default url for selenium server hub
	DEFAULT_SELENIUM_SERVER string = "http://127.0.0.1:4444/wd/hub"

	// DEFAULT_WEBDRIVER default selenium driver name
	DEFAULT_WEBDRIVER string = "firefox"
)

var (
	// SELENIUM_SERVER url for selenium server hub
	SELENIUM_SERVER string = DEFAULT_SELENIUM_SERVER

	// WEBDRIVER selenium driver name
	WEBDRIVER string = DEFAULT_WEBDRIVER
)

// NewWebClient creates a WebClient with a Selenium Web Driver
func NewWebClient(channel chan string, workwg *sync.WaitGroup) WebClientWorker {
	// FireFox driver without specific version
	// *** Add gecko driver here if necessary (see notes above.) ***
	caps := selenium.Capabilities{"browserName": WEBDRIVER}
	wd, err := selenium.NewRemote(caps, SELENIUM_SERVER)
	if err != nil {
		panic(err)
	}

	return WebClientWorker{Queue: channel, workwg: workwg, WebDriver: wd}
}

// Run starts WebClient worker
func (self WebClientWorker) Run() {
	Ligneous.Info("[WebClient] Starting google searches")
	self.workwg.Add(1)
	go self.run()
}

// run WebClient begins reading channel queue and processing jobs.
func (self WebClientWorker) run() {
	// get web driver
	wd := self.WebDriver

	// read items in queue
	for item := range self.Queue {

		Ligneous.Debug(`[WebClient] Searching for "` + item + `"`)

		// Get google.com
		wd.Get("https://www.google.com/")

		// find search input
		elem, _ := wd.FindElement(selenium.ByCSSSelector, `input[name="q"]`)
		elem.Clear()
		elem.SendKeys(item + selenium.EnterKey)

		// pause
		time.Sleep(5 * time.Second)
	}

	// pause
	time.Sleep(2 * time.Second)

	self.finish()
}

func (self WebClientWorker) finish() {
	Ligneous.Info("[WebClient] Shutting down")
	self.WebDriver.Quit()
	self.workwg.Done()
}

func (self WebClientWorker) Shutdown() {
	self.finish()
}

/*
	// Get simple playground interface
	wd.Get("http://play.golang.org/?simple=1")

	// Enter code in textarea
	elem, _ := wd.FindElement(selenium.ByCSSSelector, "#code")
	elem.Clear()
	elem.SendKeys("")

	// Click the run button
	btn, _ := wd.FindElement(selenium.ByCSSSelector, "#run")
	btn.Click()

	// Get the result
	div, _ := wd.FindElement(selenium.ByCSSSelector, "#output")

	output := ""
	// Wait for run to finish
	for {
		output, _ = div.Text()
		if output != "Waiting for remote server..." {
			break
		}
		time.Sleep(time.Millisecond * 100)
	}

	fmt.Printf("Got: %s\n", output)
*/
