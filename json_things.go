package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

// Person represents a person in our JSON data
type Person struct {
	Name    string   `json:"name"`
	Age     int      `json:"age"`
	Hobbies []string `json:"hobbies"`
}

// readJSON reads JSON from a file and returns a slice of Person
func readJSON(filename string) ([]Person, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var people []Person
	if err := json.Unmarshal(data, &people); err != nil {
		return nil, err
	}
	return people, nil
}

// writeJSON writes a slice of Person to a JSON file
func writeJSON(filename string, people []Person) error {
	data, err := json.MarshalIndent(people, "", "    ")
	if err != nil {
		return err
	}
	return os.WriteFile(filename, data, 0644)
}

func main() {
	// Create some sample data
	people := []Person{
		{Name: "Alice", Age: 30, Hobbies: []string{"reading", "hiking"}},
		{Name: "Bob", Age: 25, Hobbies: []string{"gaming", "cooking"}},
	}

	// Write to JSON file
	if err := writeJSON("people.json", people); err != nil {
		log.Fatal(err)
	}

	// Read from JSON file
	loadedPeople, err := readJSON("people.json")
	if err != nil {
		log.Fatal(err)
	}

	// Do some operations with the data
	for _, person := range loadedPeople {
		fmt.Printf("Name: %s, Age: %d\n", person.Name, person.Age)
		fmt.Printf("Hobbies: %v\n", person.Hobbies)
	}
}
