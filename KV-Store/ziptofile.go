package main

import (
	"bytes"
	"compress/gzip"
	"encoding/gob"
	"io/ioutil"
	"log"
)

func ZipToFile(kvs map[string]string, filename string) {
	//1. kvs (the key-value store passed as argument to this function)
	//is a composite datastructure. Before we can write it to file, we have to
	//encode it. For this we use gob.
	var kvs_buffer bytes.Buffer
	encoder := gob.NewEncoder(&kvs_buffer)

	err := encoder.Encode(kvs)
	if err != nil {
		log.Fatal(err)
	}

	//2. Create byte buffer that will hold compressed bytes
	//and encode (write) the bytes in the kvs_buffer.
	var zip_buffer bytes.Buffer
	gzipwriter := gzip.NewWriter(&zip_buffer)

	gzipwriter.Write(kvs_buffer.Bytes())
	//Close gzipwriter, causing bytes to be flushed to the buffer.
	gzipwriter.Close()

	//3. Write buffer containing compressed kvs data structure to file.
	err = ioutil.WriteFile(filename, zip_buffer.Bytes(), 0755)
	if err != nil {
		log.Fatal(err)
	}
}
