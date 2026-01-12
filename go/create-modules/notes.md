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

### Go Slice
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
