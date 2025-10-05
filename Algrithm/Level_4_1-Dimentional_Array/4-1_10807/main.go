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
	N, _ := strconv.Atoi(scanner.Text())
	arr := make([]int, N)

	scanner.Scan()
	temp := strings.Fields(scanner.Text())

	for i, v := range temp {
		arr[i], _ = strconv.Atoi(v)
	}

	scanner.Scan()
	target, _ := strconv.Atoi(scanner.Text())

	var result int
	for _, v := range arr {
		if v == target {
			result += 1
		}
	}
	fmt.Println(result)
}
