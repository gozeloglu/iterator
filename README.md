# iterator

## Install

You can install the package as below:

```shell
go get github.com/gozeloglu/iterator
```

## Usage

Firstly, you need to create an iter object with `New` function. You can pass elements to that function.

```go
it := iterator.New(1,2,3,4)
```

Then, you can call the methods by using `it` object. 

```go
package main

import "github.com/gozeloglu/iterator"

func main() {
	it := iterator.New(1,2,3,4)
	it.Next()   // Check whether there is element in next
	it.Value()  // Retrieve next element if exists
	it.Prev()   // Check whether there is element in prev
	it.ToSlice()    // Retrieve all elements in slice
	it.First()  // Updates the cursor by moving first index
	it.Last()   // Updates the cursor by moving last index
}
```

## Running tests

You can run the test with following command.

```shell
go test
go test -v  # verbose output
```

You can also get code coverage with the following command.

```shell
go test -cover
go test -cover -v # verbose output
```

## LICENSE

[MIT](LICENSE)