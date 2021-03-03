package utensils

/*Utensil interface*/
type Utensil interface {
	Init() Utensil
	Lock()
	Unlock()
}
