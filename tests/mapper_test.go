package mapper_test

import (
	"testing"

	"github.com/davitostes/go-mapper/mapper"
)

type testUserDto struct {
	FirstName string
	SurName   string
	Age       uint
	FullName  string
}

type testUser struct {
	FirstName string
	SurName   string
	Age       uint
}

func Test(t *testing.T) {
	profile, err := mapper.CreateProfile(testUser{}, testUserDto{})
	if err != nil {
		t.Fatal(err)
	}
	profile.ForMember("FullName", func(src testUser) any {
		return src.FirstName + " " + src.SurName
	})

	u := testUser{
		FirstName: "John",
		SurName:   "Doe",
		Age:       45,
	}

	dto := testUserDto{}

	err = mapper.Map(u, &dto)
	if err != nil {
		t.Fatal(err)
	}

	if dto.Age != u.Age {
		t.Fatal("Age not mapped correctly")
	}

	if dto.FirstName != u.FirstName {
		t.Fatal("FirstName not mapped correctly")
	}

	if dto.SurName != u.SurName {
		t.Fatal("SurName not mapped correctly")
	}

	if dto.FullName != u.FirstName+" "+u.SurName {
		t.Fatal("FullName not mapped correctly")
	}
}
