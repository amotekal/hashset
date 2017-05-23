package hashset

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"hash/fnv"
)

//Hashset implements the extendible hashing algorithm
type Hashset struct {
	depth uint32
	pages []page
}

//NewHashset creates Hashset and returns it
func NewHashset() Hashset {
	p := newPage()
	pp := []page{p}
	eh := Hashset{depth: 0, pages: pp}

	return eh
}

func hash(key interface{}) (uint32, error) {
	//convert interface to []byte
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(key)
	if err != nil {
		return 0, err
	}

	h := fnv.New32a()
	h.Write(buf.Bytes())
	return h.Sum32(), nil
}

func (eh *Hashset) getPage(key interface{}) (page, uint32, error) {
	var p page
	h, err := hash(key)
	if err != nil {
		return p, 0, err
	}
	p = eh.pages[h&((1<<eh.depth)-1)]
	return p, h, err
}

//Put is e method to put a value into the set
func (eh *Hashset) Put(v interface{}) error {
	p, k, err := eh.getPage(v)
	if err != nil {
		return err
	}

	if p.isFull() && p.depth == eh.depth {
		fmt.Printf("before len: %d cap: %d", len(eh.pages), cap(eh.pages))
		eh.pages = append(eh.pages, eh.pages...)
		fmt.Printf("after len: %d cap: %d", len(eh.pages), cap(eh.pages))
		eh.depth++
	}

	if p.isFull() && p.depth < eh.depth {
		p.put(k, v)
		p1 := newPage()
		p2 := newPage()

		for k2, v2 := range p.items {
			h, _ := hash(k2)
			h = h & ((1 << eh.depth) - 1)
			if (h>>p.depth)&1 == 1 {
				p2.put(k2, v2)
			} else {
				p1.put(k2, v2)
			}
		}

		for i, x := range eh.pages {
			if &x == &p {
				if i>>p.depth&1 == 1 {
					eh.pages[i] = p2
				} else {
					eh.pages[i] = p1
				}
			}
		}

		p1.depth = p.depth + 1
		p2.depth = p1.depth
	} else {
		p.put(k, v)
	}

	return nil
}

//PrintAll prints all the contents of the hashset
func (eh *Hashset) PrintAll() {
	for _, page := range eh.pages {
		for _, v := range page.items {
			fmt.Println(v)
		}
	}
}
