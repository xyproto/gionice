package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/xyproto/gionice"
)

func main() {
	// Make the current process "idle" (level 7)
	gionice.SetIdle(0)

	// Write to a file and delete then delete it, repeatedly
	for {
		fmt.Println("TICK")
		_ = ioutil.WriteFile("frenetic.dat", []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 0644)
		fmt.Println("TOCK")
		_ = os.Remove("frenetic.dat")
	}
}
