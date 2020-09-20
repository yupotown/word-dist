package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var sc *bufio.Scanner

func prompt(msg string) string {
	fmt.Printf("%s > ", msg)
	sc.Scan()
	return strings.TrimSpace(sc.Text())
}

func dist(a, b string) int {
	res := 0
	l := len(a)
	if len(b) < l {
		l = len(b)
	}
	for i := 0; i < l; i++ {
		d := int(b[i]) - int(a[i])
		if d < 0 {
			d = -d
		}
		res += d
	}
	return res
}

func main() {
	sc = bufio.NewScanner(os.Stdin)

	ans := prompt("answer")

	for {
		word := prompt("word?")
		fmt.Printf("distance: %d\n", dist(ans, word))
	}
}
