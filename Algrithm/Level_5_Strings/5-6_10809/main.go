package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// 입력 문자열들이 소문자 아스키 테이블의 몇번째에 있는지, 없는건 0 처리
func main() {
	var scanner *bufio.Scanner = bufio.NewScanner(os.Stdin)
	var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	scanner.Scan()
	l := scanner.Text()
	barr := []byte(l)

	tmp := make([]int, 26)
	for i := range tmp {
		tmp[i] = -1
	}

	start := byte('a') // ascii 97
	asciitable_lowcase := make([]int, 26)
	for i := range asciitable_lowcase {
		asciitable_lowcase[i] = int(start) + i
	}

	for i := range tmp {
		for j, v := range barr {
			if byte(asciitable_lowcase[i]) == v {
				if tmp[i] == -1 {
					tmp[i] = j
				}
			}
		}
	}

	result := make([]string, 26)
	for i, v := range tmp {
		s := strconv.Itoa(v)
		result[i] = s
	}
	fmt.Println(strings.Join(result, " "))
}
