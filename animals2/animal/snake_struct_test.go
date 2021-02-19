package animal

import "testing"

func TestSnakeStruct(t *testing.T) {

	var snakeInitializer Snake
	var name = "Bill"
	snake := snakeInitializer.Init(name)

	/*Checks if the struct was initialized correctly*/
	if snake.GetName() != name {
		t.Fail()
		t.Logf("Expected %s, got %s", name, snake.GetName())
	} else {
		/*If the test doesn't fail check  if the interface works correctly.*/
		var animal Animal

		animal = snake
		// print what the cow eats.
		animal.Eat()

		// print when the cow moves.
		animal.Move()

		// print when the cow speaks.
		animal.Speak()
	}
}
