package examples

import (
	"fmt"

	"github.com/davitostes/go-mapper/mapper"
)

type User03 struct {
	FirstName string
	SurName   string
	Age       uint
}

type User03DTO struct {
	FirstName string
	SurName   string
	Age       uint
	FullName  string
}

func ListMapping() {
	// Create mapping profile for User03 -> User03DTO
	profile, err := mapper.CreateProfile(User03{}, User03DTO{})
	if err != nil {
		panic(err)
	}

	// Configure custom mapping for FullName field
	err = profile.ForMember("FullName", func(src User03) any {
		return src.FirstName + " " + src.SurName
	})
	if err != nil {
		panic(err)
	}

	// Create a slice of users
	users := []User03{
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
	var dtos []User03DTO

	// Map the entire slice
	err = mapper.MapList(users, &dtos)
	if err != nil {
		panic(err)
	}

	// Print each mapped DTO
	for i, dto := range dtos {
		fmt.Printf("User03 %d: %+v\n", i+1, dto)
	}
	// Output:
	// User03 1: {FirstName:John SurName:Doe Age:45 FullName:John Doe}
	// User03 2: {FirstName:Jane SurName:Smith Age:32 FullName:Jane Smith}
	// User03 3: {FirstName:Bob SurName:Johnson Age:28 FullName:Bob Johnson}
}
