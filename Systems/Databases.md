# Databases

## Introduction

*Database* is a collection of structured data. Captures an abstract
representation of the domain of an application.

Organised as *records* in *tables/relations*, and there are relationships between *records*.
Each *relation* has a set of *attributes*, where each attribute has a type.

Schema is a *structural* description of relations in a database.

## Background

Relational model was proposed in 1970 by Ted Codd. Database abstraction
to avoid maintenance.

- Store database in simple data structures
- Save queries in a high level language

## Relational Model

- structure: definition of relations and concepts

- integrity: ensure that contents satisfy constraints

- manipulation: access and monitor contents

Note: Based on bags (duplicates) and not sets (no duplicates)

## ACID

Set of db transactions to guarantee validity in event of errors/power failures.
A set of db operations that satisfy the ACID properties is called a transaction.


### Atomicity

Each transaction is treated as a single unit, and whether it completely succeeds or fails, the DB is left unchanged. 

### Consistency

Transaction brings from one valid state to another, whetere a valid state means that constraints are satisfied. 

### Isolation

Transactions executed concurrently. Current execution of transaction leaves DB in the same state that would have been optained if it was executed sequentially.

### 

## Fundamental Concepts

1. Data representation

    - Data Model: set of constructs to describe organisation of the data
    (i.e. tables, graphs, triples, relational, key-value)

    - Conceptual Schema: description of particular collection of data using given
     data model. (i.e. schema)

     - Physical Schema: Physical organisation of the data (i.e. on the disk itself)

2. Declarative Querying and Querying Processing

    - Data Independence: separating what from how

    - Data Definition Language vs Data Manipulation Language


3. Transactions

    - Grouping atomic transactions

    - Moving between consistent states

    - Isolates from parallel execution of other actions


## Relational Algebra

Queries operate on relations to produce relations.

Operators:

1. Select Operators

Picks specific rows (by expressing constraints)

2. Project Operator

Picks specific columns (by specific columns)

3. Compose Operators

Compose the rows and columns

Set Operators

1. Union Operators

Combines relations horizontally by using a cross product.
The tuples need to have the same attributes.

2. Difference Operators

Finds difference between sets. The tuples have to have the same attributes.

3. Intersection

Doesn't add any expressive power to the language as it can be expressed with two difference operators. The tuples have to have the same attribute.

4. Product (Cross Join)

Generates a relation that has all possible combinations of tuples from the input relations.

5. Join

Generates a relation that contains all tuples that are combinations of two tuples with a common attribute.

## SQL

1. Aggregations

Returns a single value from a bag of tuples

(i.e. AVG, MIN, MAX, SUM, COUNT)

2. Window Functions

Performs a calculation across a set of tuples related to a single row

(i.e. OVER(PARTITION BY cid))
(i.e. RANK() OVER (PARTITION BY .. ORDER BY grade ASV))

3. 
