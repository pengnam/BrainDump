# Elastic Search

Elasticsearch is a distributed RESTful search engine. 
- Search engine where each index has a fully configurable number of shards. 
- Each shard has one or more replicas
- Read/search operation in replica


Built on top of Lucene


Multi Tenant
- Support for > 1 index
- Index level configuration

Set of APIs
- HTTP RESTful API
- Native Java API

Document oriented
- No need for upfront schema definition
- Schema can be defined for customizing indexing process



## Lucene
High performance, full-featured text search engine. 
150GB/hour with small RAM requirements (1MB heap)
About 20%/30% of size

Implements efficient search algorithms, available open source under Apache license

## REading Documents
### Data replication model
EAch index is divided into shards, and each shard has multiple copies. Each copy is a replacation group and kepy in sync. 

Data replication is based on the primary-backup model, and described well in PacificA paper of microsoft research. Single copy of replication group that acts as primary shard, the other copies are replication shards. 

Primary is responsible for replicating operation to other copies. 

## Basic write model
Every indexing is resolved using routing, typically based on document ID. 


## Routing

Routing is controlled using the hash of document ID's value.

For explicit control, value fed into the hash function used by router can be directly specified on a per-operation basis. 

When setting up explicit mapping, _routing field can be optionally used to direct the index operation. 

## Basic Write Model
Every indexing operation is resolved to a replication group using routing. (Refer to above). Once replicatino group determined, the operation is forwarded to the primary shard. Primary shard validates the operation and forwards it to the replcias. 

List is called in-sync copies and is maintained by the master node. 

The primary shard follows this basic flow:

Validate incoming operation and reject it if structurally invalid (Example: have an object field where a number is expected)
Execute the operation locally i.e. indexing or deleting the relevant document. This will also validate the content of fields and reject if needed (Example: a keyword value is too long for indexing in Lucene).
Forward the operation to each replica in the current in-sync copies set. If there are multiple replicas, this is done in parallel.
Once all replicas have successfully performed the operation and responded to the primary, the primary acknowledges the successful completion of the request to the client.

## Write failuer
In the case that the primary itself fails, node hosting primaru will send message to master. Indexing operation will wait for master to promote one of the copies to be primary. 

## Basic Read Model
When a read request is received by a node, that node is responsible for forwarding it to the nodes that hold the relevant shards, collating the responses, and responding to the client. We call that node the coordinating node for that request. The basic flow is as follows:

Resolve the read requests to the relevant shards. Note that since most searches will be sent to one or more indices, they typically need to read from multiple shards, each representing a different subset of the data.
Select an active copy of each relevant shard, from the shard replication group. This can be either the primary or a replica. By default, Elasticsearch will simply round robin between the shard copies.
Send shard level read requests to the selected copies.
Combine the results and respond. Note that in the case of get by ID look up, only one shard is relevant and this step can be skipped.


## Distributed
### Wait for active shards

Indexing oeprations can be configured to wait for a certain nunmber of active shard copies before proceeding with the operation. If the requisite number of active shard copies are not available, then the write operation must wait and retry. 

Write operations can work on the primary but fail on the replicated nodes. 

## Refresh

Refresh controls when changes made by the request are made available to search. 

## API

## Search Shars API
Returns the indices and shareds that a search request would be executed against. 

Sequence number and primary term uniquely identify a change. 

preference: which shard replicas to execute the search request on. (By default, randomized, between requests)

local: boolean value to indicate whether to read state locally or use master nodes cluster state.


cross-field queries: term centric approach by blending idfs, searches for matches across all fields

BoolQueries:
A query that matches documents matching boolean combinations of other queries. 

Built using 1 or more boolean clauses, which each clause mapping to a typed occurence.


