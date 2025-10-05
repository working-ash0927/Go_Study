package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var pl = fmt.Println

// defer: 이를 호출한 함수의 리턴 직전에 실행.
// Flush: 모든 데이터가 writer에 보내졌음을 의미 (꼭 써줘야한다!)

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var a, b, c int
	fmt.Fscanln(reader, &a, &b, &c) // 개행문자로 읽기 종료, Fsanln은 스페이스, 탭으로 데이터를 구분.
	arr := [...]int{a, b, c}

	if a == b && b == c {
		pl(10000 + (a * 1000))
	} else {
		var same_val int
		var cnt int

	outer:
		for i := 0; i < len(arr); i++ {
			cnt = 1
			for j := 0; j < len(arr); j++ {
				if i == j {
					continue
				}
				if arr[i] == arr[j] {
					same_val = arr[j]
					cnt++
					break outer
				}
			}
		}
		if cnt == 1 {
			maxVal := int(math.Max(math.Max(float64(a), float64(b)), float64(c)))
			pl(maxVal * 100)
		} else {
			pl(1000 + (same_val * 100))
		}
	}
}
