# Designing Data Intensive Applications

## Replication

- Keeping a copy of the same data on multiple machines that are connected via a network. 

### Reasons

1. Keep data geographically close to users

2. Allow the system to continue working even if some of the parts have failed (to increase availability)

3. Scale out the number of machines that can read queries



#### Algorithms to handle changes

Maintaining consistency among the shards

Considerations: 

1. Synchronous or async replication

2. How to handle failed replicas

3. Replication lag


#### Leaders and followers


##### Leader-based replication 

One of the replicas is designated as the leader. When clients want to write to the database, they must send their requests to the leader, which first writes the new data to its local storage. 

The other replicas are known as followers. whenever the leader writes new data to its local storage, it also sends the data change to all its followers as part of a replication log / change stream.

##### Synchronous versus asynchronous replication

- Synchronous replication: leader waits until follower has confirmed that it received a write before reporting success to the user and making it visible to other clients. 

- Asynchronous replication: leader sends a replication but does not wait for a response. 

Synchronocity:

Follower is guaranteed to have an up-to-date copy of data that is consistent withe hte leader

If synchronous follower doesn't response, write cannot be processed. 

Hence, proposal of semi-synchrocity. 


##### Setting up new followers

Clients are constantly writing to database, and data is always in flux so standard file copy would see different parts of the database at different points in time. 

1. Lock database and copy, but it goes against idea of high availability. 

2. Take consistent snapshot of leader's database at some point of time. Copy snapshot. 

#### Multi-Leader Replication

One leader, and all writes must go through it. 


