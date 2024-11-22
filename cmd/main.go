package main

import (
	"github.com/PolinaBY/lab6/internal/worker.go"
)

func main() {
	worker.RunWorkers(3)
}
