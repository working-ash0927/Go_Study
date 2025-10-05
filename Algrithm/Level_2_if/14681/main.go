package main

import (
	"bufio"
	"fmt"
	"os"
)

var p = fmt.Print
var pl = fmt.Println
var pf = fmt.Printf

// defer: 이를 호출한 함수의 리턴 직전에 실행한다.
// Flush: 모든 데이터가 writer에 보내졌음을 의미 (꼭 써줘야한다!)
func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var a, b int
	// fmt.Fscanln(reader, &a) // 개행문자로 읽기 종료, Fsanln은 스페이스, 탭으로 데이터를 구분.
	fmt.Fscan(reader, &a, &b)

	// 조건을 만족할 때 실행할 코드
	if (a >= -1000 && a <= 1000 && a != 0) && (b >= -1000 && b <= 1000 && b != 0) {
		switch {
		case a > 0 && b > 0:
			pl(1)
		case a > 0 && b < 0:
			pl(4)
		case a < 0 && b < 0:
			pl(3)
		case a < 0 && b > 0:
			pl(2)
		}
	} else {
		return
	}

}
