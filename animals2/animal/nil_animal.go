package animal

type NilAnimal struct {
	name string
}

func (nil NilAnimal) Init(name string) Animal {
	nil.name = "Nil"
	return nil
}

func (nil NilAnimal) GetName() string {
	return nil.name
}

func (nil NilAnimal) Eat() {
}

func (nil NilAnimal) Speak() {
}

func (nil NilAnimal) Move() {
}
