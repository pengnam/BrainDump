# Log-Structured Merge Tree

Data structure makes it attractive for providing indexed access to files with high insert volume, such as transactional log data. LSM keeps data is two or more separate structures. 

## Description

Two tree like structures: C0, C1, C0 in memory and smaller, C1 in disk. New records are inserted into C0. Once C0 is too large, merged with C1 on disk. Migrated to disk in a merge-sort like way. 

## How it works:

1. New write added to an in-memory sorted balanced tree (known as memtable)

2. When memtable is too big, it is flushed to disk as an SSTable file. SSTable is a sorted key-value storage on disk of the in-memory tree. 

3. Reads are directed to memtable. If keys are not found in memtable, it is searched and most recent SSTable is retrieved. 

4. Duplicate key removed from SSTable and most recent is retained (compaction). Compacted SSTable are merged into new SSTable (merge sort -esque algorithm)


