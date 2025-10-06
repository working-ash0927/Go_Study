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

	var N, M int
	scanner.Scan()
	line := scanner.Text()
	fmt.Sscan(line, &N, &M)

	arr := make([]int, N)
	result := make([]string, N)
	for i := range arr {
		arr[i] = i + 1
	}

	for range M {
		var i, j, tmp int
		scanner.Scan()
		line := scanner.Text()
		fmt.Sscan(line, &i, &j)
		if i == j {
			continue
		} else {
			t := (j-i)/2 + 1 // 배열 앞 뒤로 체크할거니 절반잘라서 처리
			for range t {
				tmp = arr[i-1]
				arr[i-1] = arr[j-1]
				arr[j-1] = tmp

				j -= 1
				i += 1
			}
		}
	}

	for i, v := range arr {
		result[i] = strconv.Itoa(v)
	}
	fmt.Print(strings.Join(result, " "))
}

func tresh() {
	// 배열 내 값을 기준해서 뒤집기
	var scanner *bufio.Scanner = bufio.NewScanner(os.Stdin)
	var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var temp [42]int
	for i := range temp { // initialize
		temp[i] = 99
	}

	var N, M int
	scanner.Scan()
	line := scanner.Text()
	fmt.Sscan(line, &N, &M)

	arr := make([]int, N)
	result := make([]string, N)
	for i := range arr {
		arr[i] = i + 1
	}
	fmt.Println(arr)
	for range M {
		var i, j, tmp int
		scanner.Scan()
		line := scanner.Text()
		fmt.Sscan(line, &i, &j)
		if i == j {
			continue
		} else {
			var start, end int
			for {
				if arr[start] == i {
					end = start
					break
				}
				start += 1
			}
			for {
				if arr[end] == j {
					break
				}
				end += 1
			}

			t := (end-start)/2 + 1
			for range t {
				tmp = arr[start]
				arr[start] = arr[end]
				arr[end] = tmp

				end -= 1
				start += 1
			}
		}
		fmt.Println(arr)
	}

	for i, v := range arr {
		result[i] = strconv.Itoa(v)
	}
	fmt.Print(strings.Join(result, " "))
}
