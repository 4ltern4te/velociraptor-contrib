//go build -o g.go file.bin

package main

import (
    "fmt"
	"os"
)
 
func main() {
	pid := os.Getpid()

    fmt.Println("executing payload")
	fmt.Println("Press the Enter Key to stop anytime")

	fmt.Printf("pid: %T, %v\n", pid, pid)
    fmt.Scanln()
}
