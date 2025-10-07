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
	line := strings.Fields(strings.TrimSpace(l))

	ra := reverseString(line[0])
	rb := reverseString(line[1])

	if ra > rb {
		fmt.Fprintln(writer, ra)
		return
	}
	fmt.Fprintln(writer, rb)
}
func reverseString(s string) string {
	bs := []byte(s)
	for i, j := 0, len(bs)-1; i < j; i, j = i+1, j-1 {
		bs[i], bs[j] = bs[j], bs[i]
	}
	return string(bs)
}
