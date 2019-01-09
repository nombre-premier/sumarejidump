package main

import (
	"log"
	"os"
)

func main() {
	cliApp := CreateCliApp()
	err := cliApp.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
