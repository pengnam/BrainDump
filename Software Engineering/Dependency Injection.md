# Dependency Injection

## Overview

A technique that has one object supplying the dependencies of another object. 

Allows for a *seperation of concern* between the construction and the use of objects. 

Allows programs to be loosly coupled.


## Types of Dependency Injection

### Constructor Injection

Constructor is used to decide an object's dependencies. Object declares in constructor the objects that it needs.
Register constructor (and relevant parameters) into the provider.


### Setter Injection

Defining setter methods in objects that need injection. 


### Interface Injection

An interface can be used to perform the injection through. 

## Alternative: Service locator

Object calls service locator object to find its dependencies. 

This, however, creates a dependency between the service locator and the object.

### Notes:

There seems to be different ways of registering a provider. 

One way would be to register with a configuration file, another one is by calling methods. I believe both methods are similar.

## Service Locator vs Dependency Injection

### Similarties:

Both allow fundamental decoupling that's missing in the naive example. 

Both gives inversion of control, but it comes at a price. It is potentially hard to understand and it leads to potential problems. 

About equally difficult to test. Dependency injection will be easier with IDE support to search up references. 

### Differences:

```
So the primary issue is for people who are writing code that expects to be used in applications outside of the control of the writer. In these cases even a minimal assumption about a Service Locator is a problem.
```

## Constructor vs Setter Injection

The choice between setter and constructor injection is interesting as it mirrors a more general issue with object-oriented programming - should you fill fields in a constructor or with setters.

Martin Fowler's principle is that as much as possible create valid objects at run time. Constructors with parameters give a clear statement of what it means to create a valid object.
