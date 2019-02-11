# WebPack

## Background
Static module bundler for modern JavaScript applications.
Internally, it builds a dependency graph that maps every module in project needs and generates one or more bundles.

Write modules that can run in the browser



## WebPack Ideas

Use javascript parser(babylon) to parse javascript code into an AST.
Traverse AST to find the modules that a given module depends on, assign given module a unique value. This abstraction of dependencies is an 'asset'.

Combining these assets will allow us to form a dependency graph.

This dependency graph can be used to write a new require function: one that finds the right module to "require" from.g


https://www.youtube.com/watch?v=Gc9-7PBqOC8

## Core Concepts
1. Entry
2. Output
3. Loaders
4. Plugins
5. Mode
6. Browser Compatibility

ECMAScript modules vs CommonJS modules

### Entry
Indicates the start point to build the internal dependency graph from.

### Output
Where to emit bundles (default: ./dist/main.js for main output file and ./dist for any other generated file)

### Loaders
Webpack only understands javascript/json. Hence, loaders are needed to process other types of files and convert them into valid modules that can be consumed by application.

### Sample Code
<code>
module.exports = {
    modern
}
</code>
