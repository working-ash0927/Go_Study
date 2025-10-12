package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	l, _ := reader.ReadBytes('\n')
	data := bytes.Fields(l)
	N, _ := strconv.Atoi(string(data[0]))
	M, _ := strconv.Atoi(string(data[1]))

	arr := make([][]int, N)
	for i := range arr {
		arr[i] = make([]int, M)
	}

	for i := range N {
		l, _ := reader.ReadBytes('\n')
		data := bytes.Fields(l)

		for j, v := range data {
			n, _ := strconv.Atoi(string(v))
			arr[i][j] = n
		}
	}

	for i := range N {
		l, _ := reader.ReadBytes('\n')
		data := bytes.Fields(l)

		for j, v := range data {
			n, _ := strconv.Atoi(string(v))
			arr[i][j] += n
		}

		for _, v2 := range arr[i] {
			fmt.Fprint(writer, v2, " ")
		}
		fmt.Fprint(writer, "\n")
	}
}
