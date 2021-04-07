package main

import (
	"encoding/gob"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

// type KV struct {
// 	Key   string
// 	Value string
// }

func main() {

	/*
		Let's:
		> create an output file
		> create an empty hash table
		> fill the hash table with random strings as key and values
		> encode the hash table using the gob encoder
		> print the encoded content into the output file
	*/

	// Creates the output file
	outFile, err := os.Create("tinyDB")
	if err != nil {
		log.Fatal("Cannot open file: ", err)
	}

	// Do not close the file right away
	defer outFile.Close()

	// Creates a seed for the pseudo random generator
	rand.NewSource(time.Now().UnixNano())

	// Creates an empty hash table
	var kv = map[string]string{}

	// Number of key-value pairs to add in the hash table
	const maxIter = 10
	for i := 0; i < maxIter; i++ {

		key := RandStringBytesRmndr(25)
		value := RandStringBytesRmndr(25)
		kv[key] = value
	}
	// Prints the content of the hash table
	fmt.Println("Encoding...")
	for key := range kv {
		fmt.Printf("%s:%s\n", key, kv[key])
	}
	// Creates a new gob encoder
	encoder := gob.NewEncoder(outFile)

	// Encodes the hash table and outputs to a file
	encoder.Encode(kv)

	/*
		Let's:
		> read the output file
		> decode the content using the gob decoder
		> print the content of the decoded content
	*/

	// Throws kv to GC
	kv = nil

	// Creates an empty hash table
	kv = map[string]string{}

	// Open the output file
	inFile, err := os.Open("tinyDB")
	if err != nil {
		log.Fatal("Cannot open file: ", err)
	}

	// Do not close the file right away
	defer inFile.Close()

	// Creates a new gob decoder
	decoder := gob.NewDecoder(inFile)

	// Decodes the file content and saves it in the hash table kv
	decoder.Decode(&kv)

	// Prints the content of the hash table
	fmt.Println("\nDecoding...")
	for key := range kv {
		fmt.Printf("%s:%s\n", key, kv[key])
	}

}

// RandStringBytesRmndr randomly generates strings
// https://stackoverflow.com/a/31832326
func RandStringBytesRmndr(n int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return string(b)
}
