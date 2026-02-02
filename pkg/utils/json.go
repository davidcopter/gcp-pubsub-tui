package utils

import (
	"bytes"
	"encoding/json"
)

// PrettyPrintJSON formats a byte slice of JSON into a pretty-printed string.
// If the input is not valid JSON, it returns the string representation of the input.
func PrettyPrintJSON(input []byte) string {
	var prettyJSON bytes.Buffer
	if err := json.Indent(&prettyJSON, input, "", "  "); err != nil {
		return string(input)
	}
	return prettyJSON.String()
}
