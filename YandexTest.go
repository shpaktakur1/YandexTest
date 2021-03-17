package main

import (
	"fmt"
	"regexp"
)

const refString = `[{ \"email\": \"shpaktakur1@yandex.ru\" \"phone\": 89651112233},
{ \"email\": \"other@domain.com\" \"phone\": 89161114455}]`

func main() {

	emailRegexp := regexp.MustCompile("Ya[A-Z]+_[a-zA-Z]+")
	first := emailRegexp.FindString(refString)
	fmt.Println("First: ")
	fmt.Println(first)

	all := emailRegexp.FindAllString(refString, - 1)
	fmt.Println("All: ")
	for _, val := range all {
		fmt.Println(val)
	}
}