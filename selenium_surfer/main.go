package main

import (
	"bufio"
	"flag"
	"os"
	"strings"
	"sync"
	"time"
)

const (
	DEFAULT_CLIENT_NUMBER int    = 1
	DEFAULT_SEARCH_FILE   string = "search.txt"
)

var (
	CLIENT_NUMBER    int
	SEARCH_FILE      string
	WebDriverCluster map[int]WebClientWorker
	SearchItems      chan string
	Pool             WebClientWorkerPool
)

func init() {
	// create web driver cluster
	WebDriverCluster = make(map[int]WebClientWorker)

	// create channel
	SearchItems = make(chan string, 100)
}

func shutDown() {
	time.Sleep(time.Millisecond * 100)
	os.Exit(1)
}

func getSearchTerms() {
	Ligneous.Info("Reading search terms from file", SEARCH_FILE)

	if !fileExists(SEARCH_FILE) {
		Ligneous.Error("Search file not found", SEARCH_FILE)
		shutDown()
	}

	file, err := os.Open(SEARCH_FILE)
	if err != nil {
		Ligneous.Error(err)
		shutDown()
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		//SearchItems <- scanner.Text()
		Pool.Add(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		Ligneous.Error(err)
		shutDown()
	}
}

func main() {

	// command line args
	flag.StringVar(&SEARCH_FILE, "f", DEFAULT_SEARCH_FILE, "search file")
	flag.StringVar(&WEBDRIVER, "b", DEFAULT_WEBDRIVER, "web driver")
	flag.IntVar(&CLIENT_NUMBER, "n", DEFAULT_CLIENT_NUMBER, "number of clients")
	flag.StringVar(&SELENIUM_SERVER, "s", DEFAULT_SELENIUM_SERVER, "selenium server url")
	flag.Parse()

	// check for valid web driver option
	if !strings.Contains("firefox chrome", WEBDRIVER) {
		Ligneous.Error("Driver not supported")
		shutDown()
	}

	// create work group for workers
	var workwg sync.WaitGroup

	/*
		// create web clients
		for i := 0; i < CLIENT_NUMBER; i++ {
			WebDriverCluster[i] = NewWebClient(SearchItems, &workwg)
		}

		for i := 0; i < CLIENT_NUMBER; i++ {
			WebDriverCluster[i].Run()
		}
	*/
	Pool = newWebClientWorkerPool(CLIENT_NUMBER, &workwg)

	getSearchTerms()

	Pool.Close()
	//close(SearchItems)

	// wait for work groups to complete
	workwg.Wait()
}
