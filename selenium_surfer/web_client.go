package main

import (
	"fmt"
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

// NewWebClient creates  
func NewWebClient(channel chan string, workwg *sync.WaitGroup) WebClient {
	// FireFox driver without specific version
	// *** Add gecko driver here if necessary (see notes above.) ***
	caps := selenium.Capabilities{"browserName": WEBDRIVER}
	wd, err := selenium.NewRemote(caps, SELENIUM_SERVER)
	if err != nil {
		panic(err)
	}

	//wc := WebClient{Queue: channel, workwg: workwg, WebDriver: wd}
	//return wc
	return WebClient{Queue: channel, workwg: workwg, WebDriver: wd}
}

func (self WebClient) Run() {
	Ligneous.Info("[WebClient] Starting google searches")
	self.workwg.Add(1)
	go self.run()
}

func (self WebClient) run() {
	// get web driver
	wd := self.WebDriver

	// read items in queue
	for item := range self.Queue {

		msg := fmt.Sprintf(`[WebClient] Searching "%v"`, item)
		Ligneous.Debug(msg)
		//Ligneous.Debug(`[WebClient] Searching "` + item + `"`)

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

func (self WebClient) finish() {
	Ligneous.Info("[WebClient] Shutting down")
	self.WebDriver.Quit()
	self.workwg.Done()
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
