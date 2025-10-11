package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	l, _ := reader.ReadString('\n')
	l = strings.TrimSpace(l)

	for i := 0; i < len(l)/2; i++ {
		if l[i] != l[len(l)-1-i] {
			fmt.Fprintln(writer, 0)
			return
		}
	}
	fmt.Fprintln(writer, 1)
}
