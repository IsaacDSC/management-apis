package util

import (
	"encoding/json"
	"os"
	"testing"
)

const filePath = "valid.json"

func init() {
	os.WriteFile(filePath, []byte(`{"key": "value"}`), 0644)
}

func TestReadFileSuccess(t *testing.T) {

	// Defer is used to ensure that the file is closed after the test ends
	defer os.Remove(filePath)

	expected := map[string]interface{}{"key": "value"}
	var result map[string]interface{}

	err := ReadFile(filePath, &result)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if !jsonEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestReadFileFileNotFound(t *testing.T) {
	filePath := "testdata/nonexistent.json"
	var result map[string]interface{}

	err := ReadFile(filePath, &result)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}

func TestReadFileInvalidJSON(t *testing.T) {
	filePath := "testdata/invalid.json"
	var result map[string]interface{}

	err := ReadFile(filePath, &result)
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
}

func jsonEqual(a, b map[string]interface{}) bool {
	aj, _ := json.Marshal(a)
	bj, _ := json.Marshal(b)
	return string(aj) == string(bj)
}
