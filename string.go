package vigenere

import (
	"bytes"
	"strings"
)

var allowed = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

func filter(r rune) rune {
	if bytes.ContainsRune(allowed, r) {
		return r
	}
	return -1
}

// Print the byte slice as utf-8 strings in ten groups of five per line
// We've stripped non-ascii so can safely index the slice
func Fives(bb []byte) string {
	var msg string
	for i := 0; i < len(bb); i += 5 {
		var group string
		if i+5 > len(bb) {
			group = string(bb[i:])
		} else {
			group = string(bb[i : i+5])
		}
		msg += group
		if i > 0 && i%60 == 0 {
			msg += "\n"
		} else {
			msg += " "
		}
	}
	return strings.TrimSpace(msg)
}

// Remove non-alphanumeric characters. Capitalize all characters. Returns copy.
func strip(bb []byte) []byte {
	out := bytes.ToUpper(bb)
	out = bytes.Map(filter, out)
	return out
}
