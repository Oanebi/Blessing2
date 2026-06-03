package ascii

import (
	"errors"
	"strings"
)

// GenerateArt takes text and a style name, and outputs simulated ASCII line blocks
func GenerateArt(text string, banner string) (string, error) {
	// Simple validation to ensure the banner template asset exists
	if banner != "standard" && banner != "shadow" && banner != "thinkertoy" {
		return "", errors.New("banner file not found")
	}

	// Simple simulation of your text block stacking logic
	var simulatedArt string
	lines := strings.Split(text, "\n")

	for _, line := range lines {
		if line == "" {
			continue
		}
		// Simulate drawing lines for characters side-by-side
		simulatedArt += " [Row 1: " + line + " Layout (" + banner + ")] \n"
		simulatedArt += " [Row 2: " + line + " Layout (" + banner + ")] \n"
		simulatedArt += " [Row 3: " + line + " Layout (" + banner + ")] \n"
	}

	return simulatedArt, nil
}
