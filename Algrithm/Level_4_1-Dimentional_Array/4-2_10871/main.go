package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var scanner *bufio.Scanner = bufio.NewScanner(os.Stdin)
	var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 입력받은 띄어쓰기 기반 데이터를 각 변수에 바로 할당하는 방법
	N, X := func() (int, int) {
		scanner.Scan()
		f := strings.Fields(scanner.Text())
		r1, _ := strconv.Atoi(f[0])
		r2, _ := strconv.Atoi(f[1])
		return r1, r2
	}()

	scanner.Scan()
	temp := strings.Fields(scanner.Text())
	if N != len(temp) {
		return
	}
	var result []string
	for _, v := range temp {
		a, _ := strconv.Atoi(v)
		if a < X {
			result = append(result, v)
		}
	}
	fmt.Println(strings.Join(result, " "))
}
