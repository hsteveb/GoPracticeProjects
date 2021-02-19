package animal

import "testing"

//import "fmt"

func TestCowStruct(t *testing.T) {

	var cowInitializer Cow
	var name = "Bill"
	cow := cowInitializer.Init(name)

	/*Checks if the struct was initialized correctly*/
	if cow.GetName() != name {
		t.Fail()
		t.Logf("Expected %s, got %s", name, cow.GetName())
	} else {
		/*If the test doesn't fail check  if the interface works correctly.*/
		var animal Animal

		animal = cow
		// print what the cow eats.
		animal.Eat()

		// print when the cow moves.
		animal.Move()

		// print when the cow speaks.
		animal.Speak()
	}

}
