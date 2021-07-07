package goxss

import (
	"bufio"
	"fmt"
	"goXss/utils"
	"io/ioutil"
	"net/url"
	"os"
	"regexp"
	"strings"
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

func ReflectionBetweenTag(urlStr string, payload string) (string, bool) {

	res, err := utils.DoGetReq(urlStr, 10)

	if err != nil {
		return "", false
	}

	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return "", false
	}

	bodyString := string(bodyBytes)

	matchRegex := "<[^<]*" + payload + ".*?</(.*?)>"
	r, _ := regexp.Compile(matchRegex)
	a := r.FindAllStringSubmatch(bodyString, -1)[0]

	var reflectedTag string
	for _, tag := range a {
		reflectedTag = tag
	}

	return reflectedTag, true

}

func ReflectionInAtributes(urlStr string, payload string) (string, string, bool) {

	res, err := utils.DoGetReq(urlStr, 10)

	fmt.Println(urlStr)
	if err != nil {
		return "", "", false
	}

	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return "", "", false
	}

	bodyString := string(bodyBytes)

	matchRegex := `<[^<]*\s([^<^>]*)="` + payload + `"`

	r, _ := regexp.Compile(matchRegex)
	a := r.FindAllStringSubmatch(bodyString, -1)

	if len(a) <= 0 {
		return "", "", false
	}

	var reflectedTag string
	var attribute string

	for _, ele := range a {
		reflectedTag = ele[0]
		attribute = ele[1]
		// fmt.Println(ele[0])
		// fmt.Println(ele[1])
	}

	return reflectedTag, attribute, true

}

func CheckReflection(urlStr string, payload string) bool {

	res, err := utils.DoGetReq(urlStr, 10)

	if err != nil {
		fmt.Println(err)
		return false
	}

	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return false
	}

	bodyString := string(bodyBytes)

	//checking if payload gets reflected in the response body
	match := strings.Contains(bodyString, payload)

	return match

}
