package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (r rot13Reader) Read(b []byte) (n int, e error) {
	n, e = r.r.Read(b)
	if e == nil {
		// This is messy but it works...
		for i := 0; i < n; i++ {
			if 'A' <= b[i] && b[i] <= 'Z' || 'a' <= b[i] && b[i] <= 'z' {
				b[i] += 13
				if b[i] > 'z' {
					b[i] -= 26
				}
			}
		}
	}
	return
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
