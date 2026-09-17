package modules

// Keylogger starts recording keystrokes.
func Keylogger(durationSeconds int) (string, error) {
	// ponytail: this requires low-level OS hooks (e.g. User32.dll on Win, X11 on Linux).
	// Returning a mock log for now.
	return "MOCK_KEYLOG_DATA: 'password123'", nil
}
