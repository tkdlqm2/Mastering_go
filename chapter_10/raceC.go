package main

import (
	"fmt"
	"sync"
	"os"
	"strconv"
)

func main() {

	arguments := os.Args
	if len(arguments)  != 2 {
		fmt.Println("Give me a natural number!")
		os.Exit(1)
	}

	numGR, err := strconv.Atoi(os.Args[1])

	if err != nil {
		fmt.Println(err)
		return
	}

	var waitGroup sync.WaitGroup

	k := make(map[int]int)
	k[1] = 12

	for i := 0; i < numGR; i++ {
		waitGroup.Add(1)
		go func() {
			defer waitGroup.Done()
			k[i] = i
		}()
	}
	k[2] = 10
	waitGroup.Wait()
	fmt.Printf("k = %#v\n", k)
}

// 여러 go 루틴이 맵 k를 동시에 접근하는 것만으로는 부족한 것 같아 sync.Wait() 함수를
// 호출하기 전에 맵 k에 접근하는 다른 문장도 추가했다.