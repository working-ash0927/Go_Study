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

	temp := make([]int, N+1)
	result := make([]string, N)
	for z := 0; z < M; z++ {
		scanner.Scan()
		line := scanner.Text()

		var i, j, k int
		fmt.Sscan(line, &i, &j, &k)
		for {
			if i == j {
				temp[i] = k
				break
			}
			temp[i] = k
			i++
		}
	}
	for i, v := range temp[1:] {
		result[i] = strconv.Itoa(v)
	}
	fmt.Println(strings.Join(result, " "))
}
