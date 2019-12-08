# Query Rewrite for Null and Low Search Results in eCommerce

This document summarizes the key points in a paper submitted to SIGIR 2017. 
http://sigir-ecom.weebly.com/uploads/1/0/2/9/102947274/paper_8.pdf

## Introduction

Goal is to design a system that reformulates query into multiple alternative queries that enlarge recall without altering main search intent. 

## Approach

The approach adopted by the paper can be broken down into two parts:

1. For verbose queries, perform POS tagging, entities, unit of measurements. 

2. For non-verbose queries, reformulate queries by replacing some tokens. Avoids turning specific query into broad query. Done using language model approach. 


## Query Term Dropping

Selection of terms drop is dropped by predetermined rules. There are three categories. Tier 1 includes brands, product names, first noun phrases. Tier 2 are attributes of the core product. Tier 3 gropup are stop words, conjunctions. 


## Identifying Brands

Subproblem is the disambiguation of brands.

*P(B|q) = \sum P(C_i|q) P(B|C_i)*

Predicts probability that B is brand in context of query q.

C_i is the shopping category. 

## Scoring variations

Form variations using multiple passes, which is similar to the system I currently implemented that adopts a pseudo- BFS approach. 

Suggestions from earlier pass will always rank higher than the ones from the later passes, as overall queries with more original tokens are believed to be closer to the original user content. 

The paper's algorithm also prefers suggestions with more brands or nouns over suggestions with more adjectives. 

Baseline tagging scoe is given by a step function over their tags. 54% of tokens are nouns. There is hence a need to rank the nouns. 

Ranking nouns can be done by using the cross-entropy of clicking probabilities conditioning on their document frequencies. 

*P_click(r)  = -P(click|r) log(1 - P(click|r))*

r is the document frequency of the token. P(click|r) is evaluated by a large sample of null queries. P(click | r) is evaluated by a large number of null queries. 


Clicking score added to tagging score using discounted cumulative gain.


## Query Term Replacement


Consider all reassignments of tokens of the n-gram model. To calculate the probability that one token is replaced by another token, 
