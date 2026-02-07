package test

import (
	"reflect"
	"testing"
)

func Assert[T comparable](t testing.TB, expected, actual T) {
	t.Helper()

	if expected != actual {
		t.Errorf("Assertion failed: expected %v, got %v", expected, actual)
	}
}

func AssertNot[T comparable](t testing.TB, expected, actual T) {
	t.Helper()

	if expected == actual {
		t.Errorf("Assertion failed: expected %v not to be %v", actual, expected)
	}
}

func AssertEqual[T any](t testing.TB, expected, actual T) {
	t.Helper()

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Assertion failed: expected '%v', got '%v'", expected, actual)
	}
}
