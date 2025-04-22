package readme

import (
	"fmt"
	"os"
	"strings"
)

const (
	startTag = "<!-- BENCHSCRIBE:START -->"
	endTag   = "<!-- BENCHSCRIBE:END -->"
)

func Update(readmePath, newBlock string) error {
	content, err := os.ReadFile(readmePath)
	if err != nil {
		return fmt.Errorf("failed to read README: %w", err)
	}
	str := string(content)

	// Check if markers exist
	if !strings.Contains(str, startTag) || !strings.Contains(str, endTag) {
		return fmt.Errorf("README markers not found. Please add %s and %s to your README", startTag, endTag)
	}

	// Split content
	parts := strings.Split(str, startTag)
	if len(parts) != 2 {
		return fmt.Errorf("invalid README format: multiple start markers found")
	}
	before := parts[0]

	parts = strings.Split(parts[1], endTag)
	if len(parts) != 2 {
		return fmt.Errorf("invalid README format: multiple end markers found")
	}
	after := parts[1]

	// Construct updated content
	updated := before + startTag + "\n" + newBlock + "\n" + endTag + after

	// Write with more permissive permissions
	if err := os.WriteFile(readmePath, []byte(updated), 0666); err != nil {
		return fmt.Errorf("failed to write README: %w", err)
	}

	return nil
}
