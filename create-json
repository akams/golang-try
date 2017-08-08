package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Circle struct {
	X, Y, R float64
}

func main() {
	fmt.Println("Hello, playground")
  
	var arrCircle []Circle

	for i := 0; i < 5; i++ {
		arrCircle = append(arrCircle, Circle{X: float64(i), Y: float64(i), R: float64(i)})
	}

	fmt.Printf("len=%d cap=%d %v\n", len(arrCircle), cap(arrCircle), arrCircle)

	b, error := json.Marshal(arrCircle)
	if error != nil {
		fmt.Println(error)
		return
	}
	fmt.Println(string(b))

	file, err := os.Create("test.json")
	if err != nil {
		// handle the error here
		return
	}
	defer file.Close()
	file.WriteString(string(b))
}
