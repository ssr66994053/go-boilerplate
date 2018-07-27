package main

import (
	"fmt"

	"../../graceful"
)

func main() {
	graceful.Watch(func() {
		fmt.Println("shut...")
	})
}
