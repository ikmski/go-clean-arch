package main

import (
	"fmt"
	"github.com/ikmski/go-clean-arch/database"
)

func main() {
	fmt.Println("hello")

	database.Migration()
}
