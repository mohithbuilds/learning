# Go Generics Tutorial Notes

## Adding a Generic Function to Handle Multiple Types

This section details how to replace multiple non-generic functions with a single generic function that can receive a map containing either integer or float values.

### Declaring Generic Functions

To support values of different types, a single function will need a way to declare what types it supports. This is achieved by declaring **type parameters** in addition to its ordinary function parameters. These type parameters make the function generic, enabling it to work with arguments of different types.

### Type Parameters and Type Arguments

*   You'll call the function with **type arguments** and ordinary function arguments.
*   Each type parameter has a **type constraint** that acts as a kind of meta-type for the type parameter.
*   Each type constraint specifies the permissible type arguments that calling code can use for the respective type parameter.
*   While a type parameter’s constraint typically represents a set of types, at compile time the type parameter stands for a single type – the type provided as a type argument by the calling code.
*   If the type argument’s type isn’t allowed by the type parameter’s constraint, the code won’t compile.

### Importance of Type Constraints

A type parameter must support all the operations the generic code is performing on it. For example, if your function’s code were to try to perform string operations (such as indexing) on a type parameter whose constraint included numeric types, the code wouldn’t compile. In the `SumIntsOrFloats` example, you'll use a constraint that allows either integer or float types.

### Example: `SumIntsOrFloats` Function Declaration

Declare a `SumIntsOrFloats` function with two type parameters (inside square brackets), `K` and `V`, and one argument that uses the type parameters, `m` of type `map[K]V`. The function returns a value of type `V`.

*   **`K` type parameter:**
    *   Specify the type constraint `comparable`.
    *   The `comparable` constraint is predeclared in Go.
    *   It allows any type whose values may be used as an operand of the comparison operators `==` and `!=`.
    *   Go requires that map keys be `comparable`, so declaring `K` as `comparable` is necessary to use `K` as the key in the map variable `m`.
    *   It also ensures that calling code uses an allowable type for map keys.

*   **`V` type parameter:**
    *   Specify a constraint that is a **union of two types**: `int64` and `float64`.
    *   Using `|` specifies a union of the two types, meaning that this constraint allows either type. Either type will be permitted by the compiler as an argument in the calling code.

*   **`m` argument:**
    *   Specify that the `m` argument is of type `map[K]V`, where `K` and `V` are the types already specified for the type parameters.
    *   `map[K]V` is a valid map type because `K` is a `comparable` type. If `K` hadn’t been declared `comparable`, the compiler would reject the reference to `map[K]V`.

### Type Inference

You can omit type arguments in calling code when the Go compiler can infer the types you want to use. The compiler infers type arguments from the types of function arguments.

**Note:** This isn’t always possible. For example, if you needed to call a generic function that had no arguments, you would need to include the type arguments in the function call.

### Constraint Interfaces

You declare a type constraint as an interface. The constraint allows any type implementing the interface. For example, if you declare a type constraint interface with three methods, then use it with a type parameter in a generic function, type arguments used to call the function must have all of those methods. Constraint interfaces can also refer to specific types.
