package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

var Password = secret{password:"myPassword"}

// 한 개의 공뷰 변수, sync.RWMutex 타입 뮤텍스.
type secret struct {
	RwM sync.RWMutex
	M sync.Mutex
	password string
}

// 이 함수는 공유 변수를 수정함. 
// 다시 말해 이 부분에서 잠금을 설정해야 함.
// 그래서 Lock(), Unlock()함수를 사용함.
func Change(c *secret, pass string) {
	c.RwM.Lock()
	fmt.Println("LChange")
	time.Sleep(10 * time.Second)

	c.password = pass
	c.RwM.Unlock()
}

// 이 함수는 RLock(), RUnlock() 함수를 이용
// 이 함수는 임계영역에서 공유 변수를 읽기 때문임.
// 이러한 공유 변수를 여러 개의 고루틴으로 읽을 수는 있지만, Lock(), UnLock() 함수 없이는 아무 것도 변경할 수 없음.
// 여기서 Lock()함수는 뮤텍스를 이용하여 공유 변수에서 뭔가 읽는 것이 있는 한 블록됨.
func show(c *secret) string {
	c.RwM.RLock()
	fmt.Print("show")
	time.Sleep(3 * time.Second)
	defer c.RwM.RUnlock()
	return c.password
}

func showWithLock(c *secret) string {
	c.M.Lock()
	fmt.Println("showWithLock")
	time.Sleep(3 * time.Second)
	defer c.M.Unlock()
	return c.password
}

func main() {
	var showFunction = func(c *secret) string {return ""}
	if len(os.Args) != 2 {
		fmt.Println("Using sync.RwMutex!")
		showFunction = show
	} else {
		fmt.Println("Using sync.Mutex!")
		showFunction = showWithLock
	}

	var waitGroup sync.WaitGroup

	fmt.Println("Pass : ", showFunction(&Password))

	for i := 0; i < 15; i ++ {
		waitGroup.Add(1)
		go func() {
			defer waitGroup.Done()
			fmt.Println("Go Pass : ", showFunction(&Password))
		}()
	}

	go func() {
		waitGroup.Add(1)
		defer waitGroup.Done()
		Change(&Password,"123456")
	}()

	waitGroup.Wait()
	fmt.Println("Pass : ", showFunction(&Password))
}