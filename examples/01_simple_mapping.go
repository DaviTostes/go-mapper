package examples

import (
	"fmt"

	"github.com/davitostes/go-mapper/mapper"
)

type User01 struct {
	FirstName string
	SurName   string
	Age       uint
}

type User01DTO struct {
	FirstName string
	SurName   string
	Age       uint
	FullName  string // This will be computed from FirstName and SurName
}

func SimpleMapping() {
	// Create mapping profile for User01 -> User01DTO
	profile, err := mapper.CreateProfile(User01{}, User01DTO{})
	if err != nil {
		panic(err)
	}

	// Configure custom mapping for FullName field
	err = profile.ForMember("FullName", func(src User01) any {
		return src.FirstName + " " + src.SurName
	})
	if err != nil {
		panic(err)
	}

	// Create a sample user
	user := User01{
		FirstName: "John",
		SurName:   "Doe",
		Age:       45,
	}

	// Create destination DTO and perform mapping
	dto := User01DTO{}
	err = mapper.Map(user, &dto)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Mapped DTO: %+v\n", dto)
	// Output: Mapped DTO: {FirstName:John SurName:Doe Age:45 FullName:John Doe}
}

