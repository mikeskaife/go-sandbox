package main

import (
	"fmt"
	"os"
)

func main() {
	name := os.Getenv("LOGNAME")

	fmt.Println(name)
}
