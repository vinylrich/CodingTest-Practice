package main

import "fmt"

func main() {
	a, b := 0, 0
	fmt.Scanf("%d %d", &a, &b)
	if a < b {
		fmt.Println("<")
	}
	if a == b {
		fmt.Println("==")
	}
	if a > b {
		fmt.Println(">")
	}
}
