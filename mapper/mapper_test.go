package mapper_test

import (
	"fmt"
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

func TestSimple(t *testing.T) {
	profile, err := mapper.CreateProfile(testUser{}, testUserDto{})
	if err != nil {
		t.Fatal(err)
	}
	err = profile.ForMember("FullName", func(src testUser) any {
		return src.FirstName + " " + src.SurName
	})
	if err != nil {
		t.Fatal(err)
	}

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

type testUserNested struct {
	Name    string
	Contact testContact
}

type testContact struct {
	Number string
	Email  string
}

type testUserNestedDto struct {
	Name    string
	Contact testContactDto
}

type testContactDto struct {
	Number string
	Email  string
}

func TestNested(t *testing.T) {
	p, err := mapper.CreateProfile(testUserNested{}, testUserNestedDto{})
	if err != nil {
		t.Fatal(err)
	}
	err = p.ForMember("Contact", func(src testUserNested) any {
		contactDto := testContactDto{}

		err := mapper.Map(src.Contact, &contactDto)
		if err != nil {
			t.Fatal(err)
		}

		return contactDto
	})
	if err != nil {
		t.Fatal(err)
	}

	_, err = mapper.CreateProfile(testContact{}, testContactDto{})
	if err != nil {
		t.Fatal(err)
	}

	u := testUserNested{
		Name: "John Doe",
		Contact: testContact{
			Number: "303-4040",
			Email:  "johndoe@email.com",
		},
	}

	dto := testUserNestedDto{}

	err = mapper.Map(u, &dto)
	if err != nil {
		t.Fatal(err)
	}

	if dto.Name != u.Name {
		t.Fatal("Name not mapped correctly " + fmt.Sprint(dto))
	}

	if dto.Contact.Email != u.Contact.Email {
		t.Fatal("Email not mapped correctly " + fmt.Sprint(dto))
	}

	if dto.Contact.Number != u.Contact.Number {
		t.Fatal("Number not mapped correctly " + fmt.Sprint(dto))
	}
}
