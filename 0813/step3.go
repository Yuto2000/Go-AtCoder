// ABC081A - Placing Marbles

package main

import (
	"fmt"
	"os"
)

func main() {
	var s int
	var ss []int
	var counter int

	fmt.Scanf("%d", &s)

	for _, v := range reverse(digit(s, ss)) {
		switch v {
		case 1:
			counter++
		case 0:
		default:
			os.Exit(1)
		}
	}
	fmt.Println(counter)
}

func digit(i int, list []int) []int {
	if i > 0 {
		return digit(i/10, append(list, i%10))
	}
	return list
}

func reverse(list []int) []int {
	for i := len(list)/2 - 1; i >= 0; i-- {
		opp := len(list) - i - 1
		list[i], list[opp] = list[opp], list[i]
	}
	return list
}
