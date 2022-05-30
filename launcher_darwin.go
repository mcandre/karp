//go:build darwin
package karp

import (
	"os"
	"os/exec"
)

// Launch triggers a graphical hook for opening URIs.
func Launch(pth string) error {
	cmd := exec.Command("open", pth)
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
