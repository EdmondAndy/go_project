package main

import "fmt"

func main() {
	var input string
	fmt.Scanf("%s\n", &input)
	answer := 1
	for _, ch := range input {
		if ch >= 'A' && ch <= 'Z' {
			answer++
		}
	}
	fmt.Println("Number of words is:", answer)
}