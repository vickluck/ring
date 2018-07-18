package balance

type balancer interface {
	add(interface{})
	get() interface{}
	len() int
}

type Mode uint16

const (
	MODE_RR Mode = iota // 0
)

func (m Mode) pattern() balancer {
	switch m {
	case MODE_RR:
		return &roundRobin{}
	default:
		return nil
	}
}

type balances map[string]balancer

type Struct struct {
	balances
}

func New() *Struct {
	return &Struct{}
}

func (b balances) SetMode(key string, m Mode) {
	b[key] = m.pattern()
}

func (b balances) Put(key string, value interface{}) {
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

func (b balances) Get(key string) interface{} {
	return b[key].get()
}

func (b balances) Len(key string) int {
	return b[key].len()
}
