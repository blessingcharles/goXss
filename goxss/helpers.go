package goxss

import (
	"bufio"
	"fmt"
	"net/url"
	"os"
)

// bruteforce with common payloads

func BruteForce(urlParsed *url.URL,
	QueryMap url.Values,
	paramKey string, reflectedTag string) {

	fmt.Println("[*] using common payloads ")

	file, err := os.Open("./Payloads/common.txt")
	if err != nil {
		fmt.Println("failed to open file")
		return
	}

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {

		QueryMap.Set(paramKey, scanner.Text())
		urlParsed.RawQuery = QueryMap.Encode()
		if CheckReflection(urlParsed.String(), scanner.Text()) {
			fmt.Println("[!]Potential Vector : ", urlParsed.String())
		}

	}
}

func BruteForceAttributes(urlParsed *url.URL,
	QueryMap url.Values,
	paramKey string, reflectedTag string) {

	fmt.Println("[*] using common attributes payloads ")

	file, err := os.Open("./Payloads/attributes.txt")
	if err != nil {
		fmt.Println("failed to open file")
		return
	}

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {

		QueryMap.Set(paramKey, scanner.Text())
		urlParsed.RawQuery = QueryMap.Encode()
		if CheckReflection(urlParsed.String(), scanner.Text()) {
			fmt.Println("[!]Potential Vector : ", urlParsed.String())
		}

	}
}
