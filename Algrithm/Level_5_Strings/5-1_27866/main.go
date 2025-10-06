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

	var N, M int
	scanner.Scan()
	line := scanner.Text()
	fmt.Sscan(line, &N, &M)

	arr := make([]int, N)
	for i := range arr {
		arr[i] = i + 1
	}

}
