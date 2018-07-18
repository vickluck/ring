package balance

type roundRobin struct {
	values []interface{}
	cursor int // == index

	// Save the index has been deleted
	delIndex []int
}

func (rr *roundRobin) add(v interface{}) {
	rr.values = append(rr.values, v)
}

func (rr *roundRobin) get() interface{} {
	p := rr.values[rr.cursor]
	rr.cursor = rr.cursor + 1

	if rr.cursor == len(rr.values) {
		rr.cursor = 0
	}
	return p
}

// Len
func (rr *roundRobin) len() int {
	return len(rr.values)
}
