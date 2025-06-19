package mapper_test

import (
	"testing"

	"github.com/davitostes/go-mapper/mapper"
)

func TestProfileNotFound(t *testing.T) {
	user := testUserDto{Name: "Davi", Age: 21, password: "davi123"}
	dto := testUserDto{}

	if err := mapper.Map(user, &dto); err == nil {
		t.Fatal("Expected error on Profile not found")
	}
}

func TestSourceIsNotStruct(t *testing.T) {
	var user = 1
	var dto = testUserDto{}

	mapper.CreateProfile(user, dto)

	if err := mapper.Map(user, &dto); err == nil {
		t.Fatal("Expected error on Source is not struct")
	}
}

func TestDestinyIsNotStruct(t *testing.T) {
	var user = testUser{}
	var dto = "test"

	mapper.CreateProfile(user, dto)

	if err := mapper.Map(user, &dto); err == nil {
		t.Fatal(err)
		t.Fatal("Expected error on Destiny is not struct")
	}
}

func Test(t *testing.T) {
	dto := testUserDto{Name: "Davi", Age: 21, password: "davi123"}
	user := testUser{}

	mapper.CreateProfile(testUserDto{}, testUser{})

	if err := mapper.Map(dto, &user); err != nil {
		t.Fatal(err)
	}
}
