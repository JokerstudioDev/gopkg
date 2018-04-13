package main

import (
	"fmt"

	"github.com/JokerstudioDev/gopkg/dac"
)

func main() {
	// counter := calculator.NewCounter(0)
	// counter.Increase()
	// counter.Increase()
	// counter.Increase()
	// counter.Increase()
	// fmt.Printf("%d", counter.Count)
	result := dac.GetAccountInformation(dac.RealAccount{})
	fmt.Printf(result)
}
