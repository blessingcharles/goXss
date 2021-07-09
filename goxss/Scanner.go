package goxss

import (
	"fmt"
	"net/url"
)

func XssScanner(urlStr string, paramKey string) {

	urlParsed, _ := url.Parse(urlStr)
	QueryMap, err := url.ParseQuery(urlParsed.RawQuery)
	if err != nil {
		return
	}

	//checking for basic reflection
	payload := "th3h04x"
	QueryMap.Set(paramKey, payload)
	urlParsed.RawQuery = QueryMap.Encode()

	fmt.Println("[*]checking for Reflection:", urlParsed.String())

	if !CheckReflection(urlParsed.String(), payload) {

		fmt.Println("[-]No Reflection in Response Body")
		return

	}

	// html context

	//[+] Between Tags
	// fmt.Println("[*]checking for Reflection Between Tags")

	// reflectedTag, OK := ReflectionBetweenTag(urlParsed.String(), payload)

	// if OK {
	// 	fmt.Println("[!]Reflected Between Tag", reflectedTag)

	// 	fmt.Println("[*] BruteForce Payloads ")
	// 	BruteForce(urlParsed, QueryMap, paramKey, reflectedTag)

	// 	fmt.Println("[*]Starting Intelligent Fuzz")
	// 	BetweenTags(urlParsed, QueryMap, paramKey, reflectedTag)

	// } else {
	// 	fmt.Println("[-]No Reflection Between Tags")
	// }

	// //[+] Html Attributes
	// fmt.Println("[+]checking for Reflection in Attributes")

	// QueryMap.Set(paramKey, payload)
	// urlParsed.RawQuery = QueryMap.Encode()

	// reflectedTag, attribute, OK := ReflectionInAtributes(urlParsed.String(), payload)

	// if OK {
	// 	fmt.Println("[!]Reflected Tag", reflectedTag)
	// 	fmt.Println("[!]Attribute", attribute)

	// 	fmt.Println("[*] BruteForce Payloads ")
	// 	BruteForceAttributes(urlParsed, QueryMap, paramKey, reflectedTag)

	// 	fmt.Println("[*]Starting Intelligent Fuzz")
	// 	BetweenAttributes(urlParsed, QueryMap, paramKey, reflectedTag)

	// } else {
	// 	fmt.Println("[-]No Reflection Tags Attributes")
	// }

	//[+] javascript Context

	fmt.Println("[+]In Javascript Context")
	QueryMap.Set(paramKey, payload)
	urlParsed.RawQuery = QueryMap.Encode()
	Reflected, OK := ReflectionJavascript(urlParsed.String(), payload)

	if OK {
		fmt.Println("[!]Script Context", Reflected)

		fmt.Println("[*] BruteForce Payloads ")
		BruteForceScript(urlParsed, QueryMap, paramKey)

		fmt.Println("[*]Starting Intelligent Fuzz")
		JavascriptContext()
	}
}
