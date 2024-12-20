package util

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func ReadFile(filePath string, value any) error {
	file, err := os.Open(filePath)
	if err != nil {
		return fmt.Errorf("error on open file: %v", err)
	}

	b, err := io.ReadAll(file)
	if err != nil {
		return fmt.Errorf("error on read file: %v", err)
	}

	if err := json.Unmarshal(b, value); err != nil {
		return fmt.Errorf("error on Unmarshal JSON: %v", err)
	}

	return nil
}
