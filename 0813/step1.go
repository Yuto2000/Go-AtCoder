// PracticeA - Welcome to AtCoder
// https://atcoder.jp/contests/abs/tasks/practice_1

package main

import (
	"fmt"
	"os"
)

func main() {
	var a, b, c int
	var s string

	// fmt.Println("aを入力してください")
	_, err := fmt.Scanf("%d", &a)
	if err != nil {
		os.Exit(1)
	}

	// fmt.Println("bcを半角スペースを間に入れて入力してください")
	_, err = fmt.Scanf("%d %d", &b, &c)
	if err != nil {
		os.Exit(1)
	}

	// fmt.Println("文字列sを入力してください")
	_, err = fmt.Scanf("%s", &s)
	if err != nil {
		os.Exit(1)
	}

	sum := sum(a, b, c)

	fmt.Printf("%d %s\n",sum, s)
}

func sum(a, b, c int) int {
	return a + b + c
}