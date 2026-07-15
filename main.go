package main

import (
	"fmt"
	"postgres/feature/postgres/simple_connection"
)

func main() {
	fmt.Println("Hello, world!")

	simple_connection.CheckConnection()
}
