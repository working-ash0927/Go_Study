package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	scanner.Scan()
	field := strings.Fields(scanner.Text())
	N := field[0]
	B, _ := strconv.Atoi(field[1])

	result, _ := strconv.ParseInt(N, B, 64)
	fmt.Fprintln(writer, result)
}
