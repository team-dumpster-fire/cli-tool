package src

import (
	"fmt"
	"os"
	"os/exec"
)

// GoForth and prosper!
func GoForth() {
	cmd := exec.Command("docker", "run", "-it", "--rm", "carflo/dumpster-fire:latest")
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin

	err := cmd.Run()
	if err != nil {
		panic("OH NO!")
	}

	fmt.Println("Nice!")
}
