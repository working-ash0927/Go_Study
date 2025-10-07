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
	T, _ := strconv.Atoi(scanner.Text())

	for range T {
		scanner.Scan()

		line := strings.Fields(scanner.Text())
		r, _ := strconv.Atoi(line[0])
		s := line[1]

		var builder strings.Builder
		for _, v := range s {
			for range r {
				builder.WriteRune(v)
			}
		}
		fmt.Fprintln(writer, builder.String())
	}

}
