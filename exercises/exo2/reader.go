package main

import (
	"io"
	"os"
	"strings"
)

type spaceEraser struct {
	r io.Reader
}

func (reader spaceEraser) Read(buffer []byte) (int, error) {
	n, err := reader.r.Read(buffer)
	char := 0
	for i := 0; i < n; i++ {
		if buffer[i] != 32 {
			buffer[char] = buffer[i]
			char++
		}
	}
	return char, err
}

func main() {
	s := strings.NewReader("H e l l o w o r l d!")
	r := spaceEraser{s}
	io.Copy(os.Stdout, &r)
}

// Implement a type that satisfies the io.Reader interface and reads from another io.Reader,
// modifying the stream by removing the spaces.
// Expected output: "Helloworld!"
