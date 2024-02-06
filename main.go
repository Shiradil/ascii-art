package main

import (
	"ascii/logic"
	"ascii/validation"
	"log"
	"os"
)

func main() {
	// args validation
	if len(os.Args) == 1 || len(os.Args) > 3 {
		log.Fatal("No arguments or too many arguments")
	}

	// style modification validation
	if err := validation.CheckStyleModification(); err != nil {
		log.Fatal(err)
	}

	logic.PrintAscii(os.Args)
}
