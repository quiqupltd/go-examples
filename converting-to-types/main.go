package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type (
	Person struct {
		Name string `json:"name"`
		Age  int    `json:"age"`
	}

	Animal struct {
		Breed string `json:"breed"`
		Age   int    `json:"age"`
		Legs  int    `json:"legs"`
	}
)

var (
	mappings = map[string]interface{}{
		"person": Person{},
		"animal": Animal{},
	}
)

func CreateInstance(key string) (interface{}, error) {
	// Find the example instance in the map.
	exampleInstance, exists := mappings[key]
	if !exists {
		return nil, fmt.Errorf("no type found for key %s", key)
	}

	// Use reflection to get the type of the instance, then create a new instance of that type.
	instanceType := reflect.TypeOf(exampleInstance)
	newInstancePtr := reflect.New(instanceType)

	return newInstancePtr.Interface(), nil
}

func main() {
	// Define a person
	person := &Person{
		Name: "Danny",
		Age:  43,
	}

	// Marshal to json
	jsonData, err := json.Marshal(person)
	if err != nil {
		panic(err)
	}
	fmt.Println("JSON data:")
	fmt.Println(string(jsonData))

	// Create a new pointer to a person instance and unmarshal json into the pointer
	personInterface, err := CreateInstance("person")
	if err != nil {
		panic(err)
	}
	json.Unmarshal(jsonData, &personInterface)
	fmt.Println("Person Interface:")
	fmt.Println(personInterface)

	// Cast the interface to a person
	if newPerson, ok := personInterface.(*Person); ok {
		fmt.Println(newPerson.Name)
		fmt.Println(newPerson.Age)
	} else {
		fmt.Print("Error: personInterface is not a Person")
	}

	// Expect to error
	if testAnimal, ok := personInterface.(*Animal); ok {
		fmt.Println(testAnimal.Breed)
	} else {
		fmt.Print("Error: personInterface is not an Animal")
	}

}
