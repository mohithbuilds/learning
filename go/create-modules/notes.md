# Go features
- a function whose name starts with a capital letter can be called by a function not in the same package
  - This is known in Go as an exported name.
  ```
  func Hello(name string) string
  Hello -> function name
  string -> parameter type
  string -> return type
  ```
- the `:=` operator is a shortcut for declaring and initializing a variable in one line
  - Go uses the value on the right to determine the variable's type.
  The long way would have been:
  ```
  var message string
  message = fmt.Sprintf("Hi, %v. Welcome!", name)
  ```

### In Go, code executed as an application must be in a main package.

Things I noticed when writing this example:
1. If a function returns x number of values/objects then when you call the function you must have x number of variables/wildcards for that.
  - `_` is wildcard for if you don't need the value returned at that position from the function call
  - If you don't handle all returned values the program won't compile, go treats this as a fatal error not just as a warning
2. You must use any variables you declare, if you don't then it won't compile
  - Some languages such as Java & Rust treat this as a warning, but go treats it as a fatal error and doesn't compile
  - This doesn't apply to:
     - Global variables, function parameters (can define parameters in the function signature and not use them in the body), and constants
3. Errors are treated as values rather than a exception based pattern with `try/catch`
  - Why?
    - This design emphasizes clarity and ensures that devs are conscious of every point where a failure can occur.

## Go Slice
  - A slice is like an array, except its size changes dynamically as we add and remove items. (similar to python lists)
  - When declaring a slice, you omit the size in the brackets like this: `[]string`
    - tells Go that the size of the array underlying the slice can be changed
    - If there was a number in the brackets `[]` then it would be an array and it's size cannot be changed once declared since **it's part of the type**.
  - A slice is just a descriptor pointing to a backing array
    - so this also means that slices are passed by **reference** (only the descriptor is copied)
    - an array is passed by **value** (entire copy of actual array is made)
  - Under the hood, a slice is represented as a struct with 3 parts:
    1. **Pointer `array`**: the memory address of the first element in the underlying array that the slice can access
    2. **Length `len`**: the number of elements currently visible in the slice
    3. **Capacity `cap`**: the max number of elements the underlying array can hold before having to reallocate memory

## For Loops
  - Looping through an array or slice:
    ```
    for i, element := range arr {
      ...
    }
    ```
    - The `range` gives the index of the current item and a copy of the item's value at that index
    - Another common practice if you don't need the index is use a `_` wildcard
    If you don't need the element you can just omit that second variables:
    ```
    for i := range arr {
      ...
    }
    ```
  > Important note: you must use index in order to modify the original slice/array
  **Traditional:**
  ```
  for i := 0; i <= len(arr); i++ {
    ...
  }
  ```

## Maps
  - A go map type looks like:
  ```map[KeyType]ValueType```
    - `KeyType` may be any type that is comparable & `ValueType` may be any type at all, including another map!
  This variable `m` is a map of string keys to int values:
    `var m map[string]int`
    - Map types are reference types, like pointers and slices, and so the value of `m` above is `nil` -> it doesn't point to an **initialized** map
    - A nil map behaves like an empty map when reading, but attempting to write to a nil map will cause a **runtime panic**
  - To initialize a map, use built in make function:
  `m = make(map[string]int)`
    - The `make` function is used to allocate and initialize maps, slices, and channels
  - Maps are not safe for concurrent use: not defined what happens when you read and write to them simultaneously
  - When iterating over a map with a range loop, the iteration order is not specified and is not gauranteed to be the same from one iteration to the next.
    - If you need a stable iteration order then you must maintain a separate data structure that specifies the order

## Testing
  - Implement test functions in the same package as the code you're testing
  - Test functions have the form __Test*Name*__, where *Name* says something about the specific test
  - Test functions take a pointer to the testing package's `testing.T type` as a parameter.
    - You use this parameter's methods for reporting and logging from your test.
  - I also ran into something unrelated with regex and noticed something overall, look at `` vs ''
