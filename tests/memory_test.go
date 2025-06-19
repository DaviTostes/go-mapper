package mapper_test

import (
	"testing"

	"github.com/davitostes/go-mapper/mapper"
)

type testUser struct {
	Name          string
	Age           uint
	password_hash string
}

type testUserDto struct {
	Name     string
	Age      uint
	password string
}

func TestSaveAndLoadStruct(t *testing.T) {
	mapper.CreateProfile(testUserDto{}, testUser{}, func() any { return "sas" })

	_, ok := mapper.LoadStruct[mapper.Profile[testUserDto, testUser]]()
	if !ok {
		t.Fatal("Expected struct, got none")
	}
}

func TestLoadWithoutSave(t *testing.T) {
	_, ok := mapper.LoadStruct[mapper.Profile[testUser, testUserDto]]()
	if ok {
		t.Error("Expected false when loading unsaved struct")
	}
}

func TestTypeMismatch(t *testing.T) {
	mapper.CreateProfile(testUserDto{}, testUser{}, func() any { return "sas" })

	_, ok := mapper.LoadStruct[mapper.Profile[testUser, testUserDto]]()
	if ok {
		t.Fatal("Expected false on type mismatch")
	}
}
