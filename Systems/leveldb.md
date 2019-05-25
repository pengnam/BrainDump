# Level DB

Description:

Fast key-value storage library, provides ordered mapping from string keys to string values. 

Keys and values are arbitrary byte arrays. 

Data stored sorted by key. 

Supports a few basic operations:
1. Put(key, value)
2. Get(key)
3. Delete(key)

Stores keys and values in arbitrary byte arrays, and data is sorted by key. Supports batching writes, forward and backward integration. 

## Method of storage
Sequence of key value pairs in file are stored in sorted ordered and partitioned into a sequence of data blocks. The data blocks are formatted according to the code in block builder, before being optionally compressed.

## General Structure

General types of blocks:
1. Data blocks
2. Meta blocks
	- Stats meta blocks
	- Filter meta blocks
3. Index blocks
	- One entry per data block

A bunch of data blocks, a bunch a meta blocks. 

The metablocks describe one entry for every other meta block where the key is the name of the meta block and the value is a BlockHandle pointing to meta block. 

## 

