# Go Slices – Quick Reference  

Based on [Go Slices: usage and internals](https://go.dev/blog/slices-intro)  

---

## Slicing a Slice  
```go
b := []byte{'p', 'a', 'r', 't', 'h'}

fmt.Println(b[1:4]) // slice from index 1 to 3
fmt.Println(b[:2])  // first two elements
fmt.Println(b[2:])  // from index 2 to end
fmt.Println(b[:])   // entire slice
```

---

## Bytes and ASCII  
- `byte` in Go is an alias for `uint8`.  
- Printing `[]byte` shows ASCII codes, not characters.  

| Char | ASCII |
|------|-------|
| 'p'  | 112   |
| 'a'  | 97    |
| 'r'  | 114   |
| 't'  | 116   |
| 'h'  | 104   |

Convert to string for characters: `fmt.Println(string(b)) // "parth"`

---

## Copy Function  
```go
func copy(dst, src []T) int
```
- Copies elements from `src` into `dst`.  
- Returns number of elements copied.  
- Handles overlapping slices safely.  

---

## The `...` Operator in Go

The `...` (ellipsis) has different meanings depending on context:

1. **In function parameter → variadic function**
   Declares a parameter that accepts **zero or more values** of the given type.

   ```go
   func f(x ...int) {
       fmt.Println(x) // x is a []int
   }
   ```

2. **In function call → expand a slice into arguments**
   Expands a slice into a list of individual values.

   ```go
   f([]int{1, 2, 3}...) // same as f(1, 2, 3)
   ```

3. **In array literal → infer length**
   Lets the compiler determine the length of the array from the elements provided.

   ```go
   arr := [...]int{1, 2, 3} // type [3]int
   ```

---

