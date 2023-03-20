/* Write a function that takes in a slice of any type and a value of the same type, and returns the index
of the first occurrence of the value in the slice. If the value is not present in the slice, the function should return -1. */

package main

import (
	"fmt"
	"reflect"
)

func main() {
	slicefruit := []string{"apple", "banana", "cherry"}
	fruitIndexValue := indexOf(slicefruit, "apple")
	fmt.Println(fruitIndexValue)

	slicenumbers := []int{1, 2, 3, 100, 4, 5}
	numberIndex := indexOf(slicenumbers, 100)
	fmt.Println(numberIndex)

	values := []bool{true, false, true, false}
	boolIndex := indexOf(values, true)
	fmt.Println(boolIndex)

}
func indexOf(slice interface{}, value interface{}) int {
	s := reflect.ValueOf(slice)
	for i := 0; i < s.Len(); i++ {
		if reflect.DeepEqual(value, s.Index(i).Interface()) {
			return i
		}
	}
	return -1
}
