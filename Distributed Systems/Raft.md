


# Raft

- Raft is a concensus algorithm that is an alternative to Paxos. 
- It offers a generic way to distributed a state machine across a cluster of computing systems, ensuring that each node in the cluster agrees upon the same series of state transactions. 
- Reliable, Replicated, Redundant, and Fault Tolerant. 


## Subproblems
1. Leader election
2. Log replication

## Leader election
Roles:
- Leader: onyl one leader
- Candidate: leader candidates in the election process
- Follower: Do not send any request, only accespt heartbeat from laeader, except forwarding requests from client to leader

### Changing of roles
1. Follower: When meets heartbeat timeout. convert 

