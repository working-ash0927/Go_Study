package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	l, _ := reader.ReadString('\n')
	l = strings.TrimSpace(l)
	N, _ := strconv.Atoi(l)

	length := 2*N - 1
	cnt := 1
	for i := 0; i < length/2; i++ {
		r := make([]string, cnt)
		for i := range r {
			r[i] = "*"
		}
		fmt.Fprint(writer, strings.Join(make([]string, N-i), " "))
		fmt.Fprintln(writer, strings.Join(r, ""))
		cnt += 2
	}
	for i := 0; i <= length/2; i++ {
		r := make([]string, cnt)
		for i := range r {
			r[i] = "*"
		}
		fmt.Fprint(writer, strings.Join(make([]string, 1+i), " "))
		fmt.Fprintln(writer, strings.Join(r, ""))
		cnt -= 2
	}

}
