package main

import (
	"flag"
	"log"

	"github.com/syywu/geojson-merger/geojson"
)

func main() {
	var outputFilename string
	flag.StringVar(&outputFilename, "output", "output.geojson", "output geojson filename")
	flag.Parse()

	inputFiles := flag.Args()
	if len(inputFiles) < 1 {
		log.Fatal("input files are required")
	}

	fc, err := geojson.Merge(inputFiles)
	if err != nil {
		log.Fatal(err)
	}

	err = geojson.Output(flag.Arg(0), outputFilename, fc)
	if err != nil {
		log.Fatal(err)
	}
}
