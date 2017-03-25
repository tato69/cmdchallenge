package main

import (
	"testing"
)

func TestReadChallenges(t *testing.T) {
	var cBlob = []byte(`[
	{"slug": "c1", "version": 4, "author": "cmdchallenge", "description": "foo", "example": "bar",
	 "expected_output": {"lines": ["herp", "derp"]}},
	{"slug": "c2", "version": 4, "author": "cmdchallenge", "description": "foo", "example": "bar",
	 "expected_output": {"lines": ["hello world"]}},
	{"slug": "c3", "author": "cmdchallenge", "description": "foo", "example": "bar", "expected_output":
	 {"order": false, "lines": ["hello world"]}},
	{"slug": "c4", "version": 4, "author": "cmdchallenge", "description": "foo", "example": "bar"}
	]`)
	ch, err := readChallenges(cBlob)
	if err != nil {
		t.Error("Expected no error, got ", err)
	}
	challenge, err := ch.Slug("c4")
	if err != nil {
		t.Error("Expected to find slug 'c4' but got", err)
	}
	if challenge.ExpectedOutput.Order == false {
		t.Error("For order expected false but got", challenge.ExpectedOutput.Order)
	}
	if ch.Len() != 4 {
		t.Error("Expected 4 challenges, got", ch.Len())
	}
}
