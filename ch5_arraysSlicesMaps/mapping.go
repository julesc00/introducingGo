package main

import "fmt"

/*
A map is an unordered collection of key-value pairs (maps are also sometimes called associative arrays, hash tables,
or dictionaries in Python). Maps are used to look up a value by its associated key.
*/
func main() {
	map1 := make(map[string]string)
	map1["name"] = "Claudia"
	map1["location"] = "Mexico"

	map2 := make(map[string]int)
	map2["age"] = 46

	map3 := make(map[int]int)
	map3[0] = 10

	map4 := make(map[int]float64)
	map4[2] = 5.20

	fmt.Println(map1, map2, map3, map4)

	// Way 1 to create a map of elements
	elements := make(map[string]string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"
	elements["Be"] = "Beryllium"
	elements["B"] = "Boron"
	elements["C"] = "Carbon"
	elements["N"] = "Nitrogen"
	elements["O"] = "Oxygen"
	elements["F"] = "Fluorine"
	elements["Ne"] = "Neon"

	fmt.Println(elements["Ne"])

	name, ok := elements["B"]
	fmt.Println(name, ok) // if not exists returns ' false' else for example 'Oxygen true'

	// Another way to check if key:value pair exists, if item doesn't exist return an empty string
	if name, ok := elements["F"]; ok {
		fmt.Println(name, ok)
	}

	shorterWay()
	mapTwo()
}

/*
	Shorter way 2 to create a map of elements
*/
func shorterWay() {
	elements := map[string]string{
		"H":  "Hydrogen",
		"He": "Helium",
		"Li": "Lithium",
		"Be": "Beryllium",
		"B":  "Boron",
		"C":  "Carbon",
		"N":  "Nitrogen",
		"O":  "Oxygen",
		"F":  "Fluorine",
		"Ne": "Neon",
	}
	fmt.Println(elements)
}

func mapTwo() {
	elements := map[string]map[string]string{
		"H": map[string]string{
			"name":  "Helium",
			"state": "gas",
		},
		"He": map[string]string{
			"name":  "Helium",
			"state": "gas",
		},
		"Li": map[string]string{
			"name":  "Lithium",
			"state": "solid",
		},
		"Be": map[string]string{
			"name":  "Beryllium",
			"state": "solid",
		},
		"B": map[string]string{
			"name":  "Boron",
			"state": "solid",
		},
		"C": map[string]string{
			"name":  "Carbon",
			"state": "solid",
		},
		"N": map[string]string{
			"name":  "Nitrogen",
			"state": "gas",
		},
		"O": map[string]string{
			"name":  "Oxygen",
			"state": "gas",
		},
		"F": map[string]string{
			"name":  "Fluorine",
			"state": "gas",
		},
		"Ne": map[string]string{
			"name":  "Neon",
			"state": "gas",
		},
	}

	if el, ok := elements["Li"]; ok {
		fmt.Println("Element:", el["name"], "State:", el["state"])
	}
}
