package utensils

import "sync"

/*Chopstick struct */
type Chopstick struct {
	wg sync.Mutex
}

/*Init initialization*/
func (chopstick *Chopstick) Init() Utensil {
	return chopstick
}

/*Lock lock mutex*/
func (chopstick *Chopstick) Lock() {
	chopstick.wg.Lock()
}

/*Unlock unlock mutext*/
func (chopstick *Chopstick) Unlock() {
	chopstick.wg.Unlock()
}
