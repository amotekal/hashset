package main

import (
	"fmt"
	"hashset/hashset"
)

type as struct {
	I int
}

func main() {
	p := hashset.NewHashset()

	for i := 0; i < 10; i++ {
		err := p.Put(as{i})
		if err != nil {
			fmt.Println(err)
		}
	}

	p.PrintAll()

}
