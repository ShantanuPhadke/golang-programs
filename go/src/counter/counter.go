package counter

import "sync"

/*Counter ...
value = Current integer value of the Counter
*/
type Counter struct {
	value int
	mu    sync.Mutex
}

/*NewCounter ...
INPUTS: None
OUTPUTS:
	Returns a pointer to a brand new Counter object
*/
func NewCounter() *Counter {
	return &Counter{}
}

/*Inc ...
INPUTS:
	counter = Pointer to a Counter object
OUTPUTS:
	Doesn't return anything
*/
func (counter *Counter) Inc() {
	counter.mu.Lock()
	defer counter.mu.Unlock()
	counter.value++
}

/*Value ...
INPUTS:
	counter = Pointer to a Counter object
OUTPUTS:
	Returns the current value of the Counter object
*/
func (counter *Counter) Value() int {
	return counter.value
}
