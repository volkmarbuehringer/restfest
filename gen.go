// The following directive is necessary to make the package coherent:

// +build ignore

// This program generates the code in gener It can be invoked by running
// go generate
package main

import "restfest/generator"

func main() {
	generator.Generator()
}
