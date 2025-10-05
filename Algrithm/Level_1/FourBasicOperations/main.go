package main

import (
	"bufio"
	"fmt"
	"os"
)

var p = fmt.Print
var pl = fmt.Println
var pf = fmt.Printf

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

	// defer: 이를 호출한 함수의 리턴 직전에 실행한다.
	// Flush: 모든 데이터가 writer에 보내졌음을 의미 (꼭 써줘야한다!)
	defer writer.Flush()

	var a, b int
	if a < 0 || b > 10 {
		return
	}
	fmt.Fscanln(reader, &a, &b) // Fsanln은 스페이스, 탭으로 데이터를 구분. 개행문자로 읽기 종료
	pl(a + b)
	pl(a - b)
	pl(a * b)
	pl(a / b)
	pl(a % b)
}
