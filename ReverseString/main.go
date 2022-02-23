package main

import "fmt"

func reverse(str string) string {
	result := ""
	for i := len(str) - 1; i >= 0; i-- {
		result += string(str[i])
	}
	return result
}

func main() {
	fmt.Println(reverse("my name is andyyyyy"))
}
