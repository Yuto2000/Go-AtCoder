// ABC085B - Kagami Mochi
// https://atcoder.jp/contests/abs/tasks/abc085_b
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	var x []int
	n := scan()
	for i := 0; i < n; i++ {
		x = append(x, scan())
	}

	m := make(map[int]struct{})
	unique := make([]int, 0, len(x))

	for _, v := range x {
		if _, ok := m[v]; ok {
			continue
		}
		unique = append(unique, v)
		m[v] = struct{}{}
	}

	fmt.Println(len(unique))
}

func scan() int {
	sc.Scan()
	i, err := strconv.Atoi(sc.Text())
	if err != nil {
		os.Exit(1)
	}
	return i
}
