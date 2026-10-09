//go:build !windows

package mounted

// GetWindowsDriveLetters is only meaningful on Windows, where the directory
// tree shows the drives; the caller never asks for them elsewhere.
func GetWindowsDriveLetters() ([]string, error) {
	return nil, nil
}
