package main

import (
	"fmt"
	"math/rand"
	"time"
)

func add(c chan int) {
	sum := 0
	t := time.NewTimer(time.Second)

	for {
		select {
		case input := <-c:
			sum = sum + input

		// time.NewTimer()를 호출할 때 지정한 시간만큼 차이머의 c 채널을 블록시킴.
		// 지정된 시간이 만료되면 타이머는 t.C 채널로 값을 보냄.
		// 그 후 select문에서 이와 관련된 브랜치가 실행되면서 c 채널에 nil 값을 할당한 뒤 sum 변수를 화면에 출력
		case <-t.C:
			c = nil
			fmt.Println(sum)
		}
	}
}

// 채널이 열려있는 동안 지속적으로 난수를 생성해서 채널에 보냄.
func send(c chan int) {
	for {
		c <- rand.Intn(10)
	}
}

func main() {
	c := make(chan int)
	go add(c)
	go send(c)

	// 두 개의 고루틴이 실행되는데 충분한 시간을 지정함.
	time.Sleep(3 * time.Second)
}