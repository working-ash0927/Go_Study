** 백준 문제 풀이 **
- 활용해야 할 필수 모듈
"bufio" : 표준 입출력 
"fmt" : 콘솔 출력
"os" : 표준 입출력을 위함

- 문법 
1. defer func () 는 상위 함수가 중료되기 전에 동작하도록 하는 예약선언. 로직 마지막에 동작하도록 선언하려는 경우 활용
2. 쿼터 (`) 는 다중 문자열을 출력하는 경우 활용
3. strconv.Itoa() : 정수형을 Ascii 로 변환. 반대는 Atoi()
4. fmt.Sprintf 는 Printf 와 달리 출력 없이 문자열로 반환되는 함수. (string print formatted)
5. var a []int 는 고정길이 배열이므로, 다른 변수의 길이로 선언하여 활용할 수 없음.
6. 이를 위해선 슬라이스 (tmp := make([]int, n)) 형태로 동적 배열을 활용해야함
7. 초기값 0이 출력되는거 방지
    `if _, err := fmt.Fscan(reader, &a, &b); err != nil {
        continue
    }`
8. 언어별 빠른 I/O 처리 방식 (문제 백준 15552) : 
    https://docs.google.com/document/d/17OUl9nU9i7vTkhk2q_qy4Q5Vl0HHE9bTLUHwbLp56WM/edit?pli=1&tab=t.0#heading=h.oq3q0ypfd1l2
9. Fprintln 는 bufio.NewWriter(os.Stdout) 으로 선언해 할당된 변수 값이 실제 출력되는 타이밍.
    단순 콘솔 출력 목적인 println 보다 더 적합하게 활용할 수 있음

10. 패키지를 아래 처럼 단순하게 사용할 수 있음
    `var pl = fmt.Println
    pl("Hello World!")`
11. 문자열 반복 생성
    `fmt.Println(strings.Repeat("*", 5)) // *****`