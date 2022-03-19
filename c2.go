package main

import (
	"fmt"
	"log"
	"net/http"

	"time"

	"github.com/PuerkitoBio/goquery"
)

// This will get called for each HTML element found

func processElement(index int, element *goquery.Selection) {
    // See if the href attribute exists on the element
    href, exists := element.Attr("href")
    if exists {
        fmt.Println(href)
    }
}


func main() {

	// Create Client

	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	// Stage HTTP request

	request, err := http.NewRequest("GET", "https://www.devdungeon.com", nil)
	if err != nil {
		log.Fatal(err)
	}

	request.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 12.3; rv:98.0) Gecko/20100101 Firefox/98.0")

	//Do the HTTP Request

	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	// Create a goquery document from the HTTP response

	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal("Error loading HTTP response body. ", err)
	}

	// Find all links and process them with the function
	// defined earlier
	document.Find("a").Each(processElement)
}
