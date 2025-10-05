package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var stdin *bufio.Reader = bufio.NewReader(os.Stdin)
	var stdout *bufio.Writer = bufio.NewWriter(os.Stdout)
	defer stdout.Flush()

	for {
		l, e := stdin.ReadString('\n')
		l = strings.TrimSpace(l)

		if e != nil || l == "" {
			break
		}
		fmt.Println(l)
		// var N int = fmt.Sscan(l, &N)
		// fmt.Sscan(l, &l)
	}
	// l, _ := stdin.ReadString('\n')
	// fmt.Println(l)
	// fmt.Println(N, a, b)
	// arr := make([]int,N)
	// for i := range arr {
	// 	arr[i] =
	// }

}
