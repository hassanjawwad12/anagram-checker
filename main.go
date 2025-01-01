package main

import "fmt"

func CheckIfStringsAreAnagram(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}

	sourceMap := make(map[rune]int) // Store frequency of each charcater
	targetMap := make(map[rune]int)

	// The string is iterated and character with its count is stored
	for _, letter := range s {
		sourceMap[letter]++
	}

	for _, letter := range t {
		targetMap[letter]++
	}

	// Iterates over sourceMap to compare its values with targetMap
	for letter, sourceCount := range sourceMap {

		// If a character exists in sourceMap but not in targetMap the strings are not anagrams.
		targetCount, ok := targetMap[letter]
		if !ok {
			return false
		}

		// If the frequency of a character in sourceMap (sourceCount) does not match the frequency in targetMap (targetCount), the strings are not anagrams.
		if sourceCount != targetCount {
			return false
		}
	}

	return true
}

func main() {

	reponse := CheckIfStringsAreAnagram("heart", "earth")
	fmt.Println(reponse)

	result := CheckIfStringsAreAnagram("hello", "olli")
	fmt.Println(result)
}
