package main

import (
	"fmt"

	"github.com/davitostes/go-mapper/mapper"
)

type User struct {
	FirstName string
	SurName   string
	Age       uint
}

type UserDTO struct {
	FirstName string
	SurName   string
	Age       uint
	FullName  string
}

func main() {
	// Create mapping profile for User -> UserDTO
	profile, err := mapper.CreateProfile(User{}, UserDTO{})
	if err != nil {
		panic(err)
	}

	// Configure custom mapping for FullName field
	err = profile.ForMember("FullName", func(src User) any {
		return src.FirstName + " " + src.SurName
	})
	if err != nil {
		panic(err)
	}

	// Create a slice of users
	users := []User{
		{
			FirstName: "John",
			SurName:   "Doe",
			Age:       45,
		},
		{
			FirstName: "Jane",
			SurName:   "Smith",
			Age:       32,
		},
		{
			FirstName: "Bob",
			SurName:   "Johnson",
			Age:       28,
		},
	}

	// Create destination slice for DTOs
	var dtos []UserDTO

	// Map the entire slice
	err = mapper.MapList(users, &dtos)
	if err != nil {
		panic(err)
	}

	// Print each mapped DTO
	for i, dto := range dtos {
		fmt.Printf("User %d: %+v\n", i+1, dto)
	}
	// Output:
	// User 1: {FirstName:John SurName:Doe Age:45 FullName:John Doe}
	// User 2: {FirstName:Jane SurName:Smith Age:32 FullName:Jane Smith}
	// User 3: {FirstName:Bob SurName:Johnson Age:28 FullName:Bob Johnson}
}

