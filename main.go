// main.go

package main

import (
	"fmt"
	"log"

	"github.com/q0d1r0v/go-translator-cli/cmd"
)

func main() {
	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}
}
