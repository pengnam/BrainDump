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


## Key concepts

1. Near Realtime

There is about a 1 sec latency from indexing to searching

2. Cluster

Collection of one or more ndoes that holds data, provides federated indexing and search capabilities across all nodes. 
Cluster identified by unique name, and node joins a cluster only through name

3. Node

Single server in cluster, stores data, participates in indexing and search. Node identified by random UUID. 

4. Index

Index is a collection of docs w similar characteries, identified by name

5. Type

Logical category/partition of index to allow one to store diff type of documents in the index (to be deprecated)

6. Document 

Expressed as JSON

7. Shards and replicas

Sharding allows one to horizontally split/scale content. Allows one to distribute and parallelize operations across shards. When you create an index, one can simply define the number of shards that one wants. 

Each shard is a fully functional index.

## Lucene
High performance, full-featured text search engine. 

150GB/hour with small RAM requirements (1MB heap)

About 20%/30% of size

Implements efficient search algorithms, available open source under Apache license

## Reading Documents
### Data replication model
Each index is divided into shards, and each shard has multiple copies. Each copy is a replacation group and kept in sync. 

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

## Search API

https://www.elastic.co/guide/en/elasticsearch/reference/current/getting-started-search-API.html

## Search Shard API
Returns the indices and shards that a search request would be executed against. 

Sequence number and primary term uniquely identify a change. 

preference: which shard replicas to execute the search request on. (By default, randomized, between requests)

local: boolean value to indicate whether to read state locally or use master nodes cluster state.


cross-field queries: term centric approach by blending idfs, searches for matches across all fields

BoolQueries:
A query that matches documents matching boolean combinations of other queries. 

Built using 1 or more boolean clauses, which each clause mapping to a typed occurence.


## A note about Indexing

An elasticsearch cluster can contain multiple *indices*, which in turn contain multiple *types*. These types hold multiple documents, and each document has multiple *fields*. 


## Index lifecycle

There are four stages in the index life cycle.

1. Hot - index being updated and queried
2. Warm - index no longer being updated, still being queried
3. Cold - no longer being updated and is seldom queried. 
4. Delete - index is no longer needed and can be safely deleted


Lifecycle policy governs index transitions through these stages:
1. Maximum size or age at which we want to roll over to new index
2. Point that index is no longer being udpdated
3. When to force merge, permanently delete documents
4. Index deletion

## Document Orientated

Objects are complete data structures that may contain dates, geo locations, other objects, arrays

These objects are stored and indexed to be searchable. JSON is used as all languages support it. 

Converting object to JSON to be indexed is simpler than converting JSON to flat table. 

Document has metadata. The three required metadata:

1. _index: where the document lives

2. _type: class of object that document represents

3. _index: collection of documents that should be grouped together for a common reason

## Mapping

Mapping is the process of defining how a document, and the fields in contains, are stored and indexed. For instance, use mappings to define:

1. Which string fields should be treated as full text fields
2. Which fields contains numbers, dates, geolocations
4. Format of date values
5. Custom rules to control mapping for dynamically added fields

Each index has a mapping type to determine how the document will be indexed. 

Mapping type has:
	- Meta-fields: used to customized how a document's meta data asscociated and treated
	
	- Fields or properties: Contains list of fields or properties pertinent to the document

Settings to prevent mappings explosion:

Too many fields in an index is a condition that can lead to mapping explosion, which can cause out of memory errors and difficult situations to recover from. This problem may be more common than expected. Consider where every new document inserted introduces new fields. 







## Creating and Updating a document

When sending a PUT request and if a document with the same _index, _type, _id exists, 
	1. If ?op_type=create is present, then  the request will be rejected with a *409 Conflict* response code
	2. If it is not present, new document will be created with a higher version number, older document is rejected 

We are also able to partially update the document using _update. Scripting using Groovy is also supported.

## Concurrency Control

The most recent request wins in ES.
Preventing race condition when updating values in index.

### Pessimistic Concurrency Control
Used in relational databases, assumes conflicting changes are likely to happen. Blocl access to a resource in order ot prevent conflicts. i.e. locking a row before reading data, ensuring only the thread that place the lock is able to make changes. 

### Optimistic Concurrency Control

Used in ES, assumes conflicts are unliekely to happen, doesn't block.
However, if underlying data has been modified between reading and writing, the update will fail. Then up to application to decide how to resolve the conflict. 


ES is async and concurrent, meaning that requests may arrive at destination *out of sequence*. Versions change at every document change, and ES uses version to ensure changes are applied in the correct order. i.e. if an old document arrives after new, it can be ignored. 

*This requires the user to retrieve a copy if the document to get its version number, before reindexing the document with the version number of the document that it shuold ne applied to*.

Note that an external data store can be used. By specifiying version_type=external parameter, ES checks that the current version is less than specified version. 

## Distributed Document Store

