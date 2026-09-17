package modules

// Screenshot captures the current screen and returns the image as bytes.
func Screenshot() ([]byte, error) {
	// ponytail: this requires CGO or external libs like robotgo.
	// Returning a mock image for now.
	return []byte("MOCK_SCREENSHOT_DATA"), nil
}
