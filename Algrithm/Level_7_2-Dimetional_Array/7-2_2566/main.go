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
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	maxval, row, column := -1, 0, 0
	for i := range 9 {
		l, _ := reader.ReadString('\n')
		field := strings.Fields(l)

		parse := make([]int, len(l))
		for i, v := range field {
			n, _ := strconv.Atoi(v)
			parse[i] = n
		}

		// fmt.Fprintln(writer, i, slices.Max(parse))

		if slices.Max(parse) > maxval {
			maxval = slices.Max(parse)
			row = i + 1
			column = slices.Index(parse, maxval) + 1
		}
	}

	fmt.Fprintln(writer, maxval)
	fmt.Fprintln(writer, row, column)

}
