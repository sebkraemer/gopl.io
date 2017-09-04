// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 8.

// Echo3 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
)

//!+
func main() {
	for i := 1; i <= len(os.Args[1:]); i++ {
		fmt.Printf("%d %s\n", i, os.Args[i])
	}

}

//!-
