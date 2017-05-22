package hashset

import (
	"bytes"
	"encoding/gob"
	"hash/fnv"
)

//Hashset implements the extendible hashing algorithm
type Hashset struct {
	depth int
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

func (eh *Hashset) getPage(key int) {

}
