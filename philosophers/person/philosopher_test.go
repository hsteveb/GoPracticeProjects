package person

import "testing"
import "../utensils"
import "sync"

var philosopherInit Philosopher

func TestPhilosopherInit(t *testing.T) {
	var lChopstick utensils.Chopstick
	var lutensil utensils.Utensil
	lutensil = lChopstick.Init()
	var rChopstick utensils.Chopstick

	var rutensil utensils.Utensil
	rutensil = rChopstick.Init()

	var philosopher = philosopherInit.Init(1, &lutensil, &rutensil)

	if philosopher.GetName() != 1 {
		t.Logf("Got %d expected %d", philosopher.GetName(), 1)
		t.Fail()
	}
}

func TestGetEat(t *testing.T) {
	var lChopstick utensils.Chopstick
	var lutensil utensils.Utensil
	lutensil = lChopstick.Init()
	var rChopstick utensils.Chopstick

	var rutensil utensils.Utensil
	rutensil = rChopstick.Init()
	var philosopher = philosopherInit.Init(1, &lutensil, &rutensil)

	if philosopher.GetEat() != 0 {
		t.Logf("Got %d expected %d", philosopher.GetEat(), 0)
		t.Fail()
	}
}

func TestStartEating(t *testing.T) {
	var wg sync.WaitGroup
	var lChopstick utensils.Chopstick
	var lutensil utensils.Utensil
	lutensil = lChopstick.Init()
	var rChopstick utensils.Chopstick

	var rutensil utensils.Utensil
	rutensil = rChopstick.Init()
	var philosopher = philosopherInit.Init(1, &lutensil, &rutensil)

	wg.Add(1)
	go philosopher.Eat(&wg)
	wg.Wait()

	if philosopher.GetEat() != 1 {
		t.Logf("Got %d expected %d", philosopher.GetEat(), 1)
		t.Fail()
	}
}
