package main

import (
	"encoding/hex"
	"fmt"
	"hash/crc32"
	"io"
	"os"
)

func hash_file_crc32(filePath string, polynomial uint32) (string, error) {
	var returnCRC32String string
	file, err := os.Open(filePath)
	if err != nil {
		return returnCRC32String, err
	}
	defer file.Close()
	tablePolynomial := crc32.MakeTable(polynomial)
	hash := crc32.New(tablePolynomial)
	if _, err := io.Copy(hash, file); err != nil {
		return returnCRC32String, err
	}
	hashInBytes := hash.Sum(nil)[:]
	returnCRC32String = hex.EncodeToString(hashInBytes)
	return returnCRC32String, nil

}

func main() {
	filePath := os.Args[1]
	hash, err := hash_file_crc32(filePath, 0xedb88320)
	if err == nil {
		fmt.Printf("%s %s\n", filePath, hash)
	}
}
