package animal

type Animal interface {
	Roar()
	Walk()
}

func Roar(animal Animal) {
	animal.Roar()
}
