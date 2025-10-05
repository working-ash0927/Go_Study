package main

import (
	"bufio"
	"fmt"
	"os"
)

var pl = fmt.Println

// defer: 이를 호출한 함수의 리턴 직전에 실행한다.
// Flush: 모든 데이터가 writer에 보내졌음을 의미 (꼭 써줘야한다!)

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var H, M, t int
	fmt.Fscanln(reader, &H, &M) // 개행문자로 읽기 종료, Fsanln은 스페이스, 탭으로 데이터를 구분.
	fmt.Fscan(reader, &t)

	if (H >= 0 && H <= 23) && (M >= 0 && M <= 59) && (t >= 0 && t <= 1000) {
		M += t //초과된 분
		H += M / 60

		if M >= 60 {
			M %= 60
		}
		if H > 23 {
			H -= 24
		}

		pl(H, M)
	} else {
		return
	}
}
