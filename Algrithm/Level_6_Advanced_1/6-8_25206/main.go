package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var scanner *bufio.Scanner = bufio.NewScanner(os.Stdin)
	var writer *bufio.Writer = bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var title, rating string
	var score float32

	subject := map[string]float32{
		"A+": 4.5,
		"A0": 4.0,
		"B+": 3.5,
		"B0": 3.0,
		"C+": 2.5,
		"C0": 2.0,
		"D+": 1.5,
		"D0": 1.0,
		"F":  0.0,
	}
	var result, average_count float32

	for range 20 {
		scanner.Scan()
		fmt.Sscan(scanner.Text(), &title, &score, &rating)
		if rating == "P" {
			continue
		}
		result += score * subject[rating]
		average_count += score
	}
	fmt.Fprintf(writer, "%6f", result/average_count)
}
