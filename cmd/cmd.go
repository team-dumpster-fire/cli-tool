package cmd

import "fmt"

// GoForth and prosper!
func GoForth(cmd string, args ...string) error {
	switch cmd {
	case "party":
		return party()
	default:
		return fmt.Errorf("huh?")
	}
}
