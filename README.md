# Golang Mocking Imports

This is a demo repo that includes all the code I wrote in [this blog post](https://colinj.hashnode.dev/mocking-imported-libraries-in-golang-unit-tests).

Feel free to run it to ensure that it works. You can also run the tests and recreate the coverage report to prove that everything works as advertised:
```shell
$ go test ./... -coverprofile=coverage.out  
$ go tool cover -html coverage.out -o coverage.html
```