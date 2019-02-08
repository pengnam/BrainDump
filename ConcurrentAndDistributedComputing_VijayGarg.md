# Concurrent and Distributed Computing

## Chapter 1: Introduction
A _parallel_ system consists of multiple processors that communicate with each other using shared memory.
Parallel processes with multple processors communicate with each other using shared memory.

_Distributed_ systems are computer systems that contain multiple processors connected by a communication network.
Processors communicate with each other using messages sent over the network.
These systems are increasingly available because of decrease in prices of the network.

### Comparing parallel and Distributed

Distributed has better scalability, greater modularity, lower cost. However, it is faster
to update a shared memory location than to send a message to another processor.
More efficient to get parallelism.

### Design Goals

- Fault Tolerance

- Transparency

- Flexibility

- Scalability

## Chapter 2: Mutual Exclusion Problem
### Introduction
When processes share data, it is important to synchronise their access to the data
so that updates are not lost as a result of concurrent accesses and the data
sets are not corrupted.

The key idea is that some operations have to be executed _atomically_ i.e. has a critical section.

The problem of ensuring that a critical section is executed atomically is called the mutual exclusion problem.

### Peterson's Algorithm
#### Attempt 1:
Use of a open door variable.

Problem:
The testing and setting of the open door variable is not done _atomically_.

#### Attempt 2:
Use of a boolean array.

Problem: Can deadlock.

#### Attempt 3:
Use of a turn variable

Problem:
Setting turn not atomic, does not guarantee ME.


#### Final Attempt:
##### Proof for mutual exclusion:

Lets assume that process 0 has just entered critical section. (T1)
--> either wantCS[1] == false or turn == 0

Case 1: wantCS[1] is false

This means that P1 has to be currently  at line 5 or before, and is not currently in CS.
However, when P1 eventually reaches line 8, turn will be set to 0 and since wantCS[0] is true as P0 is in CS, P1 will stop at line 8.

Hence case 1 is invalidated.

Case 2: turn == 0

When P1 approaches the critical section, it sets turn to 0 after P0 sets turn 1 for turn == 0 at line 8 for P0. However, P1 will hence be kept waiting at line 8 since turn == 0 and wantCS will be true.

QED


##### Proof for starvation free:
Contradiction: Process eventually enters the critical section

#TODO: I'm thinking this explanation is not rigorous enough

Case 1: P0 is waiting, P1 is trying to enter

Implies wantCS[1] is true and turn == 1. Hence, P1 will not be kept waiting at line 8.

Case 2: P1 is waiting, P0 is trying to enter <br>
Symmetric to case 1.

##### Proof for progress:
Contradiction: Neither of the processes can enter critical section.

If P0 is waiting,
Implies wantCS[1] is true and turn == 1. Hence, P1 will not be kept waiting at line 8.

...

### Lamport's Bakery Algorithm

#### Proof for mutual Exclusion
Consider a process that has the smaller process id

Proof by contradiction: both processes are currently in the CS
<Define smaller/larger process>
Consider: smaller processes is currently in the CS
Let j be the index of the larger process.
This means that there are either number[j] == 0 or the other larger process is smaller than this process.

Since the larger process is assumed to be in CS, number[j] != 0. The larger process is also larger than the smaller process so it can't be true.

#### Proof for progress
Proof by contradiction: both are stuck outside CS

1) Can't be stuck at the first while loop
2) Both can't be stuck at the same second while loop


## Chapter 3: Semaphores
### Monitors
high level object oriented construct for synchronisation in concurrent programming.

It can be viewed as  class with data variables and methods to manipulate the date. Because there exists multiple thread that can access the data, monitors have entry methods to guarantee at most one thread can be executing in any one time.
