/* Q1) Write a function that takes in a slice of integers and returns the sum of all the elements in the slice.

print(sum_slice([1, 2, 3, 4, 5]))  # Output: 15
print(sum_slice([-1, -2, -3, -4, -5]))  # Output: -15
print(sum_slice([0, 0, 0, 0, 0]))  # Output: 0  */

package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3, 4, 5} // Output: 15
	fmt.Println(sum_slice(slice1))

	slice2 := []int{-1, -2, -3, -4, -5} // Output: -15
	fmt.Println(sum_slice(slice2))

	slice3 := []int{0, 0, 0, 0, 0} // Output: 0
	fmt.Println(sum_slice(slice3))
}
func sum_slice(nums []int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}
