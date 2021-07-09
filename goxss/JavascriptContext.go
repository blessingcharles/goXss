package goxss

import (
	"bufio"
	"fmt"
	"goXss/utils"
	"io/ioutil"
	"net/url"
	"os"
	"regexp"
)

func ReflectionJavascript(urlStr string, payload string) ([][]string, bool) {

	res, err := utils.DoGetReq(urlStr, 10)

	if err != nil {
		return nil, false
	}

	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return nil, false
	}

	bodyString := string(bodyBytes)

	matchRegex := `<script>\s*.*` + payload
	r, _ := regexp.Compile(matchRegex)
	a := r.FindAllStringSubmatch(bodyString, -1)

	return a, true

}

func BruteForceScript(urlParsed *url.URL,
	QueryMap url.Values,
	paramKey string) {

	fmt.Println("[*] using common javascript escaping payloads ")

	file, err := os.Open("./Payloads/javascriptEscape.txt")
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

func JavascriptContext() {

}
