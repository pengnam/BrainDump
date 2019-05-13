# GoLang

## Basic Syntax Ideas

1. Uses tokens. There are four classes: identifiers, keywords, operators, punctutation, and literals. 
2. Uses semicolons as terminators. However, Go programs may omit these semicolons following some rules. 
3. Common types: ArrayType, StructType, PointerType, FunctionType, SliceType, MapType, ChannelType

4. Initialization of variables 
- Types are declared after the variable. 
- Variables declared without a corresponding initialization are zero-valued. 
- Either var <name> or <name> := can be used to initialize variables


## Packages

The first statement in a Go source file must be:

```
package name
```

## Basic Data Strutures

- Slices: Supports append operations
- Arrays: built-in data type
- Map :Associative data type

## Functions
-support for multiple return values. (and hence, multiple assignment) 
-variadic functions


## Iteration
for index, ele := range sliceName {}
for i := 0; i < len(sliceName) ; i++ {}

## Pointers
Go supports pointers, allowing you to pass references to values and records within your program.

## Structs

## Methods
Go support methods on struct types. 

## Interfaces

Named collections of method signatures.
1. It is a set of methods
2. It is a type

Core concept in Go's type system: 
Instead of designing our abstractions in terms of what kind of data it holds, we design abstractons based on what actions our types can execute.

The interface{} type, the empty interface, is the source of much confusion. The interface{} type is the interface that has no methods. Since there is no implements keyword, all types implement at least zero methods, and satisfying an interface is done automatically, all types satisfy the empty interface. That means that if you write a function that takes an interface{} value as a parameter, you can supply that function with any value. 


## Miscellaneous Ideas
- GoLang supports anonymous functions, which can form closures. 
- 
