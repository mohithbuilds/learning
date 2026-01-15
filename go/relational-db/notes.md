# Relational DB Example

**Using postgres instead of MySQL**

Sections:
1. Set up a database
2. Import the database driver
3. Get a database handle and connect
4. Query for multiple rows
5. Query for a single row
6. Add data

Since I used Postgresql unlike the example, I used the `github.com/jackc/pgx/v5` Go module, `pgx` library.
- Also a bit different from the example in the sense that for this module we need `context` which is something I don't quite understand right now but will look into and learn.
@TODO Look into context package which is part of Go's standard library

## defer

What is `defer`?
- `defer` is a keyword in Go that schedules a function call to be run right before the function that contains the `defer` statement finishes executing and exits.
- So, it schedules the call: it puts the call you are deferring on a list of deferred calls for the function it's called in.
- It also **guarantees execution** right before the function it's called in exits, for any reason, whether it finishes normally, panics, or hits a return statement somewhere.
  > [NOTE] `defer` does **NOT** run before a call to `os.Exit()` because this function terminates the program immediately. It does **NOT** unwind the stack, and therefore does not run any deferred calls.
  > Also since `log.Fatal()` is a wrapper that prints the log message and then immediately calls `os.Exit()` it also wouldn't run the deferred call.
  - So this means that `panic` unwinds the current function's execution stack, and as it does so, it executes any deferred function calls.

When do you use it?
- It's most commonly used for cleanup tasks like closing files or database connections.
> Placing the `defer` right after the resource is acquired [like we did with the connection to the db] makes the code **cleaner** and **safer**, because you can't forget to release the resource.

## `fmt.Errorf(...)` vs `log.Fatalf(...)`

### `fmt.Errorf(...)` -> Returns an error
- It formats a string and wraps it in a standard `error` type, then **returns** it.
  > **It doesn't stop your program**
When to use it?
- Inside a regular function (like our `booksByAuthor` function)
  - Because if the function fails, it should **NOT** crash the whole application. Instead, it should `return` an error to it;s caller, letting the caller decide how to handle it.

### `log.Fatalf(...)` -> Exits the program
- It prints the error message to the console and then immediately calls `os.Exit(1)`, which **terminates the entire application**
When to use it?
- When there's a fatal, unrecoverable error and there's no point in continuing to run, so log the fatal error and exit.

In short:
- Errorf: "Something went wrong in this function, I'm telling my caller about it."
- Fatalf: "Something has gone catastrophically wrong, the whole program must shut down now."

## What does `&` do?
- In Go, when you pass a variable to a function, it's typically **passed by value**.
  - This means the function receives a **copy** of the variable's value. If the function modifies that copy, the original variable outside the function remains unchanged.
- The `&` operator is the **address of operator**. When we write `&bk.ID`, we are getting the **memory address** of the `bk.ID` field. This memory address is a pointer.

## Miscellaneous
Something nice I noticed is that in the `go.mod` for any modules that get added due to being a required by the main module that we tried to get it automatically adds a `// Indirect` comment which is nice for tracking.

The `%v` and `%q` are called **formatting verbs** or **format verbs**.
- `%v` is a general-purpose, default format. It prints the "value" of whatever you give it in its default format.
  - When used with an `error` type, it usually prints the underlying error message string.
- `%q` formats a string value as a **double-quoted** string literal. It also handles escaping any special characters within the string (like newlines, tabs, etc.)
  - It's particularly useful for debugging or when you want to unambiguously see the boundaries of a string, or if the string itself might contain spaces or special characters that could make the log message ambiguous.


`if` with a Short Statement (Very Common and Idiomatic!)
  This is where Go `if` statements get powerful and concise, especially for error handling. You can execute a short
  statement before the condition, separated by a semicolon. The variables declared in this short statement are scoped
  only to the if and else blocks.
  Example:
  ```
  strNum := "12345"
  if num, err := strconv.Atoi(strNum); err != nil {
    fmt.Printf("Error converting %q: %v\n", strNum, err)
  } else {
    fmt.Printf("Successfully converted %q to integer: %d\n", strNum, num)
    // num and err are accessible here in the else block
  }
  // num and err are NOT accessible here, outside the if/else block
  ```
