package main

import "fmt"

func main() {
	var s string
	fmt.Scan(&s)
	//s := "yunus emre ozdiyar"
	fmt.Println(isReverse(s))
}

func Reverse(s string) string {
	var reverse string
	for i := len(s) - 1; i >= 0; i-- {
		reverse += string(s[i])
	}
	return reverse
}

func isReverse(s string) bool {
	if s == Reverse(s) {
		return true
	}
	return false
}
