package dices

import "math/rand"

// Dices can be thrown
type Dices struct {
	First, Second int
}

// Throw dices
func (d *Dices) Throw() {
	d.First, d.Second = rand.Intn(6)+1, rand.Intn(6)+1
}

// Result after thrown
func (d Dices) Result() int {
	return d.First + d.Second
}
