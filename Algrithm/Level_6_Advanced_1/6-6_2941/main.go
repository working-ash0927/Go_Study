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
	ref := []string{"c=", "c-", "dz=", "d-", "lj", "nj", "s=", "z="}

	var result int
	for _, v := range ref {
		result += strings.Count(l, v)
		l = strings.ReplaceAll(l, v, " ")
	}
	fmt.Fprintln(writer, len(strings.ReplaceAll(l, " ", ""))+result)

}
