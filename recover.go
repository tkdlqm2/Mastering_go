package main

import (
	"fmt"
	"os"
)

func main() {

	if len(os.Args) == 1 {
		panic("Not enough argruments !")
	}

	fmt.Println("Thanks for the arguments(S)!")
}
