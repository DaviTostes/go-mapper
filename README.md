# go-mapper

## Introduction
`go-mapper` is a Go package for mapping values between structs. Inspired by the powerful AutoMapper library from .NET, `go-mapper` brings similar functionality to Go, simplifying object-to-object mapping. By leveraging Go's reflection capabilities, it minimizes repetitive boilerplate code and provides developers with an intuitive API for custom mapping.

## Installation
To use `go-mapper`, install it using Go's module system:
```bash
go get -u github.com/davitostes/go-mapper
```
Ensure you have Go installed and your project initialized as a Go module.

## Usage
### Basic Example
Below is an example of how `go-mapper` can map fields between structs:
```go
package main

import (
	"fmt"

	"github.com/davitostes/go-mapper/mapper"
)

type user struct {
	FirstName string
	SurName   string
	Age       uint
}

type readUserDto struct {
	FirstName string
	SurName   string
	Age       uint
	FullName  string
}

func main() {
    // Creating mapping profile [user] -> [readUserDto]
	profile, err := mapper.CreateProfile(user{}, readUserDto{})
	if err != nil {
		panic(err)
	}
    // Setting special mapper for FullName
	profile.ForMember("FullName", func(src user) any {
		return src.FirstName + " " + src.SurName
	})

	u := user{
		FirstName: "John",
		SurName:   "Doe",
		Age:       45,
	}

	dto := readUserDto{}

    // Mapping to dto
	err = mapper.Map(u, &dto)
	if err != nil {
		panic(err)
	}

	fmt.Println(dto) // Output: {John Doe 45 John Doe}
}
```

## Additional Examples
Explore the `tests` directory for more examples demonstrating advanced use cases and error handling.

## Contact
For questions, feedback, or support, reach out to the project maintainer at [Davi Tostes](mailto:davisiqueira591@gmail.com).
