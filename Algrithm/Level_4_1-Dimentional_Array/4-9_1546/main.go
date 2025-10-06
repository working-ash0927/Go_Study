package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	var scanner *bufio.Scanner = bufio.NewScanner(os.Stdin)
	var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	scanner.Scan()
	N, _ := strconv.Atoi(scanner.Text())
	arr := make([]float64, N)

	scanner.Scan()
	fields := strings.Fields(scanner.Text())
	for i, v := range fields {
		n, _ := strconv.Atoi(v)
		arr[i] += float64(n)
	}

	max_v := slices.Max(arr)
	var result float64

	for _, v := range arr {

		result += (v / max_v) * 100

	}
	fmt.Println(result / float64(N))
}
