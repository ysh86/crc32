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
	for i := 1; i < len(os.Args); i++ {
		filePath := os.Args[i]
		f, err := os.Open(filePath)
		if err != nil {
			panic(err)
		}

		hash, err := hashCRC32(f, 0xedb88320)
		if err != nil {
			f.Close()
			panic(err)
		}

		fmt.Printf("%s  %s\n", hash, filePath)
		f.Close()
	}
}
