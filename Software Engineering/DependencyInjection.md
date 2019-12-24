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

## Service Locator vs Dependency Injection

### Benefits:

Both allow fundamental decoupling that's missing in the naive example. 

### Notes:

There seems to be different ways of registering a provider. 

One way would be to register with a configuration file, another one is by calling methods. I believe both methods are similar.
