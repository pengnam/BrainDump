# Dart

## Important Concepts

1. Everything that is placed in a variable is *an object*, and every object is an instance of a *class*. 

2. Dart is *strongly typed*, and has *type inference*.


## Dependency management

Dependencies are managed by *pub*

Dart uses *packages* to manage shared software such as *libraries* and *tools*.

Pub creates a *.packages* file that maps each package name that app depends on to corresponding package in system cache. 

### Handling dependencies
1. Create a pubspec (a file named pubspec.yaml that lists package dependencies and includes other metadata, such as a name for your package).
2. Use pub to get your package’s dependencies.
3. If your Dart code depends on a library in the package, import the library.Dependencies are managed by *pub*. o



### About Modules

#### io module

Most methods in the IO module run asynchronously.


#### async module

**Key classes: Future, Stream**

*Future* represents computation that is yet available, Future returns the value opf computation when it completes sometime in the future. Futures are often used for potentially lengthy computations. 


*Stream* provides an asynchronous sequence of data. 

Stream.listen: registers callback function that runs each time more data is available.


