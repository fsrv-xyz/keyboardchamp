package helper

import (
	"os"
	"os/exec"
)

func RunCommandDisplayZero(cmd *exec.Cmd) error {
	if _, displayAlreadySet := os.LookupEnv("DISPLAY"); !displayAlreadySet {
		cmd.Env = append(os.Environ(), "DISPLAY=:0")
	}
	if os.Getenv("DEBUG") != "" {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	}
	return cmd.Run()
}
