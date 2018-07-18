package balance

type balancer interface {
	add(interface{})
	get() interface{}
	len() int
}

type mode uint16

const (
	MODE_RR mode = iota
)

func (m mode) pattern() balancer {
	switch m {
	case MODE_RR:
		return &RoundRobin{}
	default:
		return nil
	}
}

type balances map[string]balancer

func New() balances {
	return make(balances)
}

func (b balances) SetMode(key string, m mode) {
	b[key] = m.pattern()
}

func (b balances) Push(key string, value interface{}) {
	if v, ok := b[key]; ok {
		v.add(value)
	} else {
		// If Mode is not set, it defaults to the MODE_RR
		if b[key] == nil {
			b.SetMode(key, MODE_RR)
		}
		b[key].add(value)
	}
}

func (b balances) Pull(key string) interface{} {
	return b[key].get()
}

func (b balances) Len(key string) int {
	return b[key].len()
}
