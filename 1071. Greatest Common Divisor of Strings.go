package main

import (
	"fmt"
)

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func gcdOfStrings(str1 string, str2 string) string {
	if str1+str2 != str2+str1 {
		return ""
	}

	gcdLen := gcd(len(str1), len(str2))

	return str1[:gcdLen]
}

func main() {
	fmt.Println(gcdOfStrings("ABCABC", "ABC"))
	fmt.Println(gcdOfStrings("ABABAB", "ABAB"))
	fmt.Println(gcdOfStrings("LEET", "CODE"))
}
