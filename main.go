package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"

	"github.com/1l0/lab2spl/construct"
)

var (
	outputFilename string
	fps            int
)

func init() {
	log.SetFlags(log.Lshortfile)
	flag.StringVar(&outputFilename, "o", "", "Output file path.")
	flag.IntVar(&fps, "fps", 30, "FPS.")
	flag.Parse()
}

func main() {
	args := flag.Args()
	if len(args) < 1 {
		log.Fatalln(fmt.Errorf("missing .lab file path"))
	}
	labPath := args[0]
	exp, err := regexp.Compile(`.+\.(lab)$`)
	if err != nil {
		log.Fatalln(err)
	}
	m := exp.FindStringSubmatch(labPath)
	if len(m) < 2 {
		log.Fatalln(fmt.Errorf("unsupported .lab file"))
	}
	if outputFilename == "" {
		outputFilename = m[0] + ".spl"
	}

	spl, err := construct.Lab2Spl(labPath, fps)
	if err != nil {
		log.Fatalln(err)
	}
	f, err := os.OpenFile(outputFilename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalln(err)
	}
	defer f.Close()

	if err := f.Truncate(0); err != nil {
		log.Fatalln(err)
	}
	if _, err := f.Seek(0, 0); err != nil {
		log.Fatalln(err)
	}
	if _, err := f.WriteString(spl); err != nil {
		log.Fatalln(err)
	}
}
