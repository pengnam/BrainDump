# WebPack

## Background
Static module bundler for modern JavaScript applications.
Internally, it builds a dependency graph that maps every module in project needs and generates one or more bundles.

Write modules that can run in the browser

## Core Concepts
1. Entry
2. Output
3. Loaders
4. Plugins
5. Mode
6. Browser Compatibility

ECMAScript modules vs CommonJS modules

## WebPack Ideas

Use javascript parser(babylon) to parse javascript code into an AST.
Traverse AST to find the modules that a given module depends on, assign given module a unique value. This abstraction of dependencies is an 'asset'.

Combining these assets will allow us to form a dependency graph.

This dependency graph can be used to write a new require function: one that finds the right module to "require" from.


https://www.youtube.com/watch?v=Gc9-7PBqOC8
