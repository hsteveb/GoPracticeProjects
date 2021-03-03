package person

import "../utensils"
import "sync"

/*Person inteface*/
type Person interface {
	Init(int, *utensils.Utensil, *utensils.Utensil) Person
	GetName() int
	GetEat() int
	Eat(*sync.WaitGroup)
}
