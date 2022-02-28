package command

import (
	"errors"
	"os/exec"
)

func Run(commad string) (string, error) {
	out, err := exec.Command("sh", "-c", commad).Output()
	if err != nil {
		return "", errors.New("[Error]Cannot run commands: " + commad)
	}

	if len(out) == 0 {
		return "", nil
	}
	if out[len(out)-1] == 10 {
		out = out[:len(out)-1]
	}
	return string(out), nil
}
