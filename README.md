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

Examples
--------
### Result
```
	// returning a string value
	res := Result.Value("hi")
	// int value
	resInt := Result.Value(1)
	// but if you are sending an error
	// you may need specify the type in some cases
	res := Result.Err[any](err)
```
### Option
```
	// returning a string value
	res := Option.Some("hi")
	// int value
	resInt := Option.Some(1)
	// but if you are sending an None value
	// you may need specify the type in some cases
	res := Option.None[any]()
```