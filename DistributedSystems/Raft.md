


Check thesecretlivesof.data.com/raft


PAxos is a family of protocols. 

Check lamport's website. 



# Paxos

Can be decomposed to relatively simple subproblems. 

1. correctness: under non-byzantine conditions
2. available: Any cluster of 5 can tolerate a cluster on 2
3. Does not depend on timing


## Subproblems
1. Leader election
2. Log replication
3. Safety

## Leader election
Roles:
- Leader: onyl one leader
Candidate: leader candidates in the election process
- Follower: Do not send any request, only accespt heartbeat from laeader, except forwarding requests from client to leader

### Changing of roles
1. Follower: When meets heartbeat timeout. convert 

