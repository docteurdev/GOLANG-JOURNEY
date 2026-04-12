package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	copyedFile, err := os.Open("file.txt")
	if err != nil {
		log.Fatal("====1", err)
	}

	defer copyedFile.Close()

	// readFile, err := io.Copy(os.Stdout, copyedFile)

	if err != nil {
		log.Fatal("===2", err)
	}
	_, _ = copyedFile.Seek(5, io.SeekStart)

	fmt.Printf("%s", copyedFile)
}
