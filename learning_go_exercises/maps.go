package main

import (
	"fmt"
	"sort"
)

func main() {
	// A map is an unordered collection of key: value pairs--dictionary in python
	states := make(map[string]string)
	fmt.Println("Maps", states)
	states["CA"] = "California"
	states["IL"] = "Illinois"
	states["OR"] = "Oregon"

	california := states["CA"]

	fmt.Println(states)
	fmt.Println(california)

	// Delete an item
	delete(states, "OR")
	fmt.Println(states)

	states["NY"] = "New York"

	// Loop through a map
	for k, v := range states {
		fmt.Println("People in", v, "with code", k, "\b.")
		fmt.Printf("I live in %v with code %v.\n", v, k)
	}

	// Sort a map procedure
	keys := make([]string, len(states))
	i := 0
	for k := range states {
		keys[i] = k
		i++
	}
	fmt.Println(keys)
	sort.Strings(keys)
	fmt.Println(keys)
	// Final step to sort the map, at least in print.
	for i := range keys {
		fmt.Println(states[keys[i]])
	}
}
