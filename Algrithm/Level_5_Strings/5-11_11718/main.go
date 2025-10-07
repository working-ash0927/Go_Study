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
	for {
		l, err := reader.ReadString('\n')
		if err != nil {
			return // 에러 발생
		}
		l = strings.TrimSpace(l)
		if l == "" {
			return // 입력 없으면 끝내기
		}
		fmt.Fprintln(writer, l)
		// line := strings.Fields(l)
		// fmt.Fprintln(writer, strings.Join(line, " "))
	}

}
