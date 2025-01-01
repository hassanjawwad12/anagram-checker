# anagram-checker
Check if 2 strings are anagrams using go

### Condition inside the loop
Instead of using 2 if conditions in chronological order in the last loop you can also do this
```go 
if targetCount, ok := targetMap[letter]; !ok || sourceCount != targetCount {
    return false
}
```
