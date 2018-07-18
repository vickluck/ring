package balance

type RoundRobin struct {
	values []interface{}
	cursor int // == index

	// Save the index has been deleted
	delIndex []int
}

func (rr *RoundRobin) add(v interface{}) {
	rr.values = append(rr.values, v)
}

func (rr *RoundRobin) get() interface{} {
	p := rr.values[rr.cursor]
	rr.cursor = rr.cursor + 1

	if rr.cursor == len(rr.values) {
		rr.cursor = 0
	}
	return p
}

// Len
func (rr *RoundRobin) len() int {
	return len(rr.values)
}
