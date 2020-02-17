package main

import (
	"bufio"
	"io"
	"os"
	"path/filepath"
)

func getInput(name string) *bufio.Reader {
	file, err := os.Open(filepath.Join("input", name))
	if err != nil {
		panic(err)
	}

	return bufio.NewReaderSize(file, 2<<30)
}

func getLinesFromInput(name string) (lines []string) {
	rd := getInput(name)
	var content []byte
	var prefix bool
	var err error
	for {
		content, prefix, err = rd.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		if prefix {
			panic("prefix not supported yet")
		}
		lines = append(lines, string(content))
	}
	return
}
