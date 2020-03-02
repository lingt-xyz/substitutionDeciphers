package main

import (
	"flag"
	"io/ioutil"
	"log"
	"runtime"
)

func main() {
	log.Printf("Number of CPUs: %v\n", runtime.NumCPU())

	testFile := flag.String("input", "test", "full path of the file for testing")
	flag.Parse()

	content, err := ioutil.ReadFile(*testFile)
	if err != nil {
		log.Panicf("Cannot open the file '%v', please check its existence and the reading permission is given.", *testFile)
	}
	log.Printf("%v", string(content))
}