package main

import (
	"fmt"
	"reflect"
)

func teste_make() {
	var intSlice = make([]int, 10)        // when length and capacity is same
	var strSlice = make([]string, 10, 13) // when length and capacity is different

	fmt.Printf("intSlice \tLen: %v \tCap: %v\n", len(intSlice), cap(intSlice))
	fmt.Println(reflect.ValueOf(intSlice).Kind())

	strSlice = append(strSlice, "1", "2", "3", "4", "5", "6", "7", "8", "9", "10")
	fmt.Printf("strSlice \tLen: %v \tCap: %v\n", len(strSlice), cap(strSlice))
	fmt.Println(reflect.ValueOf(strSlice).Kind())
}

func teste_map() {
	person := "Anns"
	attended := map[string]bool{
		"Ann": true,
		"Joe": true,
	}

	if attended[person] { // will be false if person is not in the map
		fmt.Println(person, "was at the meeting")
	}
}

func main() {

}
