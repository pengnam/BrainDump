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


## Keys
Kafka accepts the use of keys. Keys are useful if a strong order for a key is required, or if we require that messages with the same key go to the same parition. 

## Cluster
Cluster typically contains multiple brokers to maintain load balance. Brokers are statemless, and use zookeeper for maintaining cluster state. 
One broker can handle hundreds and thousands of reads and writes per second, with TB of messages. 
ZooKeeper coordinates the brokers 

## Principles
1. Contains serveral producers, several brokers, several consumers
2. Partition is a physical concept, topic is a logical concept. 
- Each partition append log file at the storage level. 
3. Location of each message is called an offset, offset is a long type number that uniquely identifies a message. 
4. Kafka's high reliability guaranteed by its high replication. 
5. Setting partition allow all messages to be evenly distributed to the different partitions, achieving a level of expansion. (Note that writes can be to the tail end of any of the partitions).
6. When creating a topic, one can specify the key of the mssage to dertermine the producer send to which partition. 
7. Each partition uses segments, to ensure that the size of the log can easily expand/contract. 
8. Each segment consists of a .Index file and a .log file. Index stores metadata.

## Consumer Postion

Kafka maintains information about messages sent and consumed. (For the instances where messages are sent, package lost, and the consumer ends up time-out or waiting forever). 

This creates a problem of double consumption, or broker must keep multiple states about each message. 

Kafka handles this by ensuring that each partition is consumed by one consumer group at a single moment in time. 

## Client-side assignment protocol

New consumer relies on a server side coordinator to negotioate the set of consumer processes that form the group

https://cwiki.apache.org/confluence/display/KAFKA/Kafka+Client-side+Assignment+Proposal

## Balancing Load
Kafka balances the node at 3 different locations
1. Assignment strategy
(Look at partition assignment strategy)
2. Producer level, strategy for selecting partition to store message. 
3. If consumer processes message for a long time, Kafka thinks partitino is dead and reassign partitions among other consumers. Once the job is done, partitions can be assigned to it again.

## Replication and Synchronization
Kafka has a data copy algorithm to ensure that Aif leader fails or hangs up, a new leader is elected and message is written. 

In N replicas, 1 replica is the leader replica. Leader replica handles all reads and writes to the partition. Follower passively and regularly copies data from the leader.
[image](https://pocket-image-cache.com/direct?url=https%3A%2F%2Fcdn-images-1.medium.com%2Fmax%2F1600%2F0*EKINPKA_r9-9ip5k.png&resize=w1408)
Leader is responsible for maintaining and tracking the status of follower lags in the ISR (In-Sync Replicas) which is a copy sync queue. When producer sends a message to the broker, leader writes message and copies it to all followers. 

Message replicate latency limited by the slowest follower, and it is important to detect slow copies quickly. 

Messages in partition are totally ordered but not between partitions.

## Understanding parameters

request.required.acks:
1(default): producer sends message after leader successfully received the data in the ISR and is confirmed. If leader is down, it will lose data
0: producer does not need to wait for confirmation from broker to continue sending the next batch of messages. 
-1: producer wait for all followers to confirm data sent, but not guarantee data will not be lost. 


## Notes about number of partitions

We don't increase the number of partitions to be too large because on failure, the reelectino for broker for the partition Suppose that a broker has a total of 2000 partitions, each with 2 replicas. Roughly, this broker will be the leader for about 1000 partitions. 
When this broker fails uncleanly, all those 1000 partitions become unavailable at exactly the same time. Suppose that it takes 5 ms to elect a new leader for a single partition. It will take up to 5 seconds to elect the new leader for all 1000 partitions.

## TO READ:
Quorums

Kafka server vs zoo keeper in new kafka version
