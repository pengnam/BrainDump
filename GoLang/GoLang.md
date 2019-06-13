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
Exported, unexported types are determined by whether the first letter is upper case or lower case


## Handling Timeouts

SetDeadline:
Deadlines are not timeouts. Once set they stay in force forever no matter if and how the connection is used.
Call set deadline before each and every read/write operation.

All timeouts are implemented in terms of deadlines. They *do not* reset every time data is sent or received.

Critical for HTTP server that is exposed to the internet to enforce timeouts on client connections, else slow or disappearing clients might leak file descriptors.

ReadTimeout covers the time from when connection is accepted to when the request body is fully read. Implemented using SetReadDeadline immediately after accept.

## Context Pattern
Each incoming request is handled as its own goroutine. Request handlers might spawn additional goroutines to access identity specific values but if the request is cancelled or times out, all goroutines working on request should exit quickly. Context package makes it easy to pass request scoped values, cancellation signals, and deadlines available as a context. 

## On interfaces and implementing methods
golang interface method requires pointer receiver
https://stackoverflow.com/questions/40823315/x-does-not-implement-y-method-has-a-pointer-receiver


## Useful Links

https://queue.acm.org/detail.cfm?id=2839461
(*) https://blog.cloudflare.com/the-complete-guide-to-golang-net-http-timeouts/


