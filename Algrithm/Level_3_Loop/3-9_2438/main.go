package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// defer: 이를 호출한 함수의 리턴 직전에 실행한다.
// Flush: 모든 데이터가 writer에 보내졌음을 의미 (꼭 써줘야한다!)
var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()

	var N int
	fmt.Fscan(reader, &N) // 개행문자로 읽기 종료, Fsanln은 스페이스, 탭으로 데이터를 구분.

	z := 0
	for z < N {
		a := make([]string, z+1)
		for i := range a {
			a[i] = "*"
		}

		fmt.Println(strings.Join(a, ""))
		z++
	}
}
