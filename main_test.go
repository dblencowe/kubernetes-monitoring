package main

import (
	"os"
	"testing"
)

func TestGetEnvInvalidVariable(t *testing.T) {
	_, err := getenv("test", "")
	if err == nil {
		t.Fatalf("Expected error, got non")
	}
}

func TestGetEnvInvalidVariableDefaultValue(t *testing.T) {
	value, err := getenv("test", "TheDefault")
	if err != nil || value != "TheDefault" {
		t.Fatalf("Expected value of TheDefault, got %s or error %f", value, err)
	}
}

func TestGetEnv(t *testing.T) {
	os.Setenv("TEST_VAR", "TheExpectedValue")
	value, err := getenv("TEST_VAR", "SomethingElse")
	if err != nil || value != "TheExpectedValue" {
		t.Fatalf("Expected value of TheExpectedValue, got %s or error %f", value, err)
	}
}
