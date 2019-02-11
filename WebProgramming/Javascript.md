# Javascript

## Understand this
### Credits
https://yehudakatz.com/2011/08/11/understanding-javascript-function-invocation-and-this/
### Core Primitives
Core primitive is a functions call method. <br>
The first argument is the thisValue of the method,

```
function hello(thing) {
    console.log(this + "says" + thing);
}

hello.call("Sean", "How cool is this not")
```

The short version is: a function invocation like `fn(...args)` is the same as `fn.call(window [ES5-strict: undefined], ...args)`.

**
A function does not necessarily have a consistent notion of `this`.
**


```
var bind = function(func, thisValue) {
  return function() {
    return func.apply(thisValue, arguments);
  }
}
```
