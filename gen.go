// The following directive is necessary to make the package coherent:

// +build ignore

// This program generates the code in gener It can be invoked by running
// go generate
package main

import (
	"fmt"
	"log"
	"restfest/generator"

	"github.com/jackc/pgx"
)

func main() {
	if err := generator.Generator(); err != nil {

		if pqErr, ok := err.(pgx.PgError); ok {
			log.Println(pqErr)
		}

		fmt.Println("Generation aborted", err)
	} else {
		fmt.Println("generation ok")
	}

}
