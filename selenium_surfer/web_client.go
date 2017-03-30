package main

import (
	"time"
)

import "github.com/tebeka/selenium"

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
		
		Ligneous.Degub("[WebClient] Completing job")
		
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
	Ligneous.Info("[WebClient] Complete")
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
