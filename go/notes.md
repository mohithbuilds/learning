# Rough Notes/Thoughts

## Modules
  - Your code's module is defiened by a `go.mod` file that tracks the modules that provide those packages.
  - The `go.mod` file stays in your source code repository.
    - It's responsible for dependency management.
    - Create the `go.mod` file by running `go mod init <module-name>`.
      - The name is the module's module path.
      - In actual development, the module path will typically be the repository location where your source code will be kept.
        > i.e. the module path might be `github.com/mymodule`

## Packages
  - A package here (in go) is a way to group functions, and it's made up of all the files in the same directory.
  > Go code is grouped into packages, and packages are grouped into modules

### `go.sum` files
  - This is the file that contains cryptographic hashes of **all** dependencies of that module.
    - Includes for any transitive dependencies
  - `go mod tidy` adds missing hashes and will remove any unnecessary hashes from go.sum
  - Each line in the file has fields separated by spaces
    1. a module path
      - the name of the module the hash belongs to
    2. a version (possibly ending with /go.mod)
      - the version of the module the hash belongs to
      - if the version ends with `/go.mod`, the hash is for the modules `go.mod` file only
        - otherwise, the hash is for the files within the module's .zip file
    3. a hash
      - consists of an algorithm name (like h1) and a base64-encoded cryptographic hash, separated by a colon (:)
      - right now SHA-256 (h1) is the only supported hash algorithm
  - The file may contain hashes for multiple versions of a module
