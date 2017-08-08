package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type Person struct {
	Name    string
	Age     int
	Details interface{}
}

func main() {
  	keysBody := []byte(`[{"name": "sam", "age": 27, "details": {"salary":10000} }, {"name": "john", "age": 31, "details": {"salary":12000} }]`)
	var data []Person
 
	/*json.Unmarshal is used when the input is []byte*/
	/*stores the decoded result in map named data*/
	err := json.Unmarshal(keysBody, &data)

	if err != nil {
		panic(err)
	}
  
	fmt.Printf("==data==%#v \n\n", data)
	fmt.Println("==name== ", data[0].Name)
	fmt.Println("==name== ", data[1].Name)

	/*Type assertions, check type of data["details"]*/
	details, ok := data[0].Details.(map[string]interface{})
	if ok {
		fmt.Println("==salary== ", details["salary"])
	}
}
