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

	scanner.Scan()
	l := scanner.Text()
	result := []byte(l)
	fmt.Print(result[0])
}
