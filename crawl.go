package main

import (

	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions" 
)

func main() {

	crawler := colly.NewCollector()
	extensions.RandomUserAgent(c)
	crawler.WithTransport(&http.Transport{
	Proxy: http.ProxyFromEnvironment,
	DialContext: (&net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
		DualStack: true,
	}).DialContext,
	MaxIdleConns:          100,
	IdleConnTimeout:       90 * time.Second,
	TLSHandshakeTimeout:   10 * time.Second,
	ExpectContinueTimeout: 1 * time.Second,
}
}
