/* Q4) Write a function that takes in a list of strings and a function that determines whether a given string meets certain criteria,
and returns a new list containing only the strings from the original list that pass the function's test.

  Example input: ["apple", "banana", "cherry"]
Example output: ["banana", "cherry"]   */

package main

import "fmt"

func main() {
	input := []string{"apple", "banana", "cherry"}
	output := filterString(input, func(s string) bool {
		return len(s) > 5
	})
	fmt.Println(output)

}

func filterString(strings []string, filter func(string) bool) []string {
	stringFilterd := []string{}
	for _, svalue := range strings {
		if filter(svalue) {
			stringFilterd = append(stringFilterd, svalue)
		}
	}
	return stringFilterd
}
