package main

import (
	"fmt"
	"os"
)

func main() {
	// Suppose we wanted to create a file, write to it, and then close when we’re done.
	// Here’s how we could do that with defer.
	// Immediately after getting a file object with createFile,
	// we defer the closing of that file with closeFile. This will be executed
	// at the end of the enclosing function (main), after writeFile has finished.
	f := createFile("/tmp/defer.txt")
	defer closeFile(f)
	writeFile(f)
}

func createFile(p string) *os.File {
	fmt.Println("createing")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")
}

func closeFile(f *os.File) {
	fmt.Println("closing")
	f.Close()
}
