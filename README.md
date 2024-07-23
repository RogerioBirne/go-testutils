# go-testutils
The helper container with some testing utilities for golang


## Installation

```bash
go get https://github.com/RogerioBirne/go-testutils
```

## Features
### Argument Captor
ArgumentCaptor is used to capture any argument for mocked methods like Mockito captor.

See. https://site.mockito.org/javadoc/current/org/mockito/ArgumentCaptor.html

Example of using: [argumentcaptor_test.go](testutils%2Fargumentcaptor_test.go)


### Argument Injector
Create a `mock.MatchedBy` able to set a value when the target func call the stubby.

It can be used return a value when the target call a function with a pointer parameter.

Example of using: [argumentinjector_test.go](testutils%2Fargumentinjector_test.go)