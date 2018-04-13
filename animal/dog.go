package animal

import (
	"fmt"
)

type Dog struct{}

func (Dog) Roar() {
	fmt.Println("hong hong")
}

func (Dog) Walk() {
	fmt.Println("hong hong")
}
