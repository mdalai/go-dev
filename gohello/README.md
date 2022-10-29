# Go Hello World
Learn to be Gopher by following theh GO official doc tutorial [Create a Go Module](https://go.dev/doc/tutorial/create-module)

- Go slice is like an array, except that its size change dynamically.
- Go executes init functions automatically at program startup, after global variables have been initialized.
- To get unix time, use `time.Now().UnixNano()`
- Why is seed used for random number generation? don't know.
- To initialize map: `make(map[key-type]value-type)`.
- Use the Go blank identifier _ (underscore) to ignore any returned values.
- Introduced new func Hellos() and kept Hello() as it is to preserve Backward compatibility. 
- Testing can expose bugs as you make changes.
- Go test naming convention, testing package, `go test` cmd.
- The `go test` command executes test functions (whose names begin with Test) in test files (whose names end with _test.go). You can add the -v flag to get verbose output that lists all of the tests and their results.


1. How to use local module?

Edit it to redirect Go tools from its module path to the local directory. 
`go mod edit -replace example.com/greetings=../greetings`
This will update _go.mod_ file add this line "replace example.com/greetings => ../greetings".
Then run `go mod tidy` will add *require* directive in go.mod file along with semantic version.

2. How to set logger?

```go
// Add prefix
log.SetPrefix("greetings >> ")
// Disable logging the time, src, line #
log.SetFlags(0)
```

3. How to exit the program?

One option is that use `log.Fatal()` to exit program.

