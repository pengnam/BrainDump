#POSIX

POSIX is a family of standards specified to maintain portability between operating systems.

It defines the API along with the command line shells and utility interfaces for software compatibility.

## Background
After unix was selected by general consensus as the standard system interface because its manufacturer neutral, there was a need to develop a common denominator
Especially in light of an increasing number of unix systems.

## Examples of standards
1. Process creation and control
2. Signals
3. File and directory operations
4. Pipes
5. C library
6. I/O Port interfaces


## mmap

`mmap` is a POSIX compliant UNIX syscall that maps files or devices into memory. It is a method of memory-mapped file I/O.

Implements demand paging, because file contents are not read from disk directly and initially do not use physical RAM at all. 

Actual reads are performed in a "lazy" manner after a specific location is accessed. 

