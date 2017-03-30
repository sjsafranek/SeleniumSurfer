package main

// SurfTest
// surf docs: http://surf.readthedocs.io/overview/

import (
	"fmt"
	"io/ioutil"
	"time"
	//"github.com/PuerkitoBio/goquery"
	"gopkg.in/headzoo/surf.v1"
)

func savePage(filename, text string) {
	data := []byte(text)
	err := ioutil.WriteFile(filename, data, 0644)
	if nil != err {
		panic(err)
	}
}

func main() {

	bow := surf.NewBrowser()
	bow.AddRequestHeader("Accept", "text/html")
	bow.AddRequestHeader("Accept-Charset", "utf8")

	err := bow.Open("http://gmail.com")
	if err != nil {
		panic(err)
	}

	fmt.Println(bow.Title(), bow.Url())

	fm, err := bow.Form("form#gaia_loginform")
	if err != nil {
		panic(err)
	}

	fm.Input("Email", "sjsafranek@gmail.com")
	err = fm.Click("signIn")
	if nil != err {
		panic(err)
	}

	time.Sleep(2 * time.Second)

	fm.Input("Passwd", "geolRocks")
	err = fm.Click("signIn")
	if nil != err {
		panic(err)
	}

	err = fm.Submit()
	if nil != err {
		panic(err)
	}

	fmt.Println(bow.Title(), bow.Url())
	/*
		bow.Dom().Find("body").Each(func(n int, s *goquery.Selection) {
			//bow.Dom().Find("div#view_container").Each(func(n int, s *goquery.Selection) {
			ret, err := s.Html()
			if nil != err {
				panic(err)
			}
			fmt.Println(n, s.Text(), ret)
		})
	*/

	savePage("login.html", bow.Body())

	//fmt.Println(bow.Find("body").Text(), bow.Url())

	/*
		// Create a new browser and open reddit.
		bow := surf.NewBrowser()
		err := bow.Open("http://reddit.com")
		if err != nil {
			panic(err)
		}

		// Outputs: "reddit: the front page of the internet"
		fmt.Println(bow.Title())

		// Click the link for the newest submissions.
		bow.Click("a.new")

		// Outputs: "newest submissions: reddit.com"
		fmt.Println(bow.Title())

		// Log in to the site.
		fm, _ := bow.Form("form.login-form")
		fm.Input("user", "JoeRedditor")
		fm.Input("passwd", "d234rlkasd")
		if fm.Submit() != nil {
			panic(err)
		}

		// Go back to the "newest submissions" page, bookmark it, and
		// print the title of every link on the page.
		bow.Back()
		bow.Bookmark("reddit-new")
		bow.Find("a.title").Each(func(_ int, s *goquery.Selection) {
			fmt.Println(s.Text())
		})

	*/

}
