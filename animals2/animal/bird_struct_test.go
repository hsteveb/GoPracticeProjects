package animal

import "testing"

//import "fmt"

func TestBirdStruct(t *testing.T) {

	var bird2 Bird
	var name = "Bill"
	bird := bird2.Init(name)

	/*Checks if the struct was initialized correctly*/
	if bird.GetName() != name {
		t.Fail()
		t.Logf("Expected %s, got %s", name, bird.GetName())
	} else {
		/*If the test doesn't fail check  if the interface works correctly.*/
		var animal Animal

		animal = bird
		// print what the cow eats.
		animal.Eat()

		// print when the cow moves.
		animal.Move()

		// print when the cow speaks.
		animal.Speak()
	}
}
