// ABC087B - Coins
// https://atcoder.jp/contests/abs/tasks/abc087_b

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	var counter int
	// a: A b: B c: C d: X
	a, b, c, d := Line(), Line(), Line(), Line()

	for i := a; i > -1; i-- {
		for j := b; j > -1; j-- {
			for k := c; k > -1; k-- {
				d1 := d
				d1 = d1 - i*500 - j*100 - k*50
				if d1 == 0 {
					counter++
				}
			}
		}
	}

	fmt.Println(counter)

}

func Line() int {
	sc.Scan()
	i, err := strconv.Atoi(sc.Text())
	if err != nil {
		os.Exit(1)
	}
	return i
}
