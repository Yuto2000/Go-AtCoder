// ABC085C - Otoshidama
// https://atcoder.jp/contests/abs/tasks/abc085_c

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	sc.Split(bufio.ScanWords)
	var a int
	var b int
	result := []int{}

	a = scan()
	b = scan()

label:
	for i := a; i > -1; i-- {
		for j := a; j > -1; j-- {
			b1 := b
			b1 = b1 - i*10000 - j*5000 - (a-i-j)*1000
			if b1 == 0 && a-i-j >= 0 {
				result = append(result, i, j, a-i-j)
				break label
			}
		}
	}
	if len(result) == 0 {
		fmt.Println(-1, -1, -1)
	} else {
		fmt.Println(result[0], result[1], result[2])
	}
}

func scan() int {
	sc.Scan()
	i, err := strconv.Atoi(sc.Text())
	if err != nil {
		os.Exit(1)
	}
	return i
}
