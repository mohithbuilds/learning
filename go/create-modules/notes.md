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
