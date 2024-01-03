//go:build !darwin

package karp

import (
	"os"
	"os/exec"
	"regexp"
)

// linuxLaunch triggers a graphical hook for opening URIs.
func linuxLaunch(pth string) error {
	cmd := exec.Command("xdg-open", pth)
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// windowsLaunch triggers a graphical hook for opening URIs.
func windowsLaunch(pth string) error {
	cmd := exec.Command("powershell.exe", "Start", pth)
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// brandPattern identifies COMSPEC/WSL environments.
var brandPattern = regexp.MustCompile("[Mm]icrosoft")

// Launch triggers a graphical hook for opening URIs.
func Launch(pth string) error {
	versionPath := "/proc/version"

	versionBytes, err := os.ReadFile(versionPath)

	if err != nil || !brandPattern.Match(versionBytes) {
		return linuxLaunch(pth)
	}

	return windowsLaunch(pth)
}
