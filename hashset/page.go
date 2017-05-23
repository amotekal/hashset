package hashset

const maxPageSize int = 2

type page struct {
	depth uint32
	items map[uint32]interface{}
}

func newPage() page {
	i := make(map[uint32]interface{})
	return page{depth: 0, items: i}
}

func (p *page) isFull() bool {
	return len(p.items) >= maxPageSize
}

func (p *page) put(key uint32, value interface{}) {
	p.items[key] = value
}

func (p *page) get(key uint32) interface{} {
	return p.items[key]
}
