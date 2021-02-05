package main

import(
	"fmt"
	"os"
	"io/ioutil"
)

func main() {
	outExt := ".out"
	cmpExt := ".cmp"

	path := os.Args[1]
	

	outPath := path+outExt
	cmpPath := path+cmpExt
	fmt.Printf("opening files: %s, %s\n", outPath, cmpPath)

	// Read file
	out, err := ioutil.ReadFile(outPath)
	if err != nil {
		panic(err)
	}
	// fmt.Println(string(out))

	cmp, err := ioutil.ReadFile(cmpPath)
	if err != nil {
		panic(err)
	}
	// fmt.Println(string(cmp))

	if string(out) == string(out)  {
		fmt.Printf("%s.hdl implemented correctly\n", path)
	} else {
		fmt.Printf("%s:\n%s\n%s:\n%s", outPath, string(out), cmpPath, string(cmp))
	}
	
}

