# OpenTracing

## Data Model
Traces in open tracing are defined by their Spans. Edges between Spans are called references. 

## Span
A span has an operation name, start timestamp, finish time stamp. Each span has zero or more key:value Span tags, with each key as a string. 


## References between Spans

A span may reference zero or more other SpanContexts that are causally related. 

Two types of references:
1. ChildOf
2. FollowsFrom

Both model direct relationships between childSpan and parentSpan. 


## Tracer

Tracer interface creates Spans understands how to Inject and Extract them across process boundaries. 

Operation name: string that uniquely identifies a class of `Span` instances. 
Inject a `SpanContext` into a carrier
Extract a `SpanContext` from a carrier

## SpanContext
SpanContext is immutable, and semantically the caller should be able to ite
