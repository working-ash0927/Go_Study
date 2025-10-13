package main

import "fmt"

var pl = fmt.Println

// 예시 자료형, 메모리 점유 사이즈 1608bit = 8bit + 8 * 200 bit
type Data struct {
	value int
	data  [200]int
}

func main() {
	var a int = 118
	var b *int // 초기값 nil, nil 이 아니면 어떤 유효한 메모리 주소가 할당 되어있다는 의미

	b = &a // 포인터 b에 a 변수 메모리주소 대입
	pl(b, *b)

	*b = 18
	pl(a, *b)

	var c int = 18
	var d *int = &c
	pl(b == d)   // 사이즈는 같지만 주소가 다르기 떄문에 false
	pl(*b == *d) // 값은 같으니 true

	pl("==================")

	var data Data            // 1608bit mem var
	data2 := ChageData(data) // 1608bit mem var 가 내부적으로 생성됨
	data = data2             // data2 의 데이터를 data에 복사.
	pl(data, data2)

	ChangeDataP(&data) // 인자값으로 메모리 주소가 들어갔으니 함수 내에서 실제 값이 처리됨
	pl(data, data2)

}
func ChageData(arg Data) Data {
	arg.value = 999
	arg.data[2] = 999
	return arg
}

func ChangeDataP(arg *Data) { // 메모리 주소가 입력되면 값으로 활용함
	arg.value = 8888
	arg.data[2] = 8888
}
