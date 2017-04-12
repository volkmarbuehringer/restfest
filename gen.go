// The following directive is necessary to make the package coherent:

// +build ignore

// This program generates the code in gener It can be invoked by running
// go generate
package main

import (
	"fmt"
	"restfest/generator"
)

func main() {
	if err := generator.Generator(); err != nil {
		fmt.Println("Generation aborted", err)
	} else {
		fmt.Println("generation ok")
	}

}
