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

	alphabets := [26]byte{}
	for i := range alphabets {
		alphabet := byte(65 + i)
		alphabets[i] = alphabet
	}
	chk := [...]byte{'C', 'F', 'I', 'L', 'O', 'S', 'V', 'Z'}
	var result int
	for _, v := range []byte(line[0]) {
		for i, cv := range chk {
			if v <= cv {
				result += 3 + i
				break
			}
		}
		// fmt.Fprintln(writer, string(v), result)
	}
	fmt.Fprintln(writer, result)

}
