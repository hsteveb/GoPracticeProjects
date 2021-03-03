package main

//import "fmt"
import "./utensils"
import "./person"
import "sync"

func main() {
	var wg sync.WaitGroup
	var chopsticks = make([]utensils.Utensil, 0)
	var philosophers = make([]person.Person, 0)

	for i := 0; i < 5; i++ {
		var chopstick utensils.Chopstick
		chopsticks = append(chopsticks, chopstick.Init())
	}

	for i := 0; i < 5; i++ {
		var philosopher person.Philosopher
		philosophers = append(philosophers, philosopher.Init(i+1, &chopsticks[i], &chopsticks[(i+1)%5]))
	}

	var nextPhilosopher = 0

	for !everyoneEaten(&philosophers) {
		for i := 0; i < 2; i++ {
			wg.Add(1)
			go philosophers[nextPhilosopher].Eat(&wg)
			nextPhilosopher = (nextPhilosopher + 1) % 5
		}
		wg.Wait()
	}
}

func everyoneEaten(philosophers *[]person.Person) bool {
	for i := 0; i < len(*philosophers); i++ {
		if (*philosophers)[i].GetEat() < 3 {
			return false
		}
	}
	return true
}
