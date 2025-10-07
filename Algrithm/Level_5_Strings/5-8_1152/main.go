package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// 입력값 큰지 작은지 먼저 파악하기.....!!!!!
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	l, _ := reader.ReadString('\n')
	l = strings.TrimSpace(l)
	r := strings.Fields(l)
	fmt.Fprintln(writer, len(r))
}
