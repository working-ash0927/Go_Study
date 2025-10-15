package main

import (
	"fmt"
	"strings"
	"unsafe"
)

func main() {
	t1 := "ABC"
	fmt.Println(t1[0], len(t1))

	t2 := "ABC얍"
	fmt.Println(t2[3], len(t2))

	s1 := "Hi" // 문자열 선언
	s2 := s1   // s2에 s1을 대입

	fmt.Println((unsafe.StringData(s1)), (unsafe.Pointer(&s1)))
	fmt.Println((unsafe.StringData(s2)), (unsafe.Pointer(&s2)))

	var s string
	for i := range 26 {
		s += string('A' + i)
	}
	fmt.Println(s)

	var builder strings.Builder
	for i := range 26 {
		builder.WriteRune(rune('A' + i))
	}
	fmt.Println(builder.String())

	s3 := "ZXC가나다"
	for i, v := range s3 {
		fmt.Println(i, v)
	}

	a := [3]int{1, 2, 3}
	b := [3]int{4, 5, 6}

	a = b

	fmt.Println(a, b)
}
