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

	scanner.Scan()
	data := scanner.Text()

	var result int
	for i := range N {
		n, _ := strconv.Atoi(string(data[i]))
		result += n
	}

	fmt.Println(result)

}
