// Run some code on play.golang.org and display the result
package main

import (
	"flag"
	"os"
	"strings"
	"sync"
	"time"
)

import "github.com/tebeka/selenium"

const (
	DEFAULT_SELENIUM_SERVER string = "http://127.0.0.1:4444/wd/hub"
	DEFAULT_WEBDRIVER       string = "firefox"
	DEFAULT_CLIENT_NUMBER   int    = 1
)

var (
	SELENIUM_SERVER  string
	WEBDRIVER        string
	CLIENT_NUMBER    int
	WebDriverCluster map[int]WebClient
	SearchItems      chan string
)

func init() {
	WebDriverCluster = make(map[int]WebClient)
}

func LaunchWebDriverClient(channel chan string, workwg *sync.WaitGroup) WebClient {
	// FireFox driver without specific version
	// *** Add gecko driver here if necessary (see notes above.) ***
	caps := selenium.Capabilities{"browserName": WEBDRIVER}
	//caps := selenium.Capabilities{
	//	"browserName":            "firefox",
	//	"webdriver.gecko.driver": "geckodriver",
	//}
	wd, err := selenium.NewRemote(caps, SELENIUM_SERVER)
	if err != nil {
		panic(err)
	}

	//wc := WebClient{Queue: make(chan string), workwg: workwg, WebDriver: wd}
	wc := WebClient{Queue: channel, workwg: workwg, WebDriver: wd}
	return wc
}

func main() {

	// command line args
	flag.StringVar(&WEBDRIVER, "b", DEFAULT_WEBDRIVER, "web driver")
	flag.IntVar(&CLIENT_NUMBER, "n", DEFAULT_CLIENT_NUMBER, "number of clients")
	flag.StringVar(&SELENIUM_SERVER, "s", DEFAULT_SELENIUM_SERVER, "selenium server url")
	flag.Parse()

	// check for valid web driver option
	if !strings.Contains("firefox chrome", WEBDRIVER) {
		err := newSurferError("Driver not supported")
		Ligneous.Error(err)
		time.Sleep(time.Millisecond * 100)
		os.Exit(1)
	}

	// create work group for workers
	var workwg sync.WaitGroup

	// create channel
	SearchItems = make(chan string)

	// create web clients
	for i := 0; i < CLIENT_NUMBER; i++ {
		WebDriverCluster[i] = LaunchWebDriverClient(SearchItems, &workwg)
	}

	for i := 0; i < CLIENT_NUMBER; i++ {
		WebDriverCluster[i].Run()
	}

	SearchItems <- "TEST"
	SearchItems <- "apples"
	SearchItems <- "mac"
	close(SearchItems)

	// wait for work groups to complete
	workwg.Wait()
}
