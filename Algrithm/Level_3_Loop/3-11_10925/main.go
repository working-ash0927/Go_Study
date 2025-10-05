package main

import (
	"bufio"
	"fmt"
	"os"
)

// defer: 이를 호출한 함수의 리턴 직전에 실행한다.
// Flush: 모든 데이터가 writer에 보내졌음을 의미 (꼭 써줘야한다!)
var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()
	var a, b int
	for {
		fmt.Fscan(reader, &a, &b) // 개행문자로 읽기 종료, Fsanln은 스페이스, 탭으로 데이터를 구분.

		if a == b && a == 0 {
			break
		}
		if (a < 0 && a > 10) || (b < 0 && b > 10) {
			break
		}
		fmt.Println(a + b)
	}
}
