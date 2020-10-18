package main

import (
	"fmt"

	"github.com/code-pride/sweet.up/pkg/repository"
)

func main() {
	fmt.Println("Hello, world.")
	repository.NewMongoRepository("mongodb://localhost:27017")
}
