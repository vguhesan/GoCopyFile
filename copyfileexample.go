package main

import (
	"bufio"
	"io"
	"os"
)

func main() {
	// Objective: Copy contents of FileA to FileB
	// Perform this step in chuncks instead of using the
	// built-in copy functionality.
	// Why?
	// Illustrates how to use bufio library with a reader and writer

	// Source
	inFile, err := os.Open("./source-file.txt")
	if err != nil {
		panic(err)
	}
	defer inFile.Close()

	// Target
	outFile, err := os.Create("./target-file.txt")
	if err != nil {
		panic(err)
	}
	defer outFile.Close()

	// Read a bit then write a bit -- repeat until EOF
	reader := bufio.NewReader(inFile)
	writer := bufio.NewWriter(outFile)
	for {
		if slice, err := reader.ReadSlice('\n'); err == nil || err == io.EOF {
			writer.Write(slice)
			// IMPORTANT: Invoke flush()
			writer.Flush()
			if err == io.EOF {
				break
			}
		}
	}

}
