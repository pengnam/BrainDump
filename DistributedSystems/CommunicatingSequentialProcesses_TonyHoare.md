# Communicating Sequential Processes by Tony Hoare

## Alternative Notes

*Guarded Command Language* is a language that consists of guarded statements. Guarded statements have a proposition, which must be true before a command is executed.

*Hoare Logic* formal system with a set of logical rules for reasoning rigorously about the correctness of computer programs. Provides axioms and inference rules for all constructs of a simple imperative programming language.

## Processes
A process is defined by describing the whole range of its potential behaviour. It is also designed recursively.

*Kind of like using recursive properties of functions to simulate finite state machines*

## Traces


## Concurrency
Choice between several different actions, for example. On each such occassion, the choice of which even will actually occur can be controlled by the environment. The *environment itself can be a process*. <br>
Bottom line: the entire system can be described, prescribed, analysed as interacting processes.

### Interaction

When two processes are brought together to evolve concurrently, the usual intention is that they will interact with each other.
These interactions are regarded as events that require simultaneous participation of both the processes

> P || Q <br>
> to denote that P and Q interacting in lock-step synchronisation as described above

Two parallel processes can be evaluated as a single process, by identifying possible paths/events between both processes.

### Laws
Notable laws:
 || is symmetric

 || with a deadlocked process results in the entire system being deadlocked

 (c !P)||(c !Q)=(c !(P ||Q

traces(P ||Q)=traces(P) [intersect] traces(Q)

(P||Q)/s=(P/s)||(Q/s)

### Example: limiting number of individuals at the dining table
FOOT0 =(x:D!FOOT1)
FOOTj =(x:D!FOOTj+1|y:U!FOOTj 1) forj2{1,2,3}
FOOT4 = (y : U ! FOOT3)

### Problems
- Infinite overtaking

A constantly greedy neighbour.

- Absence of deadlock

Taking an arbitrary trace and showing that in all cases there is at least one event by which s can be extended and still remain in traces.

### Deterministic (Mathematical models)
- αP - the set of events in which the process is in principle capable of engaging

- traces(P) - the set of all sequences of events in which the process can actually participate in if required

## Non-Determinism

Processes are called deterministic, because whenever there is more than one event possible, the choice between them is determined externally by the environment of the process.
Sometimes a process has a range of possible behaviours, but the environment of the process does not have any ability to influence of even observe the selection between the alternatives.

The choice is made, as it were internally, by the machine itself in an arbitrary and nondeterministic fashion.

Non-determinism arises form a deliberate decision to ignore the factor which influence the selection. Nondeterminism is useful for maintaining a high level of abstraction in descriptions of the behaviour of physical systems and machines.

> P n Q <br>
> is used to denote process that behaves either like P or Q, where the decision is made arbitrarily

α(P u Q) = αP + αQ

### Implementation

The purpose of non-determinism is to abstract away details of implementation. There may be many different implementations of nondeterministic process P each with an observably different pattern of behaviour. The differences arise from different permitted resolutions of nondeterminism.

Needs to ensure that his system works correctly for both choices.

### General choice
 the environment can control which of P and Q will be selected, provided that this control is exercised on the very first action.

## Communication
Communication is an event described by pair c,v where c is the name of the channel on which the communication takes places and v is the value of the message which passes.


> (c!v -> P)=(c.v -> P) <br>
>  initially prepared to input any value x communicable on the channel c, and then behave like P(x)

### Pipes
> P >> Q <br>
> denotes that there is a process that send all inputs from P to Q

### Buffers and protocols
A buffer is a process which outputs on the right exactly the same sequence of messages as it has input from the left, although it is after some delay.



## Sequential Processes

SP is defined as a process that terminated successfully. (Not STOP)



## Shared Resources

## Personal Thoughts

Many ideas here are analogous to ideas in distributed systems. The book provides a strong mathematical/theoretical foundation to understanding concurrent programming as a whole  by introducing expressions and laws that allow us to express and construct concurrent programming systems.

Keys ideas include processes, determinism and non determinism, communication, traces.
