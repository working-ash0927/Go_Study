package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// 해시 테이블은 순서 보장이 없기에 이러한 구조에서 권장되지 않고 쓰잘데기 없이 복잡함
// 그냥 단순히 make[]string 으로 하는게 나음
func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	data := make(map[int]string)
	var maxlen int
	for i := range 5 {
		l, _ := reader.ReadString('\n')
		l = strings.TrimSpace(l)
		data[i] = l
		if maxlen < len(l) {
			maxlen = len(l)
		}
	}
	// 공백을 임시값으로 할당
	for i, v := range data {
		if len(v) != maxlen {
			var builder strings.Builder
			builder.WriteString(data[i])

			for range maxlen - len(v) {
				builder.WriteString("_")
				data[i] = builder.String()
			}

		}
	}

	/*
		map[
			0:AABCDD
			1:afzz
			2:09121
			3:a8EWg6
			4:P5h3kx
		]
	*/
	for i := range maxlen {
		for j := range len(data) {
			v := string(data[j][i])
			if v == "_" {
				continue
			} else {
				fmt.Fprint(writer, v)
			}
		}
	}
}
