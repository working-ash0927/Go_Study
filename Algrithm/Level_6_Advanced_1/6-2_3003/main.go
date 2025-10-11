package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var reader *bufio.Reader = bufio.NewReader(os.Stdin)
	var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	l, err := reader.ReadString('\n')
	l = strings.TrimSpace(l)
	if err != nil || l == "" {
		return
	}

	refer := [...]string{"1", "1", "2", "2", "2", "8"}
	input := strings.Fields(l)
	for i, v := range input {
		iv, _ := strconv.Atoi(v)
		rv, _ := strconv.Atoi(refer[i])
		refer[i] = strconv.Itoa(rv - iv)
	}
	// [:] 선언하면 슬라이스로 반환
	fmt.Fprintln(writer, strings.Join(refer[:], " "))

}
