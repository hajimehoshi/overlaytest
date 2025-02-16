This is a test project for [golang/go#71783](https://github.com/golang/go/issues/71783).

The main package does these things:

* Runs a function `Foo` defined in `internal/foo/foo.go`. This function just outputs 'Original' by default.
* Runs `go list -f {{.Dir}} github.com/hajimehoshi/overlaytest/internal/foo` to get the local path for the package.
* Generates a JSON file to add a new file `internal/foo/bar.go` to replace the output message to 'Replaced'.
