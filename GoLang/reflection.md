# Reflection

## Introduction

Reflection is the ability of a program to examine its own structure, especially through types as a form of metaprogramming. 

GoLang is statically typed.

Consider interface, an important sub type. 

Interface Variable can store any concrete value as long as value implements interface methods. 

This implies that an empty interface is satisfied by any value, since any value has zero or more methods. 

Go's interfaces are statically typed (all of them are), a variable of interface type always has the same static type. 

Russ Cox has written a detailed blog post about the representation of interface values. Interface values store a pair, concrete value assigned to variable, and taht value's type descriptor. 

Static type of interface determines what methods may be invoked with an interface variable, even though the concrete value inside may have a larger set of methods. 

Empty interface will again contain that same pair. 

Type assertion is needed for interfaces that do not inherit things that are of a direct subset of another. 

## General idea of reflect package

Implements run-time reflection, allowing program to manipulate objects with arbitrary types. Use: take a value with static type interface and extract its dynamic type information by calling "TypeOf", which returns a Type

Reflection goes from interface value to reflection object, and refleciton object to interface value



