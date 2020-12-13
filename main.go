package main

import (
	"dumpster-fire/cmd"
	"log"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Give me something to do!")
	}

	if err := cmd.GoForth(os.Args[1], os.Args[2:]...); err != nil {
		log.Fatal(err)
	}
}
