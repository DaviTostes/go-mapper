package mapper_test

import (
	"testing"

	"github.com/davitostes/go-mapper/mapper"
)

func TestProfileSourceIsNotStruct(t *testing.T) {
	if _, err := mapper.CreateProfile(1, testUserDto{}); err == nil {
		t.Fatal("Expected error on Source is not struct")
	}
}

func TestProfileDestinyIsNotStruct(t *testing.T) {
	if _, err := mapper.CreateProfile(testUser{}, map[string]string{}); err == nil {
		t.Fatal("Expected error on Destiny is not struct")
	}
}
