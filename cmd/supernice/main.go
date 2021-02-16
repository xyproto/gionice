package main

import (
	"io/ioutil"
	"os"

	"github.com/xyproto/gionice"
)

func main() {
	// Set the current process group CPU niceness to 10
	_ = gionice.Nice()

	// Set the current process group IO priority to "idle" (level 7)
	_ = gionice.Idle()

	// Generate I/O activity
	for {
		_ = ioutil.WriteFile("frenetic.dat", []byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 0644)
		_ = os.Remove("frenetic.dat")
	}
}
