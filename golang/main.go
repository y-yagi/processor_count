package main

import (
	"fmt"
	"log"
	"runtime"

	"go.uber.org/automaxprocs/maxprocs"
)

func main() {
	fmt.Printf("runtime.NumCPU: %d\n", runtime.NumCPU())
	maxprocs.Set(maxprocs.Logger(log.Printf))
}
