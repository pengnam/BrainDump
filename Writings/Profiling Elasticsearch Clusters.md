# Profiling Elasticsearch clusters

## Background:
I spent some time in Shopee (Southeast Asia's largest ecommerce platform) over summer of 2019 and was interning in the search team. One of my projects is to profile elasticsearch clusters (and design them).

This guide aims to come up with a mental framework to approach cluster design.

## Scope
This guide assumes the user to have at least a basic understanding of: elasticsearch (ES), ES partitions, ES cross-cluster search, and the components of a single ES search request (scattering/gathering).

This project does not include ideas of adjusting heap size/file cache size) as from our garbage collection profiles we do not appear to have this problem. It is a separate problem that can be analysed independently from this tuning.

https://www.elastic.co/blog/a-heap-of-trouble is the best guide for memory size tuning.

## Lifecycle of an ES search request
TODO


## Defining Performance
The goal of cluster profiling is get more information about the steps that would help improve its performance.
The first step is to determine and to agree on a set of metrics to optimise the cluster for. These metrics are limited by the elasticsearch API, and by metrics that we can determine from the cluster itself.

We can separate performance to CPU based performance and Off-CPU based performance.

From the analysis in previous section, we can agree that a gain in one type of performance will lead to a decline in the other hence the goal would be to determine the sweet spot that gives a balance of both.

We came up with this tentative list of metrics for each category:

**CPU Based Performance**
1. CPU % of ES Cluster node(< 80% is ideal) 
A higher CPU% is an indication of higher unreliability.

**Off-CPU Based Performance**
1. Search rate per node (lower is generally better)
This metric refers to the search rate on an individual node.
Here, off-CPU performance strongly correlates with the search rate of the cluster. Minimizing search rate per node generally decreases off-CPU time and also decreases time spent “scattering” and “gathering”.
Search latency of requests (lower the better) [Overall performance]
A lower latency is better.
Thread pool rejections (up till 100 is roughly acceptable) [Overall performance]
Thread pool rejections is an indication whether the capacity of processing on a single node is reaching its limits.
The QPS the cluster can handle before suffering a distinct decline in the performance of above [Overall performance]

Note that most of these metrics have similar contributing factors. (For example, CPU and search latency).

These metrics should be evaluated on two types of requests: search requests and indexing requests. Search requests are more important than indexing requests as they involve user experience directly.
Profiling
The first step would be to profile your requests. By understanding the type and nature of the requests, we can understand the general load profile. An example of the importance of this is the general increase in QPS limit from 17000 to 21000 by utilizing the \_routing features of cross-cluster.
Make sure to get real, actual requests.
Parameters:

Fixed Parameters
Having 1 machine per shard (regardless of whether its a master or a replica) will help balance the workload of the search and indexing requests.

Main parameters to tune:
Increasing clusters (cross cluster search).
An increased number of data clusters can be seen as a method to cope with the size of data. (increase number of clusters, decrease data size on each cluster)
With a decrease in the size of data of each individual cluster, then there is a better likelihood that performance can be better managed.
Increasing/decreasing shards
Increasing number of shards will increase throughput (and latency) up to a point before decreasing
Increasing number of shards will decrease shard size, allowing processing to be distributed across more machine nodes
But increasing number of shards will result in more cpu and off-cpu time for “scattering” and “gathering” phases
Increasing/decreasing replicas [least important]
Increasing number of replicas will decrease throughput (and latency) of search requests
Increasing number of replicas will increase throughput (and latency) of indexing requests
