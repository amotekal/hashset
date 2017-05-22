package hashset

import (
	"bytes"
	"encoding/gob"
	"hash/fnv"
)

//Hashset implements the extendible hashing algorithm
type Hashset struct {
	depth uint32
	pages []page
}

//New creates Hashset and returns it
func New() Hashset {
	p := page{}
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

func (eh *Hashset) getPage(key interface{}) (page, error) {
	var p page
	h, err := hash(key)
	if err != nil {
		return p, err
	}
	p = eh.pages[h&((1<<eh.depth)-1)]
	return p, err
}

//Put is e method to put a value into the set
func (eh *Hashset) Put(v interface{}) error {
	p, err := eh.getPage(v)
	if err != nil {
		return err
	}

	if p.isFull() && p.depth == eh.depth {
		eh.pages = append(eh.pages, eh.pages...)
		eh.depth++
	}
	return nil
}
