package animal

import (
	"fmt"
)

type Cat struct{}

func (Cat) Roar() {
	fmt.Println("meaw meaw")
}

func (Cat) Walk() {
	fmt.Println("hong hong")
}
