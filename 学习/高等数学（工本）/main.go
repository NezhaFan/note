package main

import "fmt"

func main() {
	for i := 8330; i <= 8380; i++ {
		fmt.Println(string(rune(i)))
	}
	// fmt.Println("ₓ", []rune("ₓ"))
}
