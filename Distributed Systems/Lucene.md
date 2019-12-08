# Lucene

## Packages

1. analysis: defines an abstract analyzer API for converting text from a *Reader* into a *TokenStream*, an enumeration of token *Attributes*

2. codes: provides an abstraction over the encoding and decoding of the inverted index structure, as well as different implementations that can be chosen depending on application

3. document: provides simple *Document* class. A set of named fields whose names may be strings or instances of *Reader*.

4. index: provides two primary classes: *IndexWriter* which creates and adds documents to incides

5. search: provides data structures to represent queries (*TermQuery* for individual words, *PhraseQuery* for phrases)

6. store: abstract class to store persistent data

7. util: contains handy data structures and util classes

## Basic Concepts
### Index 

Similar to a database table, but less constraints. There are multiple segments of an index. 

Document: similar to a row

Field: smallest definable unit of a data index. Data index provides many types of fields.

Term: Term is produced when field is put through an analyzer.

Segment: Index is composed of 1 or more sub-indexes. 

First writes to an in-memory buffer (buffer not readable). Once buffer is large enough, it flushes to segment. Each segment has own index, independently searchable but not changed. 

A DocId is not actually unique to the Index but is unique to a Segment. Lucene does this mainly to optimize writing and compression. Since it is only unique to a Segment, how can a Doc be uniquely identified at the Index level? The solution is simple. The segments are ordered. To take a simple example, an Index has two segments and each segment has 100 docs respectively. The DocId's in the Segment are 0-100 but when they are converted to the Index level, the range of the DocId's in the second Segment is converted to 100-200.
DocId's are unique within a Segment, numbered progressively from zero. But this does not mean that the DocId's are continuous. When a Doc is deleted, there is a gap.

### Index type 

Lucene supports many types of fields and each type of field determines the supported data types and index modes. 



## Codecs

Provides abstraction over the encoding and the decoding of the interveted index structure, as well as different implementations that can be chosen depending on application

Whenver Lucene needs to access the index, it does the accessing through codec APIs. 


Codec is set-up per segment and every segment is free to use a different codec, although that is uncommon. 

### Format

1. StoredFieldsFormat: encodes each document's stored fields (block compression)

2. TermVectorsFormat: encodes per-document term vectors. (block compression)

3. LiveDocsFormat: bit set to mark deletion. 

4. NormsFormat: encodes per-document and per-indexed-field index-time score normalizatino factors as encoded by the similarity. Contains compact encoding of field's length and field level boosting

5. SegmentInfoFormat: saves metadata about the segment. How many docs, unique id. 

6. FieldsInfosFormat: per-field metadata, whether nad how it was indexed, stored.

7. PostingsFormat: covers the inverted index, including fields, terms, documents. 

8. DocValuesFormat: column stride per document values. 

9. CompoundFormat: collapses separate index files, written by the above formulas into two files per segment. 

## Commits

Data is **only** searchable after a commit. Commit is a two stage operation. The first is prepareCommit, commit is called to complete a step. 


## Concurrency Model

Core interfaces are thread-safe and have internal concurrency tuning to optimize multi-thread writing. 


1. Multiple threads concurrently invoke an IndexWriter interface and IndexWriter's specific internal request is executed by DocumentsWriter. Before DocumentsWriter internally processes the request, it allocates DocumentsWriterPerThread based on the thread currently executing the operation.
2. Every thread processes data in its own independent DocumentsWriterPerThread space, including tokenizing, relevance scoring and indexing.
3. After data processing is finished, some post-processing is carried out at the DocumentsWriter level, such as triggering a FlushPolicy decision.


First and third steps are locked.

Each DWPT includes its own in-memory buffer, which flushes into an independent file/segment. 

## Deleting

Deleting a document goes through a deletion queue. Each DWPT has its own independent deletion queue known as pending updates. DWPT pending updates and global deletion queue schnc bi-directionally because doc deletion is global in scope. 

## Indexing Chain

The key concept in internal indexing is indexingChain, which as 

## Compression Techniques

1. Bitpacking: first bit is a continuation bit. Used for postings lists, numeric doc values. 

2. LZ4: Used to stored fields. Stored fields take about 70% of the memory.

3. FSTs: Map<string,?>, prefix suffix, terms index

## What happens when running a term query

1. Lookup terms in terms index (Uses FST, in memory). Only stores prefixes of terms. Something like a prefix tree. As traversing sum up the weights. The sum is the offsets in the term dictionary. Doesn't load alot of things in memory. 

2. Term dictinary: compressed based on shared prefixes, similarly to a burst trie called the "BlockTree terms dict". Read sequentially till the term is found. Contains information about frequency, score (I think)

3. Postings lists:
block encoding. Encoded using Frame of Reference (FOR) delta

4. Stored Fields

Find the stored fields per document of the block. 



## Disk Seaks

2 disk seeks per field for search:
find position in term dict + find postings list


1 disk seek per doc for stored fields:
Finding the stored values


Typically all in cache


## Bottlenecks:

First bottle neck is when index grows larger than filesystem cache. Term dict/postings list not fully in cache. 


