package main

import (
	"fmt"
	"maps"
)

func main() {
	myMap := make(map[string]int)
	//fmt.Println(myMap)

	myMap["key1"] = 1
	myMap["code"] = 12

	fmt.Println(myMap["key1"])
	fmt.Println(myMap["key2"])

	myMap["code"] = 18

	delete(myMap, "key1")

	fmt.Println(myMap)

	myMap["key1"] = 1
	myMap["key2"] = 12

	value, association := myMap["key2"]

	//clear(myMap)
	fmt.Println(value)
	fmt.Println(association)

	myMap2 := map[string]int{"key1": 1, "key2": 2}
	fmt.Println(myMap2)

	fmt.Println(maps.Equal(myMap2, myMap))

	for key, value := range myMap2 {
		fmt.Println(key, value)
	}

	key4, value4 := myMap2["hh"]

	fmt.Println(key4, value4)

	var myMap4 map[string]int
	fmt.Println(myMap4)
	fmt.Println(myMap4 == nil)

}
