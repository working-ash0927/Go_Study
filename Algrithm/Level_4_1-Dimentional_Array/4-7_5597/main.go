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

	var students [30]int
	var N int
	for range 28 {
		scanner.Scan()
		line := scanner.Text()
		fmt.Sscan(line, &N)
		students[N-1] = N
	}

	for i, v := range students {
		if v == 0 {
			fmt.Println(i + 1)
		}
	}
}
