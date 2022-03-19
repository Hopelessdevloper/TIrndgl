package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// This will get called for each HTML element found

func processElement(index int, element *goquery.Selection) {
	// See if the p attribute exists on the element
	p, exists := element.Attr("p")
	if exists {
		fmt.Println(p)
	}
}

func main() {

	//Create log file
	fmt.Println("Created Log file", time.Now())
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)

	// Create Client
	log.Println("Client Created")

	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	// Stage HTTP request
	log.Println("HTTP Request Created")
	request, err := http.NewRequest("GET", "https://www.devdungeon.com", nil)
	if err != nil {
		log.Fatal(err)
	}

	request.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 12.3; rv:98.0) Gecko/20100101 Firefox/98.0")

	//Do the HTTP Request

	log.Println("Sent The HTTP Request")

	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	// Create a goquery document from the HTTP response

	log.Println("goquery document created form response body")
	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal("Error loading HTTP response body. ", err)
	}

	// Find all links and process them with the function
	// defined earlier
	document.Find("p").Each(processElement)
}
