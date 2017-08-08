package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
)

type Person struct {
	Name    string
	Age     int
	Details interface{}
}

func main() {
	file, e := ioutil.ReadFile("./config.json")
	if e != nil {
		fmt.Println("File error:", e)
		os.Exit(1)
	}
	fmt.Println(string(file))
  	keysBody := []byte(string(file))
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
