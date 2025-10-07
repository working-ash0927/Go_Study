package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var scanner *bufio.Scanner = bufio.NewScanner(os.Stdin)
	var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	scanner.Scan()
	N, _ := strconv.Atoi(scanner.Text())

	for range N {
		scanner.Scan()
		l := scanner.Text() // 개행 포함 안되있으니

		switch len(l) {
		case 1:
			fmt.Print(l, l+"\n")
		case 2:
			fmt.Print(string(l[0]), string(l[1])+"\n")
		default:
			fmt.Print(string(l[0]), string(l[len(l)-1])+"\n")
		}
	}
}
