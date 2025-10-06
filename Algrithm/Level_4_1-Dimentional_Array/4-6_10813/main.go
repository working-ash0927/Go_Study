package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var scanner *bufio.Scanner = bufio.NewScanner(os.Stdin)
	var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	scanner.Scan()
	line := scanner.Text()
	var N, M int
	fmt.Sscan(line, &N, &M)

	arr := make([]string, N)
	for i := range N {
		n := strconv.Itoa(i + 1)
		arr[i] = n
	}

	for range M {
		scanner.Scan()
		line := scanner.Text()
		var i, j int
		var tmp string
		fmt.Sscan(line, &i, &j)

		tmp = arr[i-1]
		arr[i-1] = arr[j-1]
		arr[j-1] = tmp
	}
	fmt.Println(strings.Join(arr, " "))
}
