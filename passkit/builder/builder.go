package builder

import (
	"encoding/json"
	"errors"
	"log/slog"
	"os"
	"path/filepath"

	"com.naing/apple-passkit/passkit/models"
)

func BuildSource(passName string, pass models.Pass) {
	dirName := passName + ".pass"
	err := os.Mkdir(dirName, 0750)

	if err != nil {
		if !errors.Is(err, os.ErrExist) {
			slog.Error("Error creating dir", err)
			return
		}

		slog.Warn("Directory " + dirName + " already exists")
	}

	filePath := filepath.Join(dirName, "pass.json")
	file, err := os.Create(filePath)

	if err != nil {
		slog.Error("Error creating file", err)
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			slog.Error("Error closing file", err)
		}
	}(file)

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(pass)

	if err != nil {
		slog.Error("Error encoding pass to file", err)
	}
}
