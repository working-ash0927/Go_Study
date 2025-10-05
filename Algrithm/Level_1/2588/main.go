package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
	str_b := strconv.Itoa(b) // int to ascii
	for i := 2; i >= 0; i-- {
		bv, _ := strconv.Atoi(string(str_b[i]))
		pl(a * bv)
	}
	pl(a * b)

}
