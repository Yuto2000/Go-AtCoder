// ABC086A - Product

package main

import (
	"fmt"
	"os"
)

func main() {
	var a, b int
	_, err := fmt.Scanf("%d %d", &a, &b)
	if err != nil {
		os.Exit(1)
	}
	if c := a*b%2; c == 0 {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}
}