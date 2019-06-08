# dep

dep is  a package manager

## Core Concepts
Dep has a four state system: 
1. Current project source code
2. Manifest
3. Lock (contains the dependency graph)
4. Source code of the dependency itself


## Functional Flow
1. Generate Gopkg.lock from project src and Gopkg.toml
2. Generate src dependency from gopkg.lock



Add introduces new dependencies into the depgraph
- which means a new gopkg lock is generated w new dependency
- or a version constraint is added


Update (confusing)


