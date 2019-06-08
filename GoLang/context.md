# Context


In servers, each incoming request is handled in its own goroutine. Request handlers often start additional goroutines to access backends such as databases and RPC services. 

The set of gortouines need request specific values. Vontext packages makes it asy to pass request-scoped values, cancellation signals, deadlines to all goroutines. 


```
// A Context carries a deadline, cancelation signal, and request-scoped values
// across API boundaries. Its methods are safe for simultaneous use by multiple
// goroutines.
type Context interface {
    // Done returns a channel that is closed when this Context is canceled
    // or times out.
    Done() <-chan struct{}

    // Err indicates why this context was canceled, after the Done channel
    // is closed.
    Err() error

    // Deadline returns the time when this Context will be canceled, if any.
    Deadline() (deadline time.Time, ok bool)

    // Value returns the value associated with key or nil if none.
    Value(key interface{}) interface{}
}
```

Done method returns a channel that acts as a cancellation signal on functions that run on behalf of context. When channel is closed, functions should abandon its work and return. 

## Derived Contexts
Functions used to derive new context values from existing ones. 

WithCancel returns a copy of parent whose Done channel is closed as soon as parent.Done is closed.  CancelFunc cancels a context.
