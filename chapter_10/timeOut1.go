package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)

	// 익명함수가 통상적으로 3초이내 채널에 메시지를 write한다는 것.
	go func() {
		time.Sleep(time.Second * 3)
		c1 <- "c1 OK"
	}()

	select {
	case res := <-c1:
		fmt.Println(res)
	
	// time.After()를 호출한 이유는 저장한 시간만큼 기다리기 위해서임. 
	// 이 함수에서 리턴하는 실제 값에는 관심 없고, 이 함수가 끝났다는 사실, 다시 말해 그 만큼 시간이 지났다는 사실만 중요	
	case <-time.After(time.Second * 1):
		fmt.Println("timeout c1")
	}

	c2 := make(chan string)

	go func() {
		time.Sleep(3 * time.Second)
		c2 <- "c2 OK"
	}()

	select {
	case res := <-c2:
		fmt.Println(res)

	// 4초 이내로 채널로부터 메시지를 수신 못하면 타임아웃.
	case <-time.After(4 * time.Second):
		fmt.Println("timeout c2")
	}
}