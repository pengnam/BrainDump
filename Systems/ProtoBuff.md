# Google Protobuff

## Description

Protocol buffers: language-neutral, platform-neutral, extensible way of serializing structured data for use in communications procols, data storage, and more. 



- Name-value pairs
- One or more unique numbered pairs
- Specify optional fields, required fields, repeated fields. 

Compiler generates data access classes, which provides simple accessors for each field as well as methods to serialize/parse teh whole structure to/from raw bytes. 

Running the compiler will generate class called persons. 


## General problems that protocol solves

Deals with index server request/response protocol, where hand marshalling/unmarshalling has to be done. 

- Generates serialization and deserialization code for hand parsing.
- Being used for short-lived RPC 


## General Syntax

### Data Types
- bool
- int32
- float
- double
- string

### General Structure

message NameOfStructure {
  dataType nameField = val;
}

dataType: 1 of above data types
nameField: name of the field
val: markers on each element that uniquely identify the unique tag that field uses in the binary encoding



