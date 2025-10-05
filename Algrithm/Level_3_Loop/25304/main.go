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
func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var x, n, a, b int
	fmt.Fscan(reader, &x, &n) // 개행문자로 읽기 종료, Fsanln은 스페이스, 탭으로 데이터를 구분.

	if (1 <= x && x <= 1000000000) && (1 <= n && n <= 100) {
		var result int
		i := 0
		for i < n+1 {
			fmt.Fscanln(reader, &a, &b)
			if (1 <= a && a <= 1000000) && (1 <= b && b <= 10) {
				result += (a * b)
			}
			i++
		}
		if x == result {
			pl("Yes")
		} else {
			pl("No")
		}
	} else {
		return
	}
}
