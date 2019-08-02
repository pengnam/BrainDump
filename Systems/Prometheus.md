# Prometheus system

## Fundamental Concepts

### Data Model

Prometheus stores all data as time series: streams of timestamped values belonging to the same metric and the same set of labeled dimensions. It also generates derived time series from these data.

### Jobs and instances

An endpoint that you can scrape is an instance. A collection of instances with the same purpose, a process replicated for scalability or reliability, is called a job. 

When prometheus scrapes a target, it attaches some labels automatically to scraped time series which serve to identify the scraped target. 

Prometheus stores a sample in the time series. 


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



