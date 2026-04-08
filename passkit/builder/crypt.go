package builder

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"io"
	"log/slog"
	"os"
)

// fileSHA1 computes the SHA1 hash of the file at filePath and returns it as a
// lowercase hexadecimal string.
func fileSHA1(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	hash := sha1.New()
	if _, err := io.Copy(hash, file); err != nil {
		return "", err
	}

	hashedData := hash.Sum(nil)
	hashedString := hex.EncodeToString(hashedData)

	return hashedString, nil
}

// GenPassManifest generates the manifest.json file for the pass bundle
// identified by passName. It computes the SHA1 hash of every file in the
// <passName>.pass directory — excluding any pre-existing manifest.json — and
// writes the resulting filename-to-hash map to manifest.json inside that
// directory.
func GenPassManifest(passName string) error {
	dirName := passName + ".pass"
	files, err := os.ReadDir(dirName)
	if err != nil {
		slog.Error("Error reading pass files", "err", err)
	}

	manifestMap := make(map[string]string)

	for _, file := range files {
		if file.Name() == "manifest.json" {
			continue
		}
		hash, err := fileSHA1(dirName + "/" + file.Name())
		if err != nil {
			return err
		}
		manifestMap[file.Name()] = hash
	}

	manifestFile, err := os.Create(dirName + "/manifest.json")
	if err != nil {
		slog.Error("Error creating manifest file", "err", err)
		return err
	}
	defer manifestFile.Close()

	encoder := json.NewEncoder(manifestFile)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(manifestMap); err != nil {
		slog.Error("Error encoding manifest to JSON", "err", err)
		return err
	}

	slog.Debug("Successfully wrote JSON data to manifest.json")
	return nil
}
