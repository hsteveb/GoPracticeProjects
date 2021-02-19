package animal

type Animal interface {
	Init(name string) Animal
	Eat()
	Speak()
	Move()
	GetName() string
}
