// ABC088B - Card Game for Two
// https://atcoder.jp/contests/abs/tasks/abc088_b

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	sc.Split(bufio.ScanWords)
	n := scan()
	var x []int
	var a int
	var b int

	for i := 0; i < n; i++ {
		x = append(x, scan())
	}

	sort.Sort(sort.IntSlice(x))
	reverse(x)
	for i, v := range x {
		if i%2 == 0 {
			a += v
		} else {
			b += v
		}
	}

	fmt.Println(a - b)
}

func scan() int {
	sc.Scan()
	i, err := strconv.Atoi(sc.Text())
	if err != nil {
		os.Exit(1)
	}
	return i
}

func reverse(list []int) []int {
	for i := len(list)/2 - 1; i >= 0; i-- {
		opp := len(list) - i - 1
		list[i], list[opp] = list[opp], list[i]
	}
	return list
}
