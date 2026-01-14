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

## Miscellaneous
Something nice I noticed is that in the `go.mod` for any modules that get added due to being a required by the main module that we tried to get it automatically adds a `// Indirect` comment which is nice for tracking.
