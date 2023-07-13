package main

import "testing"

func TestIsValid(t *testing.T) {
	if !isValid(111111, false) {
		t.Error("Par one fail: 111111 should be true")
	}

	if isValid(223450, false) {
		t.Error("Par one fail: 223450 should be false")
	}

	if isValid(123789, false) {
		t.Error("Par one fail: 123789 should be false")
	}

	if !isValid(112233, true) {
		t.Error("Par two fail: 111122 should be true")
	}

	if isValid(123444, true) {
		t.Error("Par two fail: 123444 should be false")
	}

	if !isValid(111122, true) {
		t.Error("Par two fail: 112233 should be true")
	}
}
