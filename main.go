package main

import (
	"encoding/hex"
	"fmt"
	"hash/crc32"
	"io"
	"os"
)

func hashCRC32(in io.Reader, polynomial uint32) (string, error) {
	table := crc32.MakeTable(polynomial)
	hash := crc32.New(table)
	if _, err := io.Copy(hash, in); err != nil {
		return "", err
	}

	hashInBytes := hash.Sum(nil)[:]
	hashStr := hex.EncodeToString(hashInBytes)

	return hashStr, nil
}

func main() {
	filePath := os.Args[1]
	f, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	hash, err := hashCRC32(f, 0xedb88320)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s: %s\n", filePath, hash)
}
