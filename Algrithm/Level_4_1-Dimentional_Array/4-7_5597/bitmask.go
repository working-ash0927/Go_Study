package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func using_bitmask() {
	// test with copilot
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 학생 번호는 1~30
	const N = 30
	var mask int

	// 입력: 28개 번호
	for i := 0; i < 28; i++ {
		scanner.Scan()
		var v int
		fmt.Sscan(scanner.Text(), &v)
		mask |= 1 << v
	}

	// 전체 집합 (1~30)
	var fullmask int
	for i := 1; i <= N; i++ {
		fullmask |= 1 << i
	}

	// 빠진 번호 찾기
	missing := fullmask &^ mask

	var cnt []int
	for i := 1; i <= N; i++ {
		if missing&(1<<i) != 0 {
			cnt = append(cnt, i)
		}
	}

	slices.Sort(cnt)
	for _, v := range cnt {
		fmt.Fprintln(writer, v)
	}
}