When a document is indexed, it is stored on a single primary shard.

Routing can be used direct requests to the right shard. 


### Store

Allows one to control how index data is stored and accessed on disk

*fs*:
	Default file system implementation. Pick best implementation depending on the OS. 

*simplefs*: 
	Straightforward implementation of file system storage

*niofs*: 
	Stores the shard index on the file system



## Elasticsearch Cluster

- Master node: creates, deletes indices. Adds, removes nodes. On cluster state change, master node broadcasts the changes to other nodes in the cluster. Only 1 master node in cluster.

- Data node: holds data in shards, performs CRUD, search, aggregations. If 1 data node stops, cluster reorganizes, and continues operation. 

- Client node: routes cluster related requests to the master node, data-related requests to the data nodes. 

- Tribe node: talk to multiple clusters to perform search and other operations. 

- Ingest node: pre-processing documents before actual indexing.


### Adding node to cluster

When starting a node, node pings all nodes in the cluster to find master node. Once master is found, it will ask master to join by sending a join request. Master accepts it as new node of cluster, notify all nodes in cluster about presence of new node. 

If joined node is data node, master will reallocate data evenly across nodes. 


### Remove node

If we stop a node that is unresponsive in specific amount of time, master node will remove it from cluster and reallocate the data. 


### Elasticsearch Basics

1. Elasticsearch is a search engine (not a datastore)

2. Horizontal scaling: we can build a cluster with an infinity of hosts, depending on needs and bottlenecks. Running dataset on alot of small machines better than using a few large hosts. 

### Hardware

- CPU

CPU is needed for running complex filtered queries, intensive indexing, word analyses. The right CPU can drastically change the performance. 

ES breaks the CPU use into threadpools:

	1. Generic: for standard operations such as discovery

	2. Index: for indexing

	3. Get: for get operations

	4. Bulk: for bulk operations such as bulk indexing

	5. Percolate: for percolation


Each pool runs a number of threads, that are allocated to the pools.

- Memory

ES runs on Java, Java is garbage collected. Memory divided into 2 parts: Java heap space and everything else. (ES does not rely on Java only). More memory given to the heap, more time Java spends garbage collecting. 

CMS runs multiple concurrent threads to scan the heap for objects that can be recycled. "Stop the world" 

Java 8 brings a brand new garbage collector called Garbage First designed for heaps greater than 4GB. G1 uses background threads to divide the heap into regions, and scans the regions that contai the most garbage objects first. 

For operational reasons, this is good. However, the new JVM might lead to data corruption. 

ES has multiple buffers to perform in memory operations, as well as chache to store query results. 

	1. Indexing buffer: buffer data during indexing

	2. Buffer pools


More memory better outside heap. Off heap memory manage threads and for filesystem to cache data. 

File system storage affects performance. NIO FS and MMAPFS. NIOFS lets the kernel manage the file system cache instead of relying on the rboken, oom generator mmapfs. 

You can also commit the exact amount of memory I want ot allocate the heap at startup. Prevents the node from swapping when trying to allocate memory. 

- Network

ES performs alot of network consuming operations from transferring data during queries to reallocating shards. 

ES generally has low network related settings, and we should consider raising the bar. 

- Sharding

Shards are allocated logical parts that can be allocated on all the cluster data nodes. Sharding defined when index created. Only way to change number of shards is to delete indices, create them, reindex. Resize ES in production. 

Don't create more shards than needed, especially if ES routing parameter is used. There will be alot of empty shards. 

Alot of shards on huge indices is good for large cluster (20 data nodes or more). Multiple shards allow for better allocation. Small shares make cluster recovery faster when losing a data node or shutting down cluster. 


### ElasticSearch cross cluster replication

Replication

Replication managed at index level, performed at shard level. Follow index is configured to have identical number of shards as leader index. 


### Remote Clusters

Remote clusters enable unidirectional connections to remote cluster. Configuring remote cluster, connecting only to a limited number of nodes. Each remote clsuter referenced by aa name and a list of seed nodes. When remote cluster registered, clsuter state is retrieved, up to 3 gateway nodes are selected to be connected as part of remote cluster requests. All 

## Requests

Requests have two phases:

1. Scattering phase

Forwards requests to data nodes that holds the data. Each data node executes the request locally and returns results to the coordinating node


2. Gathering phase

coordinating node reduces each data node's results to a single global ruleset

coordinating node has a priority queue to sort results globally

sorted by relevant


## Concensus

Zen discovery module only has two parts:

ping: process to discover each other

unicast: list of hostnames to control which nodes to ping

ES is a P2P systen where all nodes communicates with each other and there is one active master tht controls and updates cluster wide state and operations. 


## Translog

Flush is performed every 30mins, hence its possible to lose commits in buffer. ES gets around that by using a translog, and translog is `fsync`ed after every operation.


## 
