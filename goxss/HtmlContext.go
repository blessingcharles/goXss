package goxss

import (
	"fmt"
	"goXss/utils"
	"io/ioutil"
	"net/url"
	"regexp"
	"strings"
)

func BetweenTags(urlParsed *url.URL,
	QueryMap url.Values,
	paramKey string, reflectedTag string) {

}

func BetweenAttributes(urlParsed *url.URL,
	QueryMap url.Values,
	paramKey string, reflectedTag string) {

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
