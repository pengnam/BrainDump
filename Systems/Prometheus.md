# Prometheus system

## Current system at workplace

1. Prometheus

Metrics collection, storage, calculation

2. M3DB

Efficient storage for prometheus data. Stores time series database. Distribute time series database. 


3. Grafana

Visualization tool and alterting tool


## The system

Components:

ZK stores IP address of each node

Script stores IP addresses into prometheus, 

Fanout Storage
-Reads/writes series data

- Local storage

- Remote storage

 
## M3DB
TSZ: 1.45 bytes/datapoint

All writes persist to a commit log

Sharding and replication makes horizontal scaling easier. 

Provide more fine grained lock granularity.

Data consistency level when writing and reading. 

Cache to speed up read operation



