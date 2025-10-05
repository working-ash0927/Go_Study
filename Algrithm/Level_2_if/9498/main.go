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

	var a int
	fmt.Fscanln(reader, &a) // 개행문자로 읽기 종료, Fsanln은 스페이스, 탭으로 데이터를 구분.
	// fmt.Fscan(reader, &a, &b)

	if a%4 == 0 && (!(a%100 == 0) || (a%400 == 0)) {
		pl(1)
	} else {
		pl(0)
	}
}

/*
연도가 주어졌을 때, 윤년이면 1, 아니면 0을 출력하는 프로그램을 작성하시오.
윤년은 연도가 4의 배수이면서, 100의 배수가 아닐 때 또는 400의 배수일 때이다.
예를 들어, 2012년은 4의 배수이면서 100의 배수가 아니라서 윤년이다. 1900년은 100의 배수이고 400의 배수는 아니기 때문에 윤년이 아니다. 하지만, 2000년은 400의 배수이기 때문에 윤년이다.
*/
