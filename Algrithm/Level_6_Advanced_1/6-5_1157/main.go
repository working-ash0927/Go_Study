package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

// 중복제거 및 Count 함수로 인해 O(n3제곱) 으로 계산됨.
// 실제로 999,999 문자열을 계산 시, 0.7초로 계산되는데, 최적화하면 0.018초 까지 줄일 수 있음.
// 해시테이블인 map 을 활용해서 중복처리를 위한 별도 처리 없이 진행할것
func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	l, _ := reader.ReadString('\n')
	l = strings.ToLower(strings.TrimSpace(l))
	if len(l) == 1 {
		fmt.Fprintln(writer, strings.ToUpper(l))
		return
	}

	tmp := string(removeDuplicate(l))
	maxcnt := 0
	result := ""

	for _, v := range tmp {
		target := string(v)
		n := strings.Count(l, target)

		if n > maxcnt {
			maxcnt = n
			result = target
		} else if n == maxcnt {
			result = "?"
		}
	}

	fmt.Fprintln(writer, strings.ToUpper(result))
}

// 쓸데없대요
func removeDuplicate(str string) []byte {
	tmp := []byte(str)
	var result []byte
	for i, v := range tmp {
		if i == 0 {
			result = append(result, v)
		}
		if slices.Contains(result, v) {
			continue
		}
		result = append(result, v)
	}
	return result
}
