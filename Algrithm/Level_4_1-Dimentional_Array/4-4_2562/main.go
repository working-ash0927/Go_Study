package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

// 64kb 를 사용하는 NewScanner 는 기본값으로 입력 데이터를 64kb (8byte) 까지만 받을 수 있음.
// 본 문제는 두번째 입력 데이터가 공백을 포함하면 10MB 까지 늘어날 수 있기에 다른 방안 체크 필요
func main() {
	using_scanner()
	// using_reader()
}

func using_reader() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	l, _ := reader.ReadString('\n')
	l = strings.TrimSpace(l) // 개행 제거
	N, _ := strconv.Atoi(l)

	r := make([]int, N)

	l2, _ := reader.ReadString('\n')
	// l2 = strings.TrimSpace(l) // 개행 제거
	data := strings.Fields(l2)
	for i, v := range data {
		n, _ := strconv.Atoi(v)
		r[i] = n
	}

	fmt.Println(slices.Min(r), slices.Max(r))
}
func using_scanner() {
	var scanner *bufio.Scanner = bufio.NewScanner(os.Stdin)
	var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	// 버퍼 사이즈 증가
	const maxCap = 10 * 1024 * 1024
	buf := make([]byte, 0, 64*1024)
	scanner.Buffer(buf, maxCap)

	// 입력받은 띄어쓰기 기반 데이터를 각 변수에 바로 할당하는 방법
	scanner.Scan()
	N, _ := strconv.Atoi(scanner.Text())

	scanner.Scan()
	temp := strings.Fields(scanner.Text())

	result := make([]int, N)
	for i, v := range temp {
		a, _ := strconv.Atoi(v)
		result[i] = a

	}
	fmt.Println(slices.Min(result), slices.Max(result))
}
