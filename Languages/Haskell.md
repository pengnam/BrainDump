# Haskell

Exerpts from "Real World Haskell".


## Features

### Characteristics
1. Strongly typed

An expression that obeys language type rules is "well typed", an expression that disobeys is "ill typed" and will cause a "type error". 

Haskell's strong typing will not automatically coerce values from one type to another. 

(Strong type means to have more typing rules. An example is coercsion tends to imply weak typing)

2. Static typed

Type is know at compile time. 

This is in contrast to *duck* typing. However, Haskell's *typeclasses* provides the benefit of dynamic typing in a safe and convenient form. 


3. Type inference

Compiler can deduce type of almost all expressions. Type need not be explicitly declared.

### Aspects of language

1. Expression orientated

This means that there is a distinction between an if *expression* and if *statement*. Expressions have to evaluate to a value, hence all if statements in haskell needs an *else* clause.


## Functions

### Useful default functions

*take, drop, head, tail, fst*

### Variables 

Once variables are bound to a particular expression, its value does not change. 
We can always use variable instead of writing out the expression. 


### Lazy Evaluation

Strict evaluation: arguments to a function evaluated before the function is applied

Haskell uses non-strict evaluation. 

Non-strict evaluation means that a *thunk* (something like a promise) is created. Only when the value is needed then the value is evaluated. 


### Function Types and Purity

*side effect*: introduces a dependency between the global state of the system and behaviour of a funciton. 

In haskell code, we should separate *pure* functions from *impure* functions. 

### Data Types

Construct new data types using *data* keyword.

*Currying* is partial function application that takes only one argument. 


### Defining new data type

```
data DataType = Book Int String [String]
                  deriving (Show)
```

algebraic data types can have more than 1 value constructor. 

data Bool = False | True


### Partial functions vs Total Functions

Total functions are functions that return valid results over their entire input domains.


### Guards vs if-else vs Case
Case is used when there are *multiple code paths* and every code path is guided by the structure of the value. (i.e. Pattern matching)

Guards make decisions based on the *value*. 

https://stackoverflow.com/questions/9345589/guards-vs-if-then-else-vs-cases-in-haskell

