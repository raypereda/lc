package main

import (
	"bytes"
	"fmt"
	"io"
	"os"

	"github.com/raypereda/qnum"
)

func main() {
	args := os.Args[1:]

	if len(args) != 1 {
		fmt.Fprintln(os.Stderr, "Expected 1 argument: File path")
		os.Exit(1)
	}

	path := args[0]

	file, err := os.Open(path)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open file: %v", err)
		os.Exit(1)
	}

	count, err := lineCounter(file)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Encountered error while counting: %v", err)
		os.Exit(1)
	}

	fmt.Println(qnum.F(float64(count)))
}

func lineCounter(r io.Reader) (int, error) {
	buf := make([]byte, 32*1024)
	count := 0
	lineSep := []byte{'\n'}

	for {
		c, err := r.Read(buf)
		count += bytes.Count(buf[:c], lineSep)

		switch {
		case err == io.EOF:
			return count, nil

		case err != nil:
			return count, err
		}
	}
}
