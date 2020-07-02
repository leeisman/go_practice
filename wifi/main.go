package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {

	cmd := exec.Command("networksetup", "-setairportnetwork", "en0", "Silkrode OA", "Silkrode50832747")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	fmt.Printf("combined out:\n%s\n", string(out))
}
