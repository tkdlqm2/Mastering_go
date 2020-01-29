package main

import (
	"fmt"
	"time"
)

// A() 는 매개변수 a에 저장된 채널에 의해 블록됨.
// main() 에서 이 채널의 블록 상태가 풀리면 A() 함수가 실행되기 시작함.
// 마지막으로 b 채널을 닫는데, 이렇게 하면 다른 함수의 블록 상태가 풀림.
func A(a,b chan struct {}) {
	<- a
	fmt.Println("A() !")
	time.Sleep(time.Second)
	close(b)
}

// 채널 a가 닫힐 때까지 블록된다. 이 채널이 닫히면 작업을 수행한 뒤 b채널을 닫는다.
// 여기서 a와 b 채널은 모두 이 함수의 매개변수 이름을 가리킴.
func B(a,b chan struct {}) {
	<- a
	fmt.Println("B()!")
	close(b)
}

// a 채널이 닫힐 때까지 블록되어 있다가 채널이 닫히면 실행을 시작.
func C(a chan struct {}) {
	<- a
	fmt.Println("C()!")
}

func main() {
	x := make(chan struct {})
	y := make(chan struct {})
	z := make(chan struct {})

	go C(z)
	go A(x,y)
	go C(z)
	go B(y,z)
	go C(z)

	close(x)
	time.Sleep(3 * time.Second)
}