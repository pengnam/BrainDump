# System Tools

## IOStat

IOStat is primarily used to decide if disk is or is not a bottleneck. 


## Background

There are different layers in the linux IO subsystem. 
Each layer abstracts out a specific operation. 

https://blog.serverfault.com/2010/07/06/777852755/

Notable ideas include the fact that there is page caching. Before it writes data to disk it puts in memory, adn before it reads data from disk it checks to see if it is in memory already. 

iostat shows requests to the device and not read and write requests from the user space. Table in IOstat is reading below the disk cache. 

IOStat can break down the statistics at both the partition levelthen the device level. 

## IOStat output
IO CPU Wait time, device utilization. 
Below disk cache layer, and does not provide information about cache miss hit ratios. 

tps: number of IO operations per second
iowait: CPI metric that measures the percentage of time the CPU is idle while waiting for an IO operation to complete.

## Profiling

### Profiling I/O 

Kernels usually provide facilities for tracing I/O types. I/O is mostly off CPU time, tracking I/O is a form of off CPU analysis

Note that synchronous IO operations cause application latency, and the resulting flame graph is a direct measurement of 

Asynchronous IO operations do not directly affect the the out and run time, but it can affect other calculations. 

Block IO flame graphs vs normal IO flame graphs



