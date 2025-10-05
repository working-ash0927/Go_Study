package main

import (
	"bufio"
	"fmt"
	"os"
)

var pl = fmt.Println
var pf = fmt.Printf // 개행문자가 필요하나 fmt.Sprintf 는 string print formatted

// defer: 이를 호출한 함수의 리턴 직전에 실행한다.
// Flush: 모든 데이터가 writer에 보내졌음을 의미 (꼭 써줘야한다!)
var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()

	var T, a, b int
	fmt.Fscan(reader, &T) // 개행문자로 읽기 종료, Fsanln은 스페이스, 탭으로 데이터를 구분.

	z := 0
	for z < T {
		fmt.Fscan(reader, &a, &b)
		msg := fmt.Sprintf("Case #%d: %d\n", z+1, a+b)
		fmt.Print(msg)
		// fmt.Printf("Case #%d: %d\n", T, a+b)

		// fmt.Printf(fmt.Sprintf("Case #%d: %d\n", T, a+b))
		z++
	}
}
