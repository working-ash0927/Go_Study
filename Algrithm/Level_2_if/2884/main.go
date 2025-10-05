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

	var H, M int
	fmt.Fscanln(reader, &H, &M) // 개행문자로 읽기 종료, Fsanln은 스페이스, 탭으로 데이터를 구분.
	// fmt.Fscan(reader, &H, &M)

	if H >= 0 && H <= 23 && M >= 0 && M <= 59 {
		M -= 45
		if M < 0 {
			M += 60
			H -= 1
		} else if M > 59 {
			M -= 60
			H += 1
		}
		if H > 23 {
			H = 0
		} else if H < 0 {
			H += 24
		}
		pl(H, M)
	} else {
		return
	}
}
