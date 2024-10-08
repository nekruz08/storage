package main

import (
	"fmt"
	"github.com/nekruz08/storage/internal/storage"
)

func main() {
	st := storage.NewStorage()
	fmt.Println("it works", st)
}
