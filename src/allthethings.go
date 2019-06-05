package src

import (
	"fmt"
	"os"
	"os/exec"
)

// GoForth and prosper!
func GoForth() {
	cmd := exec.Command("docker", "run", "-it", "--rm", "jmhobbs/terminal-parrot:latest", "--loops", "10")
	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin

	err := cmd.Run()
	if err != nil {
		panic("OH NO!")
	}

	fmt.Println("Nice!")
}
