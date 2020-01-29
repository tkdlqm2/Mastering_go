package main

import(
	"fmt"
	"os"
	"strconv"
	"time"
)

var times int

// 일반적인 형태인 int 채널을 선언한 뒤에, 이를 채널에 대한 채널 변수로 전달
// 일반 int 채널로부터 데이터를 읽거나 시그널 채널인 f를 사용해 함수를 종료하는 select문 구축.
func f1 (cc chan chan int, f chan bool) {
	c := make(chan int)
	cc <- c
	defer close(c)

	sum := 0
	select {
	case x := <-c:
		for i := 0; i <= x; i ++ {
			sum = sum + i
		}
		c <- sum
	case <-f:
		return
	}
}

func main() {
	arguments := os.Args
	if len(arguments) != 2 {
		fmt.Println("Need just one integer argument!")
		return
	}

	times, err := strconv.Atoi(arguments[1])
	if err != nil {
		fmt.Println(err)
		return
	}

	cc := make(chan chan int)

	for i := 1; i < times + 1; i ++ {

		// f 채널은 시그널 채널로서 실제 작업이 끝날 때 고 루틴을 종료하는데 사용함.
		f := make(chan bool)
		go f1(cc,f)
		// 채널에 대한 채널 변수로부터 일반 채널을 받아와서 ch <- i 를 이용해 이 채널에 int 값을 보냄.
		ch := <-cc
		ch <- i

		// f1() 값 하나만 되돌려 ㅂ다도록 작성했지만, 여러 값을 읽을 수 있음.
		// 이때 각 i는 다양한 고루틴으로부터 값을 받음.
		for sum := range ch {
			fmt.Print("Sum(", i, ")=", sum)
		}
		fmt.Println()
		time.Sleep(time.Second)
		close(f)
	}
}