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

	var a, b, c int
	fmt.Fscanln(reader, &a, &b, &c) // Fsanln은 스페이스, 탭으로 데이터를 구분. 개행문자로 읽기 종료
	if a >= 2 && a <= 10000 && b >= 2 && b <= 10000 && c >= 2 && c <= 10000 {
		pl((a + b) % c)
		pl(((a % c) + (b % c)) % c)
		pl((a * b) % c)
		pl(((a % c) * (b % c)) % c)
	} else {
		return
	}

}
