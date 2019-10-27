package main

import (
	"encoding/hex"
	"fmt"
	"hash/crc32"
	"io"
	"os"
)

func hashFileCRC32(filePath string, polynomial uint32) (string, error) {
	fin, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer fin.Close()

	tablePolynomial := crc32.MakeTable(polynomial)
	hash := crc32.New(tablePolynomial)
	if _, err = io.Copy(hash, fin); err != nil {
		return "", err
	}

	hashInBytes := hash.Sum(nil)[:]
	hashString := hex.EncodeToString(hashInBytes)

	return hashString, nil
}

func main() {
	filePath := os.Args[1]
	hash, err := hashFileCRC32(filePath, 0xedb88320)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s: %s\n", filePath, hash)
}
