package main

import "fmt"

func reverseString(str string) string {
	if len(str) == 0 {
		return ""
	}
	return string(str[len(str)-1]) + reverseString(str[:len(str)-1])
}

func main() {
	fmt.Println(reverseString("yoyo mastery"))
}
