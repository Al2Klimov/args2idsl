package main

import (
	"bytes"
	"io"
	"os"
)

func main() {
	var buf bytes.Buffer

	buf.Write([]byte("[\n"))

	for _, arg := range os.Args {
		parts := bytes.Split([]byte(arg), []byte("}}}"))
		for i := range parts {
			parts[i] = append(append([]byte("{{{"), parts[i]...), []byte("}}}")...)
		}

		buf.Write([]byte("  "))
		buf.Write(bytes.Join(parts, []byte(` + "}}}" + `)))
		buf.Write([]byte(",\n"))
	}

	buf.Write([]byte("]\n"))

	io.Copy(os.Stdout, &buf)

	os.Exit(4)
}
