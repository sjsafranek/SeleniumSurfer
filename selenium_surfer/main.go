package main

import (
	"bufio"
	"flag"
	"os"
	"os/signal"
	"strings"
	"sync"
	"time"
)

const (
	DEFAULT_CLIENT_NUMBER int    = 1
	DEFAULT_SEARCH_FILE   string = "search.txt"
)

var (
	CLIENT_NUMBER int
	SEARCH_FILE   string
	WCPool        WebClientWorkerPool
)

func init() {
	// Graceful shut down
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt)
	go func() {
		for sig := range sigs {
			// sig is a ^C, handle it
			Ligneous.Info("Recieved ", sig)
			Ligneous.Info("Gracefully shutting down")
			Ligneous.Info("Waiting for WebClients to shutdown...")
			WCPool.Shutdown()
			os.Exit(0)
		}
	}()
}

func shutDown() {
	time.Sleep(time.Millisecond * 100)
	WCPool.Shutdown()
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
		WCPool.Add(scanner.Text())
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
	if !strings.Contains("firefox chrome htmlunit", WEBDRIVER) {
		Ligneous.Error("Driver not supported")
		shutDown()
	}

	// create work group for workers
	var workwg sync.WaitGroup

	WCPool = newWebClientWorkerPool(CLIENT_NUMBER, &workwg)

	getSearchTerms()

	WCPool.Close()

	// wait for work groups to complete
	workwg.Wait()
}
