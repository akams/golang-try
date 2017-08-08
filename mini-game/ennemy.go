package main

import (
	"fmt"
	"io/ioutil"
)

func readFile() {
	bs, err := ioutil.ReadFile("./files/ennemy.txt")
	if err != nil {
		return
	}
	str := string(bs)
	fmt.Println(str)
}

// OP 2 25 5 9 9 3
// Lantern 5 15 15 19 19 13
