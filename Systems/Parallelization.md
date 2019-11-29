# Parallel Computing
## Background

![](assets/Parallelization-721166aa.png)

![](assets/Parallelization-3738f8df.png)

## Parallel Computer Memory Organisation

1. Distributed-Memory

2. Shared-Memory

    - Uniform Memory Access

    - Non-Uniform Memory Access

    - Cache-only Memory Access


3. Hybrid (Distributed-Shared Memory)

## Processor Performance Gain

Levels of parallelisation:

- Bit level

- Instruction level

- Thread level

- Process level

- Processor level

### Instruction Level Execution
Pipelining:

Splitting the instruction into multiple stages. allowing multiple instructions to occupy different stages in the same clock cycle.

Superscalar:

Allowing multiple instruction to pass through the same stage. Scheduling how instructions are executed together.

## Steps

1. Decomposition

breaking a task into sub components

2. Scheduling

of tasks to processes/threads.
3. Mapping

of processes to cores


## Process

A process is a program in execution.

A *process* is broken into several components.

1. Executable program

2. Global data


Process are abstraction that let you exploit multiple cores/parallel processes.

Processes were initially created to allow for multitasking.

There are two types of multitasking.

1. Time slicing of execution
2. Parallel execution on multiple resources.

## Parallel Architecture Taxonomy

1. Single Instruction Single Data (SISD)

- A single instruction stream is executed.
- Each instruction work on single data
- Most uniprocessors fall into this category.


2. Single Instruction Multiple Data (SIMD)

- Single stream of instructions
- Each instruction works on multiple data
- Exploits data parallelism

3. Multiple Instruction Single Data

̨

## Parallel Programming Pattern

### Parallelism Concepts

1. Data Parallelism

Same operation applied to different elements. *Operations are independent*

Partition the data to solve problem.

Exploited by SIMD computers extensively. (Also SPMD).



2. Task Parallelism

Partition the task to solve the problem.

Analysed using a *task dependence graph*. Form direct acyclic graph, which represents the control depdendency between tasks.

Analysed using:

- Critical path length: fastest completion time

- Degree of concurrency = Total work/ critical path length/

### Parallelism Models

1. Fork-Join

Task T creates a number of child tasks. Tasks then work in parallel to excecute the program. After that, a join is used.


2. Parabegin-Parend

Specifies function call to be executed.

3. SPMD&SIMD

Same program executed on different processors but operate on different data. Single instructions executed synchronously by the different threads on different data.

4. Master-Slave (Worker)

Master assigns tasks to slaves.

5. Client-Server

MPMD model.

Server computes requests from multiple client tasks concurrently.

6. Pipelining

Passing data between components of a pipeline.

7. Task Pool

Data structure that threads can access to retrieve tasks for execution.

8. Producer Consumer

Producer produces data into a common buffer which are then consumed by the consumers.
