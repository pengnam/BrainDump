# Profiling Go Code

Enable profiling to profile Go code. 

## Initial Thoughts

Computer performance will strongly affect the profiling of Go code. Modern CPUs rely heavily on active thermal management that can affect benchmark results.

The general method of deciphering the expected complexity is for the program to go loop finding. 

# Profile
A Profile is a collection of stack traces showing the call sequences that led to instances of a particular event such as allocation. Packages can create and maintain their profils, most common use is for tracking resources that must be explicity closed. 

# pprof
pprof operates on data in the profile.proto format. Each file is a collection of samples, where each sample is associated to a point in a location hierachy, one or more numeraic values, and a set of labels. 



Types of profiles:
1. goroutine: stack traces of all current goroutines
2. help: sampling of heap allocation
3. threadcreate: stack traces that led to the creation of new OS threads
4. block: stack traces that led to blocking on synchronization primitives
5. mutex: stack traces of holders of contended mutexes

# Useful links
https://blog.golang.org/profiling-go-programs
https://flaviocopes.com/golang-profiling/
