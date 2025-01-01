package main

import "fmt"

func CheckIfStringsAreAnagram(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}

	sourceMap := make(map[rune]int)
	targetMap := make(map[rune]int)

	for _, letter := range s {
		sourceMap[letter]++
	}

	for _, letter := range t {
		targetMap[letter]++
	}

	for letter, sourceCount := range sourceMap {

		if targetCount, ok := targetMap[letter]; !ok || sourceCount != targetCount {
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
