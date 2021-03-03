package person

import "../utensils"
import "fmt"
import "sync"

/*Philosopher struct*/
type Philosopher struct {
	name       int
	eat        int
	lChopstick *utensils.Utensil
	rChopstick *utensils.Utensil
}

/*Init philosopher*/
func (philosopher *Philosopher) Init(name int, lChopstick *utensils.Utensil, rChopStick *utensils.Utensil) Person {
	philosopher.name = name
	philosopher.eat = 0
	philosopher.lChopstick = lChopstick
	philosopher.rChopstick = rChopStick

	return philosopher
}

/*GetEat gets how many times the philosopher eats*/
func (philosopher Philosopher) GetEat() int {
	return philosopher.eat
}

/*GetName get name of philosopher*/
func (philosopher Philosopher) GetName() int {
	return philosopher.name
}

/*Eat philosopher eats*/
func (philosopher *Philosopher) Eat(wg *sync.WaitGroup) {
	if philosopher.eat < 3 {
		var lChopstick utensils.Utensil
		var rChopstick utensils.Utensil

		lChopstick = *philosopher.lChopstick
		lChopstick.Lock()
		rChopstick = *philosopher.rChopstick
		rChopstick.Lock()
		philosopher.eat++
		fmt.Printf("Starting to eat %d\n", philosopher.name)
		fmt.Printf("Finishing eating %d\n", philosopher.name)
		lChopstick.Unlock()
		rChopstick.Unlock()
	}
	wg.Done()
}
