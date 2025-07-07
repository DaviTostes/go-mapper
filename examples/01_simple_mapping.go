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
	FullName  string // This will be computed from FirstName and SurName
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

	// Create a sample user
	user := User{
		FirstName: "John",
		SurName:   "Doe",
		Age:       45,
	}

	// Create destination DTO and perform mapping
	dto := UserDTO{}
	err = mapper.Map(user, &dto)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Mapped DTO: %+v\n", dto)
	// Output: Mapped DTO: {FirstName:John SurName:Doe Age:45 FullName:John Doe}
}

