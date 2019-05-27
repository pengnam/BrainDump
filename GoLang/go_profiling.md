# Profiling Go Code

Enable profiling to profile Go code. 

## Initial Thoughts

Computer performance will strongly affect the profiling of Go code. Modern CPUs rely heavily on active thermal management that can affect benchmark results.

The general method of deciphering the expected complexity is for the program to go loop finding. 

# Profile
A Profile is a collection of stack traces showing the call sequences that led to instances of a particular event such as allocation. Packages can create and maintain their profils, most common use is for tracking resources that must be explicity closed. 

# pprof
pprof operates on data in the profile.proto format. Each file is a collection of samples, where each sample is associated to a point in a location hierachy, one or more numeraic values, and a set of labels. 

# Stack traces

Stack traces at time of allocation: hence, might be for code that is not running anymore, which means functions to blame (for example, a function that frees memory space, is not the same as the function allocating the code)

# alloc_space vs inuse_space
alloc_space: allocated memory and object counts
inuse: in-use memory and object counts

Inuse metrics are important in determining the amount of memory being used. 
Allocations describe the performance/time spent on garbage collection

MemProfile rate records the fraction of memory allocations that are recorded and reported in the memory profile.

Types of profiles:
1. goroutine: stack traces of all current goroutines
2. help: sampling of heap allocation
3. threadcreate: stack traces that led to the creation of new OS threads
4. block: stack traces that led to blocking on synchronization primitives
5. mutex: stack traces of holders of contended mutexes

# Useful links
https://blog.golang.org/profiling-go-programs
https://flaviocopes.com/golang-profiling/
https://jvns.ca/blog/2017/09/24/profiling-go-with-pprof/



