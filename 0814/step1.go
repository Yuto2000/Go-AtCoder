// ABC081B - Shift only

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
	n := scan()
	var x []int
	var counter int
	var scounter int

	for i := 0; i < n; i++ {
		x = append(x, scan())
	}

eternal:
	for {
		for i, v := range x {
			if a := v % 2; a == 0 {
				x[i] = v / 2
				counter++
			} else {
				break eternal
			}
		}
		if counter == len(x) {
			scounter++
		}
		counter = 0
	}
	fmt.Println(scounter)
}

func scan() int {
	sc.Scan()
	i, err := strconv.Atoi(sc.Text())
	if err != nil {
		os.Exit(1)
	}
	return i
}
