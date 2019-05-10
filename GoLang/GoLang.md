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



## Miscellaneous Ideas
- GoLang supports anonymous functions, which can form closures. 
- 
