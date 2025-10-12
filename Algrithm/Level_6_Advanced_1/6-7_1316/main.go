package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

// map 을 써서 썻던거에 대한 플래그로 false 처리하면 될 수 있음
func main() {
	var scanner *bufio.Scanner = bufio.NewScanner(os.Stdin)
	var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var N int
	scanner.Scan()

	fmt.Sscan(scanner.Text(), &N)

	var result int
	for range N {
		scanner.Scan()
		d := scanner.Text()

		var barr []byte
		cnt := true
		barr = append(barr, d[0]) // 한번이라도 사용됐던 데이터 저장
		for i := 1; i < len(d); i++ {
			if !slices.Contains(barr, d[i]) { // contain이 병목 가능성이 있어 map 을 권장
				barr = append(barr, d[i])
			} else {
				if d[i-1] != d[i] {
					cnt = false
					break
				}
			}
		}

		if cnt {
			result += 1
		}

	}
	fmt.Fprintln(writer, result)

}
