package main

import (
	"log"
	"os"
	"testing"
)

func TestRemove_SimpleCase(t *testing.T) {
	expect_len := 2
	expect_result := []string{"a", "c"}
	input := []string{"a", "b", "c"}

	result := remove(input, 1)

	if expect_len != len(result) {
		t.Errorf("Result not match, got: %d, want: %d.", len(result), expect_len)
	}

	if result[0] != expect_result[0] || result[1] != expect_result[1] {
		t.Errorf("Result not match, got: %s, want: %s.", result, expect_result)
	}
}

func TestRemove_LenOutOfBound(t *testing.T) {
	expect_len := 3
	expect_result := []string{"a", "b", "c"}
	input := []string{"a", "b", "c"}

	result := remove(input, 4)

	if expect_len != len(result) {
		t.Errorf("Result not match, got: %d, want: %d.", len(result), expect_len)
	}

	if result[0] != expect_result[0] || result[2] != expect_result[2] {
		t.Errorf("Result not match, got: %s, want: %s.", result, expect_result)
	}
}

func TestRemove_LenOutOfBound_MinusValue(t *testing.T) {
	expect_len := 3
	expect_result := []string{"a", "b", "c"}
	input := []string{"a", "b", "c"}

	result := remove(input, -1)

	if expect_len != len(result) {
		t.Errorf("Result not match, got: %d, want: %d.", len(result), expect_len)
	}

	if result[0] != expect_result[0] || result[2] != expect_result[2] {
		t.Errorf("Result not match, got: %s, want: %s.", result, expect_result)
	}
}

func TestRemove_RemoveFirstIndex(t *testing.T) {
	expect_len := 2
	expect_result := []string{"b", "c"}
	input := []string{"a", "b", "c"}

	result := remove(input, 0)

	if expect_len != len(result) {
		t.Errorf("Result not match, got: %d, want: %d.", len(result), expect_len)
	}

	if result[0] != expect_result[0] || result[1] != expect_result[1] {
		t.Errorf("Result not match, got: %s, want: %s.", result, expect_result)
	}
}

func TestReadLines_NullSafetyTesting(t *testing.T) {
	_, err := readLines("")
	if nil == err {
		t.Error("Error not being raised!")
	}
}

func TestWriteLines_NullSafetyTesting(t *testing.T) {
	err := writeLines(nil, "")
	if err.Error() != REQUIRED_PARAM_ERROR {
		t.Error("Error does not matched")
	}

	lines := []string{"a", "b"}
	err = writeLines(lines, "")
	if nil == err {
		t.Error("Error not being raised!")
	}
}

func TestWriteLines_SimpleCase(t *testing.T) {
	lines := []string{"a", "b"}
	err := writeLines(lines, "./path.txt")
	if nil != err {
		t.Errorf("Error is raised, got: %s, want: no error", err.Error())
	}

	// os.IsExist("./path.txt")

	e := os.Remove("./path.txt")
	if e != nil {
		log.Fatal(e)
	}
}
