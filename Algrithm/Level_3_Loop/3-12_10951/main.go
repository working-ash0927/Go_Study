package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// defer: 이를 호출한 함수의 리턴 직전에 실행한다.
// Flush: 모든 데이터가 writer에 보내졌음을 의미 (꼭 써줘야한다!)
var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func main() {
	defer writer.Flush()
	var a, b int

	for {
		line, _ := reader.ReadString('\n')
		line = strings.TrimSpace(line)
		// 엔터만 입력 → 종료
		if line == "" {
			break
		}
		// Sscan 은 io.reader 대신 문자열 데이터를 처리
		fmt.Sscan(line, &a, &b) // 공백 구분해서 값 할당 (줄바꿈도 공백취급)
		// if err != nil { // 잘못된 값 들어오면
		// 	break
		// }
		if (a < 0 && a > 10) || (b < 0 && b > 10) {
			break
		}
		fmt.Println(a + b)
	}
}
