package main

import (
	"fmt"
)

func main() {

	// param2 : 최대 다섯 개의 정수를 저장하도록 설정한 것.
	numbers := make(chan int, 5)
	counter := 10


	// 이 부분은 numbers 채널에 10개의 정수를 대입한다.
	// 하지만 numbers 채널 용량은 5개의 정수만 가질 수 있기 때문에
	// 여기ㅓㅅ 지정한 10개의 정수를 모두 저장할 수 없다.
	for i := 0; i < counter; i++ {
		select {
		case numbers<- i:
		default:
			fmt.Println("Not enough space for", i)
		}
	}

	// for 루프와 select문을 통해 numbers 채널에 담긴 내용을 읽음
	// numbers 채널에 읽을 내용이 있다면 select문의 첫 번째 브랜치가 실행됨.
	// numbers 채널이 비어 있다면 defalut 브랜치가 실행됨.
	for i := 0; i < counter + 5; i++ {
		select {
		case num := <- numbers:
			fmt.Println(num)
		default:
			fmt.Println("Nothing more to be done!")
			break
		}
	}
}