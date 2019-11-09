package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/therealplato/vigenere"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Pass cipher key on command line")
	}
	fmt.Println("Input ciphertext followed by EOF. Whitespace will be discarded.")
	key := os.Args[1]
	key = strings.ToUpper(key)
	bb, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatalf("error reading plaintext: %v\n", err)
	}
	plain := vigenere.Decode(bb, []byte(key))
	fmt.Println("")
	fmt.Println(string(plain))
}
