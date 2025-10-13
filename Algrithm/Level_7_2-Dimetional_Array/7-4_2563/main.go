package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	l, _ := reader.ReadString('\n')
	field := strings.Fields(l)
	N, _ := strconv.Atoi(field[0])

	// var arr [100][100]int 로 동일하게 초기화 가능; bool 로도 가능;
	arr := make([][]int, 100)
	for i := range arr {
		arr[i] = make([]int, 100)
	}

	var maxx, maxy int
	for range N {
		l, _ := reader.ReadString('\n')
		field := strings.Fields(l)
		x, _ := strconv.Atoi(field[0])
		y, _ := strconv.Atoi(field[1])

		// x -= 1 // 위치가 0이면 음수가 되서 안됨
		// y -= 1

		for i := range 10 {
			for j := range 10 {
				if arr[y+i][x+j] == 0 {
					arr[y+i][x+j] = 1
				}
			}
		}
		if maxx < x+10 {
			maxx = x + 10
		}
		if maxy < y+10 {
			maxy = y + 10
		}
	}

	// 제출할때 지우기
	for i := 1; i <= maxy; i++ {
		fmt.Fprintf(writer, "%02v: ", maxy-i)
		fmt.Fprintln(writer, arr[maxy-i][:maxx]) // 보여주기에 맞게
	}

	var result int
	for i := 1; i <= maxy; i++ {
		var sum int
		for _, v := range arr[maxy-i][:maxx] {
			sum += v
		}
		result += sum
	}
	fmt.Fprintln(writer, result)

}
