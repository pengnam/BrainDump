# Hadoop

Hadoop streaming allows one to create and run Map/Reduce jobs with any executable or script as the mapper/reducer

## How does streaming work
Mapper and reducer are executables that read the input from stdin, and emit output to std out. 

Specify executable for mappers, each mapper launches executable as a seperate process when the mapper is initialized. Mapper task runs, converts inputs into lines and feeds lines to stdin of process. 

Prefix of line up to first tab character is the key and the rest of the line will be the value. 

Each executable is a seperate process, converts key value pairs into lines and feeds the line oriented outputs from the stdout of the process. 

## Hadoop Distributed File System (HDFS)
Distributed file system desgined to run on commodity hardware. Provides high throughput access, POSIX requirements for streaming. 

 Hadoop Distributed File System (HDFS)

Distributed file system desgined to run on commodity hardware. Provides high throughput access, POSIX requirements for streaming. 


## Data Format
Hadoop uses google protobuff to specify data format
