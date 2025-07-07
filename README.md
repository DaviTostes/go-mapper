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

## Advanced Examples

### Nested Struct Mapping
The mapper can handle nested structs. Here's an example of mapping nested structures:

```go
package main

import (
	"fmt"
	"github.com/davitostes/go-mapper/mapper"
)

type Contact struct {
	Number string
	Email  string
}

type User struct {
	Name    string
	Contact Contact
}

type ContactDto struct {
	Number string
	Email  string
}

type UserDto struct {
	Name    string
	Contact ContactDto
}

func main() {
	// Create mapping profile for User -> UserDto
	profile, err := mapper.CreateProfile(User{}, UserDto{})
	if err != nil {
		panic(err)
	}

	// Configure nested struct mapping
	profile.ForMember("Contact", func(src User) any {
		contactDto := ContactDto{}
		err := mapper.Map(src.Contact, &contactDto)
		if err != nil {
			panic(err)
		}
		return contactDto
	})

	// Create mapping profile for Contact -> ContactDto
	_, err = mapper.CreateProfile(Contact{}, ContactDto{})
	if err != nil {
		panic(err)
	}

	// Create a user with nested contact
	user := User{
		Name: "John Doe",
		Contact: Contact{
			Number: "303-4040",
			Email:  "johndoe@email.com",
		},
	}

	// Map to DTO
	dto := UserDto{}
	err = mapper.Map(user, &dto)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", dto)
	// Output: {Name:John Doe Contact:{Number:303-4040 Email:johndoe@email.com}}
}
```

## Contact
For questions, feedback, or support, reach out to the project maintainer at [Davi Tostes](mailto:davisiqueira591@gmail.com).
