# go-error-handlers
A better enum based error handling for go. Inspired from rust's enums for error handling, these enums are for more ergonomic way of handling errors and values that could be nil.

```
// instead of 
val, err := someFunc
if err != nil {
	....

res := someFunc
if res.IsErr() {
	...
}
// use value
res.Value()
```
