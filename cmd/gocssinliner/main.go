package main

import (
	"flag"
	"log"
	"os"

	"github.com/astockwell/css_inliner"
)

func main() {
	inFilePtr := flag.String("i", "", "In-file (to be parsed)")
	outFilePtr := flag.String("o", "", "Out-file name")
	flag.Parse()
	inFile := *inFilePtr
	outFile := *outFilePtr

	if inFile == "" {
		log.Fatal("No In-file provided")
	}
	if outFile == "" {
		log.Fatal("No Out-file provided")
	}

	// Prep In-file
	src, err := os.Open(inFile)
	if err != nil {
		log.Fatal(err)
	}
	defer src.Close()

	// Prep Out-file
	dst, err := os.OpenFile(outFile, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.FileMode(0644))
	if err != nil {
		log.Fatal(err)
	}
	defer dst.Close()

	i := cssInliner.NewInliner()
	err = i.Inline(dst, src)
	if err != nil {
		log.Fatal(err)
	}
}
