package utils

import (
	"bytes"
	"encoding/json"
)

// PrettyPrintJSON indents JSON for human-friendly output
func PrettyPrintJSON(raw []byte) string {
	var out bytes.Buffer
	err := json.Indent(&out, raw, "", "  ")
	if err != nil {
		return string(raw) // fallback to raw output
	}
	return out.String()
}
