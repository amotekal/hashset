package hashset

const maxPageSize int = 200

type page struct {
	depth int
	items map[int]interface{}
}

func (p *page) isFull() bool {
	return len(p.items) >= maxPageSize
}

func (p *page) put(key int, value interface{}) {
	p.items[key] = value
}

func (p *page) get(key int) interface{} {
	return p.items[key]
}
