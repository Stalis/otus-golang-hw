package main

import (
	"log"
	"os"
)

func main() {

	env, err := ReadDir(os.Args[1])
	if err != nil {
		log.Fatalf("Unexpected error: %v", err)
	}

	res := RunCmd(os.Args[2:], env)
	os.Exit(res)
}
