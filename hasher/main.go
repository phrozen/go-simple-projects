package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"hash"
	"io"
	"os"
	"sort"
)

func main() {
	// Define the hash algorithms by name and implemantation
	hashes := map[string]hash.Hash{
		"md5":    md5.New(),
		"sha1":   sha1.New(),
		"sha256": sha256.New(),
		"sha512": sha512.New(),
	}
	// Create a flag map to store decisions
	flags := make(map[string]*bool)
	for k := range hashes {
		// Create a boolean flag for each hash algorithm
		flags[k] = flag.Bool(k, false, fmt.Sprintf("Calculates %s hash of file", k))
	}
	// Parse flags from command line
	flag.Parse()
	// writers slice needed for io.MultiWriter
	var writers []io.Writer
	// keys slice for sorting purposes
	var keys []string
	for k, v := range flags {
		if *v {
			keys = append(keys, k)
			writers = append(writers, hashes[k])
		}
	}
	// If no valid flags, print usage and Exit
	if len(writers) == 0 {
		flag.Usage()
		os.Exit(1)
	}
	// Open the file given as argument and panic if it fails
	in, err := os.Open(flag.Arg(0))
	if err != nil {
		panic(err)
	}
	// Copy the file Reader into a MultiWriter for each hash
	io.Copy(io.MultiWriter(writers...), in)
	// Sort keys for consistency and iterate through them
	sort.Strings(keys)
	for _, k := range keys {
		// Print every calculated sum in order
		fmt.Printf("%6s: %x\n", k, hashes[k].Sum(nil))
	}
}
