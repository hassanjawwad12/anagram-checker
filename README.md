# anagram-checker
Check if 2 strings are anagrams using go

### Code Detail

* Instead of using 2 if conditions in chronological order in the last loop you can also do this
```go 
if targetCount, ok := targetMap[letter]; !ok || sourceCount != targetCount {
    return false
}
```
* I used the `make(map[rune]int)` to also handle the specialm character in the string but if you only want for alphabetic strings u can also use `make(map[string]int)`