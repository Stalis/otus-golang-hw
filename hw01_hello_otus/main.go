package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

func main() {
	reversed := stringutil.Reverse("Hello, OTUS!")
	fmt.Println(reversed)
}
