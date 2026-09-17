package evasion

import (
	"fmt"
	"os/exec"
)

// Obfuscate uses shroud to obfuscate the agent payload.
func Obfuscate(inputPath, outputPath string, encoder string) error {
	// ponytail: calls shroud as an external binary.
	cmd := exec.Command("shroud", "-i", inputPath, "-o", outputPath, "-e", encoder)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("shroud failed: %w", err)
	}
	return nil
}
