// ABC083B - Some Sums
// https://atcoder.jp/contests/abs/tasks/abc083_b

package main

import (
	"fmt"
	"os"
)

func main() {
	var a, b, n int
	var s []int
	var t int
	_, err := fmt.Scanf("%d %d %d", &n, &a, &b)
	if err != nil {
		os.Exit(1)
	}

	for i := n; i > 0; i-- {
		s1 := 0
		for _, v := range digit(i, s) {
			s1 += v
		}
		if a <= s1 && s1 <= b {
			t += i
		}
	}
	fmt.Println(t)
}

func digit(i int, list []int) []int {
	if i > 0 {
		return digit(i/10, append(list, i%10))
	}
	return list
}
