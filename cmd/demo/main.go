package main

import (
	"fmt"
	"os"

	"github.com/claesp/demo/internal/auth"
)

var (
	Version string = "0.0.0"
)

func main() {
	fmt.Fprintf(os.Stdout, "v%s\n", Version)

	authErr := auth.Authenticate()
	if authErr != nil {
		fmt.Fprintf(os.Stderr, "%s\n", authErr.Error())
		os.Exit(1)
	}
}
