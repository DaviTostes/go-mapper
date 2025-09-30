package examples

import (
	"fmt"

	"github.com/davitostes/go-mapper/mapper"
)

type Contact02 struct {
	Number string
	Email  string
}

type User02 struct {
	Name    string
	Contact02 Contact02
}

type Contact02DTO struct {
	Number string
	Email  string
}

type User02DTO struct {
	Name    string
	Contact02 Contact02DTO
}

func NestedMapping() {
	// Create mapping profile for User02 -> User02DTO
	profile, err := mapper.CreateProfile(User02{}, User02DTO{})
	if err != nil {
		panic(err)
	}

	// Configure nested struct mapping
	err = profile.ForMember("Contact02", func(src User02) any {
		contactDTO := Contact02DTO{}
		err := mapper.Map(src.Contact02, &contactDTO)
		if err != nil {
			panic(err)
		}
		return contactDTO
	})
	if err != nil {
		panic(err)
	}

	// Create mapping profile for Contact02 -> Contact02DTO
	_, err = mapper.CreateProfile(Contact02{}, Contact02DTO{})
	if err != nil {
		panic(err)
	}

	// Create a sample user with nested contact
	user := User02{
		Name: "John Doe",
		Contact02: Contact02{
			Number: "303-4040",
			Email:  "johndoe@email.com",
		},
	}

	// Create destination DTO and perform mapping
	dto := User02DTO{}
	err = mapper.Map(user, &dto)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Mapped DTO: %+v\n", dto)
	// Output: Mapped DTO: {Name:John Doe Contact02:{Number:303-4040 Email:johndoe@email.com}}
}

