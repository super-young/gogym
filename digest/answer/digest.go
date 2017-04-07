package main

import (
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"hash"
	"io"
	"os"
)

func main() {
	digestalgo := os.Args[1]
	var hash hash.Hash

	switch digestalgo {
	case "sha256":
		hash = sha256.New()
	case "sha512":
		hash = sha512.New()
	case "md5":
		hash = md5.New()
	}

	io.Copy(hash, os.Stdin)

	var digest []byte
	digest = hash.Sum(digest)

	fmt.Println(hex.EncodeToString(digest))
}
