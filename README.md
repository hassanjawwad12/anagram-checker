# anagram-checker
Check if 2 strings are anagrams using go

### Code Detail

* Instead of using 2 if conditions in chronological order in the last loop you can also do this
```go 
if targetCount, ok := targetMap[letter]; !ok || sourceCount != targetCount {
    return false
}
```
* I used the `make(map[rune]int)` to also handle the special characters in the string but if you only want to check for alphabetic strings u can also use `make(map[string]int)`