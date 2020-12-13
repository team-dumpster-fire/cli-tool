package cmd

import (
	"fmt"
	"os"
	"os/exec"
)

func party() error {
	dir := os.TempDir() + "/dumpster-fire/carflo"
	if err := os.MkdirAll(dir, 0755); err != nil {
		return fmt.Errorf("could not make temporary directory: %w", err)
	}
	defer os.RemoveAll(dir)

	cmd := exec.Command("git", "clone", "https://github.com/carflo/party-dumpster-fire.git")
	cmd.Dir = dir
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("missing carflo powers! %w", err)
	}

	// Carlos why didn't you Go mod!? :_(
	if _, err := os.Stat(dir + "/party-dumpster-fire/dumpster-fire/go.mod"); os.IsNotExist(err) {
		cmd = exec.Command("go", "mod", "init", "party-dumpster-fire")
		cmd.Dir = dir + "/party-dumpster-fire/dumpster-fire"
		cmd.Stderr = os.Stderr
		cmd.Env = append(os.Environ(), "GO111MODULE=on")
		if err := cmd.Run(); err != nil {
			return fmt.Errorf("carflo's stuff couldn't get its modules on! %w", err)
		}
	}

	cmd = exec.Command("go", "build", "-o", "dumpster-fire")
	cmd.Dir = dir + "/party-dumpster-fire/dumpster-fire"
	cmd.Stderr = os.Stderr
	cmd.Env = append(os.Environ(), "GO111MODULE=on")
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("carflo's stuff couldn't build! %w", err)
	}

	cmd = exec.Command("./dumpster-fire")
	cmd.Dir = dir + "/party-dumpster-fire/dumpster-fire"
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("carflo's stuff couldn't run! %w", err)
	}

	fmt.Println("Nice!")
	return nil
}
