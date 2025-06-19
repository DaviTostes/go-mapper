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
The following demonstrates how `go-mapper` can map fields between structs, including custom transformations:
```go
package main

import (
	"fmt"
	"log"

	"github.com/davitostes/go-mapper/mapper"
	"golang.org/x/crypto/bcrypt"
)

type createUserDto struct {
	Name     string
	Age      uint
	Password string
}

type user struct {
	Name         string
	Age          uint
	PasswordHash string
}

func main() {
	profile, err := mapper.CreateProfile(createUserDto{}, user{})
	if err != nil {
		log.Fatal(err)
	}

	profile.ForMember("PasswordHash", func(dto createUserDto) any {
		hash, err := bcrypt.GenerateFromPassword([]byte(dto.Password), bcrypt.DefaultCost)
		if err != nil {
			log.Fatal(err)
		}
		return string(hash)
	})

	dto := createUserDto{Name: "davi", Age: 21, Password: "123"}
	u := user{}

	err = mapper.Map(dto, &u)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(u)
}
```

## Examples
Additional examples can be found in the `tests` directory, showcasing various use cases and advanced features.

## Contributing
Contributions to `go-mapper` are highly encouraged! Whether it's reporting bugs, suggesting features, or submitting pull requests, your input is invaluable. Please visit [GitHub](https://github.com/davitostes/go-mapper) to contribute or discuss.

## Contact
For questions, feedback, or support, reach out to the project maintainer at [Author Name](mailto:davisiqueira591@gmail.com).
