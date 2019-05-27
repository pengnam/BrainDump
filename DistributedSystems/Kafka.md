# Kafka

## Introduction

Distributed streaming platform that can:
1. Publish and subscribe to streams of records, similar to a mesage queue. 
2. Store streams of records in a fault tolerate and durable way. 
3. Process streams as they occur. 

Used for:
1. Building real-time streaming applications that reliably get data.
2. Building real-time streaming applications that transform data.

## Concepts

Kafka is run as a cluster on one or more servers that can span multiple data centers. 
Kafka cluster stores streams of records in categories called topics 


## Distribution 
Each cluster handle a share of the partitions. Each partition is replicated across a number of servers. Each partition has 1 leader and other followers. 

Leader handles all read and write request, followers replicate. If leader fails, one of the followers will take over.  

### Topics and Logs
Topic is a category to which records are published. Topics in Kafka are always multi-subscriber: i.e. a topic can have zero, one, or many consumers. 


## Kafka as a messaging system
1. Queueing
Allows us to divide the processing over multiple consumer instances, which lets us scale processing. 

2. Publish-subscribe
record is broadcasted to all consumers

Kafka is a generalized notion of stream processing that subsumes batch processing as well as message-driven applications. 



