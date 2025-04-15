package readme

import (
	"os"
	"strings"
)

func Update(readmePath, newBlock string) error {
	content, err := os.ReadFile(readmePath)
	if err != nil {
		return err
	}
	str := string(content)

	startTag := "<!-- BENCHSCRIBE:START -->"
	endTag := "<!-- BENCHSCRIBE:END -->"

	before := strings.Split(str, startTag)[0]
	after := strings.Split(str, endTag)[1]

	updated := before + startTag + "\n" + newBlock + "\n" + endTag + after
	return os.WriteFile(readmePath, []byte(updated), 0644)
}
