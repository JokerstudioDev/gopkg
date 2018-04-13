package dac

import (
	"testing"
)

type MockAccount struct{}

func (MockAccount) GetAccount() string {
	return "joker"
}

func TestGetAccountInfo(t *testing.T) {
	expected := "joker studio"
	result := GetAccountInformation(MockAccount{})
	if result != expected {
		t.Errorf("unexpected")
	}
}
