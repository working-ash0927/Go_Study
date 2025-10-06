package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var scanner *bufio.Scanner = bufio.NewScanner(os.Stdin)
	var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var temp [42]int
	for i := range temp { // initialize
		temp[i] = 99
	}

	var N int
	for range 10 {
		scanner.Scan()
		line := scanner.Text()
		fmt.Sscan(line, &N)

		r := N % 42
		temp[r] = r
	}

	var cnt int
	for _, v := range temp {
		if v != 99 {
			cnt += 1
		}
	}
	fmt.Println(cnt)
}
