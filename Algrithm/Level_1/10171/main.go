package main

import (
	"fmt"
)

var p = fmt.Print
var pl = fmt.Println
var pf = fmt.Printf

// defer: 이를 호출한 함수의 리턴 직전에 실행한다.
// Flush: 모든 데이터가 writer에 보내졌음을 의미 (꼭 써줘야한다!)
func main() {
	// var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	// var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
	// defer writer.Flush()

	// 다중 문자열은 역슬래시
	pl(`\    /\
 )  ( ')
(  /  )
 \(__)|`)
}
