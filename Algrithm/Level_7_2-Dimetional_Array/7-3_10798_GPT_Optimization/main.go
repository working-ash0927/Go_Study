package main

import (
	"bufio"
	"os"
)

func main() {
	reader := bufio.NewScanner(os.Stdin)
	reader.Buffer(make([]byte, 0, 1024), 1024*1024)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	data := make([]string, 5)
	maxlen := 0

	for i := range 5 {
		reader.Scan()
		data[i] = reader.Text() // 한줄씩 읽어서 가장 긴 입력값 길이 체크
		if len(data[i]) > maxlen {
			maxlen = len(data[i])
		}
	}

	for col := range maxlen {
		for row := range 5 {
			if col < len(data[row]) {
				writer.WriteByte(data[row][col])
			}
		}
	}
}
