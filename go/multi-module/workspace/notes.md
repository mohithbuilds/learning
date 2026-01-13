# My Notes

## Workspace Directory

- The `go work init` command tells go to create a `go.work` file for a workspace containing the modules in the specified Directory
  - In this example it is: `go work init ./hello`
- The `go.work` file has a similar syntax to `go.mod`
- The go directive tells Go which version of Go the file should be interpreted with. It's similar to the `go.mod` file.
- The `use` directive tells Go that the module in the directory should be main modules when doing the build
- So in any subdirectory of `workspace` the module will be active
- The Go command includes all the modules in the workspace as main modules. This allows us to refer to a package in the module, even outside the module. Running the go run command outside the module or the workspace would result in an error because the go command wouldn’t know which modules to use.

- The `go work use` command adds a new module to the go.work file

- Since the two modules are in the same workspace it's easy to make a change in one module and use it in another

### Comands
`go work use [-r] [dir]`
- adds a use directive to the `go.work` file for `dir`, if it exists, and removes the use directory if the argument directory doesn't exist.
- The `-r` flag examines subdirectories of `dir` recursively
`go work edit`
- edits the `go.work` file similarly to `go mod edit`
`go work sync`
- syncs dependencies from the workspace's build list into each of the workspace modules
