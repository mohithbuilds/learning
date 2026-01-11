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
