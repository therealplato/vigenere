package vigenere

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func TestDecode(t *testing.T) {
	in, err := ioutil.ReadFile("testdata/homonym.cipher")
	if err != nil {
		t.Fatalf("failure reading golden file: %v", err)
	}

	expected, err := ioutil.ReadFile("testdata/homonym.plain")
	if err != nil {
		t.Fatalf("failure reading golden file: %v", err)
	}
	expected = bytes.TrimSpace(expected)

	actual := Decode(in, []byte("HOMONYM"))
	if string(expected) != string(actual) {
		t.Fatalf("mismatching plaintext. Got: %q", actual)
	}
}

func TestEncode(t *testing.T) {
	in, err := ioutil.ReadFile("testdata/homonym.plain")
	if err != nil {
		t.Fatalf("failure reading golden file: %v", err)
	}

	expected, err := ioutil.ReadFile("testdata/homonym.cipher")
	if err != nil {
		t.Fatalf("failure reading golden file: %v", err)
	}
	expected = bytes.TrimSpace(expected)

	actual := Encode(in, []byte("HOMONYM"))
	if string(expected) != string(actual) {
		t.Fatalf("mismatching ciphertext. Got: %q", actual)
	}
}

func TestFives(t *testing.T) {
	in, err := ioutil.ReadFile("testdata/homonym.plain")
	if err != nil {
		t.Fatalf("failure reading golden file: %v", err)
	}
	in = bytes.TrimSpace(in)

	expected, err := ioutil.ReadFile("testdata/homonym.plain.fives")
	if err != nil {
		t.Fatalf("failure reading golden file: %v", err)
	}
	expected = bytes.TrimSpace(expected)

	actual := Fives(in)
	if string(expected) != string(actual) {
		t.Fatalf("five character groupings not as expected. Got: %q", actual)
	}
}
