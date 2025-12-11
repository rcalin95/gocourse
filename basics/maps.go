package basics

import (
	"fmt"
	"maps"
)

func main() {

	// var a map[key]valueType
	// mapVariable := make(map[string]int)

	// using a Map literal
	// mapVariable := map[keyType]	valueType{
	// 	"key1": value1,
	// 	"key2": value2,
	// }

	myMap := make(map[string]int)
	fmt.Println(myMap)

	myMap["key1"] = 9
	myMap["key2"] = 18
	fmt.Println(myMap)
	fmt.Println(myMap["key1"])
	fmt.Println(myMap["key"]) // initial value 0 for int

	myMap["code"] = 21
	fmt.Println(myMap)

	delete(myMap, "key2")
	fmt.Println("After deletion:", myMap)

	// clear(myMap)
	// fmt.Println("After clearing:", myMap)

	_, ok := myMap["key1"]
	if ok {
		fmt.Println("Value at Key1 is present")
	} else {
		fmt.Println("Value at Key1 is not present")
	}
	fmt.Println("Present?", ok)

	myMap2 := map[string]int{
		"alpha": 1,
		"beta":  2,
	}
	myMap3 := map[string]int{
		"alpha": 1,
		"beta":  2,
	}
	fmt.Println("Map 2:", myMap2)

	if maps.Equal(myMap3, myMap2) {
		fmt.Println("Maps are equal")
	}

	for _, value := range myMap3 {
		fmt.Printf("Value: %d\n", value)
	}

	var myMap4 map[string]string

	if myMap4 == nil {
		fmt.Println("myMap4 is initialized to nil")
	} else {
		fmt.Println("myMap4 is not initialized to nil")
	}

	val := myMap4["key"]
	fmt.Println("Value at key in nil myMap4:", val)

	// myMap4["key"] = "Value"
	// fmt.Println(myMap4)

	myMap4 = make(map[string]string)
	myMap4["key"] = "Value"
	myMap4["name"] = "Gopher"
	myMap4["language"] = "Go"
	fmt.Println("After initializing myMap4:", myMap4)

	fmt.Println("Length of myMap4:", len(myMap4))

	myMap5 := make(map[string]map[string]string)

	myMap5["map1"] = myMap4
	fmt.Println("Nested Map:", myMap5)
}
