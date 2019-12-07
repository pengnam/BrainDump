# Javascript

Content was largely take from "You don't know JS".

## Pillars

### Scope and Closure

Organization of variables into units of scope is one of the most foundational characteristics of any language. 



- Lexically scoped variables (Static, deduced at compile time)

The scope unit boundaries is determined at the time the program is parsed. Author time descision. 
Variables at that level of scope nesting are accessible, variables outside are not. 

- *Hoisting*: All variables declared anywhere in scope are treated as they are declared at the beginning og the scope. 

- **var**: Declared variables are function scoped even if they appear inside a block. 

### Prototypes

Prototypal inheritance was initially used to implement the class design pattern. 

However, after `class` came in ES6, language doubled down on its inclination to OO.

But behaviour delegation is still important for organising behaviour and data in the program.

### Types and Coercision

Not just solved with type aware tooling. 


