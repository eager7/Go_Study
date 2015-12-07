package main

import (
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	type ColorGroup struct {
		ID     int
		Name   int
		Colors []string
	}
	group := ColorGroup{
		ID:     0,
		Name:   2,
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
	}
	b, err := json.Marshal(group)
	if err != nil {
		fmt.Println("error:", err)
	}
	os.Stdout.Write(b)
	println("")
}