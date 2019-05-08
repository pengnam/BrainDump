# Idiomatic Python

## Idioms
|--|--|--|
|Identifier Type| Format |
| Class | CamelCase|
|Variable/Function| words joined by underscore
|Internal functions| prefixed by a single underscore|


## Default Parameters
Using dict.get to provide default values
i.e. instead of:

log_severity = None
configuration.get('severity', log.Info)


## Avoid repeating variable names in the if statement:
i.e. if name in ("x","y","z") instead of if name == "x" ...



## Use long descriptive names when necessary

## Use short mnemonic names when possible

## Use decomposition instead of commments
i.e. Abstract code out into a helper function

## List comprehensions when possible

But not excessively


https://hackernoon.com/going-beyond-the-idiomatic-python-a321b6c6a5e6
https://jeffknupp.com/writing-idiomatic-python-ebook/
