// ABC049C - 白昼夢
// https://atcoder.jp/contests/abs/tasks/arc065_a

package main

import (
	"bufio"
	"fmt"
	"os"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	dream := "maerd"
	erase := "esare"
	dreamer := "remaerd"
	eraser := "resare"
	sc.Scan()
	s1 := sc.Text()
	s := reverse(s1)
	for {
		if len(s) >= 5 && string(s[:5]) == dream {
			s = s[5:]

		} else if len(s) >= 7 && string(s[:7]) == dreamer {
			s = s[7:]

		} else if len(s) >= 5 && string(s[:5]) == erase {
			s = s[5:]

		} else if len(s) >= 6 && string(s[:6]) == eraser {
			s = s[6:]

		} else if len(s) == 0 {
			fmt.Println("YES")
			break

		} else {
			fmt.Println("NO")
			break
		}
	}
}

func reverse(s string) []rune {
	rs := []rune(s)
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		rs[i], rs[j] = rs[j], rs[i]
	}
	return rs
}
