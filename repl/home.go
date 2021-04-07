package repl

import (
	"fmt"
	"math/rand"
)

func printRepl(){
	fmt.Print("♦" , currentPrice())
}

func currentPrice() float32{
	return rand.Float32() * 100
}
