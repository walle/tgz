package main

import (
	"fmt"
	"os"

	"github.com/walle/targz"
)

const (
	COMPRESS = iota
	EXTRACT
	NOOP
)

func main() {
	args := os.Args
	argsLen := len(args)
	inputFile := ""
	outputFile := ""
	mode := NOOP
	if argsLen == 2 {
		inputFile = args[1]
		mode = EXTRACT
	} else if argsLen == 3 {
		inputFile = args[1]
		outputFile = args[2]
	} else {
		printUsage()
	}

	fi, err := os.Stat(inputFile)
	if err != nil {
		fmt.Fprintf(os.Stderr, "No such file: %s\n", inputFile)
		printUsage()
	}
	if fi.IsDir() {
		if outputFile == "" {
			fmt.Fprintf(os.Stderr, "input_file is a directory, you must specify an output_file to compress to")
			printUsage()
		} else {
			mode = COMPRESS
		}
	} else {
		mode = EXTRACT
	}

	switch mode {
	case COMPRESS:
		err := targz.Compress(inputFile, outputFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error compressing: %s\n", err)
			os.Exit(1)
		}
	case EXTRACT:
		err = targz.Extract(inputFile, outputFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error extracting: %s\n", err)
			os.Exit(1)
		}
	}

	os.Exit(0)
}

func printUsage() {
	fmt.Fprintf(os.Stderr, "usage: %s input_file [output_file]\n", os.Args[0])
	fmt.Fprintf(os.Stderr, "  input_file: if input_file is a file %s tries to extract it, if input_file is a directory %s tries to compress it to output_file\n", os.Args[0], os.Args[0])
	fmt.Fprintf(os.Stderr, "  output_file: only required if input_file is a directory, otherwise the current directory is used\n")
	os.Exit(1)
}
