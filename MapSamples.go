package main

import "fmt"

func main() {

	//map should be initialized, before using it.
	//var x map[string]int
	//	x["AAA"] = 1

	map2 := make(map[string]bool)
	map2["AAA"] = false
	fmt.Printf("%v\n", map2)

	map1 := map[string]int{
		"AAA": 50,
		"BBB": 21,
		"CCC": 41,
		"DDD": 29,
	}

	fmt.Printf("%#v\n", map1)
	fmt.Printf("%v\n", map1)

	fmt.Println("val of AAA = ", map1["AAA"])

	map1["EEE"] = 40
	fmt.Println(map1)

	delete(map1, "EEE")
	fmt.Println(map1)

	value, valueStatus := map1["key"]
	fmt.Println("valueStatus:", valueStatus)
	fmt.Println("value:", value)

	//Map insde a map
	mapInsideMap := map[string]map[int]string{
		"H": map[int]string{
			1: "name1",
			2: "name2",
		},
		"He": map[int]string{
			3: "name3",
			4: "name4",
		},
	}
	fmt.Println(mapInsideMap)

}

/*
map[AAA:false]
map[string]int{"AAA":50, "BBB":21, "CCC":41, "DDD":29}
map[AAA:50 BBB:21 CCC:41 DDD:29]
val of AAA =  50
map[EEE:40 AAA:50 BBB:21 CCC:41 DDD:29]
map[AAA:50 BBB:21 CCC:41 DDD:29]
valueStatus: false
value: 0
map[H:map[1:name1 2:name2] He:map[4:name4 3:name3]]
[Finished in 0.344s]*/
