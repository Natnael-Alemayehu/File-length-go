package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
)

func getFile(name string) (*os.File, func(), error) {
	file, err := os.Open(name)
	if err != nil {
		return nil, nil, err
	}
	return file, func() {
		file.Close()
	}, nil
}

func fileLen(name string) (int, error) {
	f, closer, err := getFile(name)
	if err != nil {
		return 0, errors.New("file not specified")
	}
	defer closer()

	data := make([]byte, 1_000_000)

	count, err := f.Read(data)
	if err != nil {
		if err != io.EOF {
			return 0, err
		}
	}

	return count, nil
}

func main() {
	if len(os.Args) < 2 {
		log.Fatal("File not specified")
	}

	len, err := fileLen(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(len)

}
