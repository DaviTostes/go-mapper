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

type ContactDTO struct {
	Number string
	Email  string
}

type UserDTO struct {
	Name    string
	Contact ContactDTO
}

func main() {
	// Create mapping profile for User -> UserDTO
	profile, err := mapper.CreateProfile(User{}, UserDTO{})
	if err != nil {
		panic(err)
	}

	// Configure nested struct mapping
	err = profile.ForMember("Contact", func(src User) any {
		contactDTO := ContactDTO{}
		err := mapper.Map(src.Contact, &contactDTO)
		if err != nil {
			panic(err)
		}
		return contactDTO
	})
	if err != nil {
		panic(err)
	}

	// Create mapping profile for Contact -> ContactDTO
	_, err = mapper.CreateProfile(Contact{}, ContactDTO{})
	if err != nil {
		panic(err)
	}

	// Create a sample user with nested contact
	user := User{
		Name: "John Doe",
		Contact: Contact{
			Number: "303-4040",
			Email:  "johndoe@email.com",
		},
	}

	// Create destination DTO and perform mapping
	dto := UserDTO{}
	err = mapper.Map(user, &dto)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Mapped DTO: %+v\n", dto)
	// Output: Mapped DTO: {Name:John Doe Contact:{Number:303-4040 Email:johndoe@email.com}}
}

