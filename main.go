package main

import (
	"bufio"
	"crypto/tls"
	"flag"
	"fmt"
	"goXss/goxss"
	"net/http"
	"net/url"
	"os"
)

func main() {

	var proxy string

	//	utils.Banner()

	flag.StringVar(&proxy, "proxy", "", "proxy configuration")

	flag.Parse()

	if proxy != "" {
		//setting default proxy and ignore ssl secure
		proxyUrl, err := url.Parse(proxy)

		if err != nil {
			panic(err)
		}

		http.DefaultTransport = &http.Transport{

			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			Proxy:           http.ProxyURL(proxyUrl),
		}

		fmt.Println("[+]proxy --> ", proxy)

	} else {
		http.DefaultTransport = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			Proxy:           nil,
		}
	}

	StdinScanner := bufio.NewScanner(os.Stdin)

	//reading from stdin
	// pipe the urls
	for StdinScanner.Scan() {

		urlStr := StdinScanner.Text()

		urlParsed, err := url.Parse(urlStr)

		if err != nil {
			println(err)
			continue
		}

		QueryMap, err := url.ParseQuery(urlParsed.RawQuery)

		if err != nil {
			println(err)
			continue
		}

		for key := range QueryMap {

			println("[+] checking param :", key)
			goxss.XssScanner(urlStr, key)

		}

	}

}
