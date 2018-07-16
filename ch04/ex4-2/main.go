// ex4.2 prints the SHA hash of stdin.
package main

import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

var (
	hashAlgo = flag.Int("sha", 256, "set hash algorithm as sha384")
)

func main() {
	flag.Parse()

	var encoder func([]byte) []byte
	switch *hashAlgo {
	case 384:
		encoder = hashBySha384
	case 512:
		encoder = hashBySha512
	default:
		encoder = hashBySha256
	}

	sc := bufio.NewScanner(os.Stdin)
	for sc.Scan() {
		fmt.Printf("sha%d: %x\n", *hashAlgo, encoder([]byte(sc.Text())))
	}
}

func hashBySha256(b []byte) []byte {
	h := sha256.Sum256(b)
	return h[:]
}

func hashBySha384(b []byte) []byte {
	h := sha512.Sum384(b)
	return h[:]
}

func hashBySha512(b []byte) []byte {
	h := sha512.Sum512(b)
	return h[:]
}
