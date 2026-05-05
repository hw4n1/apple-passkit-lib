package builder

import (
	"archive/zip"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"com.naing/apple-passkit/passkit/models"
)

// BuildPass orchestrates and assembles a complete .pass bundle for the given passName.
// Creates pass elements (root json, assets and manifest) and signs it
func BuildPass(passName string, pass models.Pass, assets []models.Asset) error {
	if err := build(passName, pass); err != nil {
		slog.Error("build failed", "err", err)
		return err
	}

	if len(assets) == 0 {
		slog.Info("no assets to write")
		return nil
	}

	if err := writeAssets(passName, assets); err != nil {
		slog.Error("writing assets failed", "err", err)
		return err
	}

	if err := GenPassManifest(passName); err != nil {
		slog.Error("manifest generation failed", "err", err)
		return err
	}

	if false == SignPass(passName) {
		return fmt.Errorf("pass signing failed")
	}

	if err := zipPass(passName); err != nil {
		slog.Error("zipping pass failed", "err", err)
		return err
	}

	return nil
}

// BuildPassBytes orchestrates and assembles a complete .pkpass bundle for the given pass
// and returns the bytes of the final zipped .pkpass file instead of writing to disk.
// Uses temporary files internally for the build process but returns only the zipped bytes.
func BuildPassBytes(pass models.Pass, assets []models.Asset) ([]byte, error) {
	// Create a temporary directory for intermediate files
	tmpDir, err := os.MkdirTemp("", "passkit-*")
	if err != nil {
		slog.Error("failed to create temp directory", "err", err)
		return nil, err
	}
	defer os.RemoveAll(tmpDir)

	// Use a temporary pass name based on the temp directory
	tempPassName := filepath.Join(tmpDir, "temp-pass")

	// Build the pass using existing functions
	if err := build(tempPassName, pass); err != nil {
		slog.Error("build failed", "err", err)
		return nil, err
	}

	if len(assets) > 0 {
		if err := writeAssets(tempPassName, assets); err != nil {
			slog.Error("writing assets failed", "err", err)
			return nil, err
		}
	}

	if err := GenPassManifest(tempPassName); err != nil {
		slog.Error("manifest generation failed", "err", err)
		return nil, err
	}

	if !SignPass(tempPassName) {
		return nil, fmt.Errorf("pass signing failed")
	}

	// Create the pkpass in memory by zipping the temporary directory contents
	passDir := tempPassName + ".pass"
	entries, err := os.ReadDir(passDir)
	if err != nil {
		return nil, fmt.Errorf("reading pass directory: %w", err)
	}

	buf := new(bytes.Buffer)
	w := zip.NewWriter(buf)

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		src, err := os.Open(filepath.Join(passDir, entry.Name()))
		if err != nil {
			w.Close()
			return nil, fmt.Errorf("opening %s: %w", entry.Name(), err)
		}

		dst, err := w.Create(entry.Name())
		if err != nil {
			src.Close()
			w.Close()
			return nil, fmt.Errorf("adding %s to archive: %w", entry.Name(), err)
		}

		if _, err = io.Copy(dst, src); err != nil {
			src.Close()
			w.Close()
			return nil, fmt.Errorf("writing %s to archive: %w", entry.Name(), err)
		}

		src.Close()
	}

	if err := w.Close(); err != nil {
		return nil, fmt.Errorf("closing zip writer: %w", err)
	}

	slog.Info("created pkpass bytes", "size", buf.Len())
	return buf.Bytes(), nil
}


// writeAssets writes given assets into the <passName>.pass directory.
// Assets with an empty Name field are skipped with a warning. It returns an
// error if the pass directory does not exist or any file cannot be written.
func writeAssets(passName string, assets []models.Asset) error {
	dir := passName + ".pass"

	if stat, err := os.Stat(dir); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return fmt.Errorf("pass directory does not exist: %s", dir)
		}
		return err
	} else if !stat.IsDir() {
		return fmt.Errorf("path exists but is not a directory: %s", dir)
	}

	for _, a := range assets {
		if a.Name == "" {
			slog.Warn("skipping asset with empty name")
			continue
		}

		path := filepath.Join(dir, a.Name)
		if err := os.WriteFile(path, a.Data, 0644); err != nil {
			slog.Error("writing asset failed", "asset", a.Name, "err", err)
			return err
		}

		slog.Info("wrote asset", "asset", a.Name, "path", path)
	}

	return nil
}

// build creates the <passName>.pass directory (or updates it if it already
// exists), validates and fills in default field values via ensureRequiredFields,
// and writes the encoded pass as pass.json inside that directory.
func build(passName string, pass models.Pass) error {
	// Validate and set defaults
	if err := ensureRequiredFields(&pass); err != nil {
		return err
	}

	dirName := passName + ".pass"
	err := os.Mkdir(dirName, 0750)

	if err != nil {
		if !errors.Is(err, os.ErrExist) {
			slog.Error("Error creating dir", "err", err)
			return err
		}

		slog.Warn("Directory " + dirName + " already exists. Overwriting pass")
	}

	filePath := filepath.Join(dirName, "pass.json")
	file, err := os.Create(filePath)

	if err != nil {
		slog.Error("Error creating file", "err", err)
		return err
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			slog.Error("Error closing file", "err", err)
		}
	}(file)

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(pass)

	if err != nil {
		slog.Error("Error encoding pass to file", "err", err)
		return err
	}

	slog.Info("Built pass", "dir", dirName, "serial", pass.SerialNumber)
	return nil
}

// zipPass archives the contents of <passName>.pass/ directly into <passName>.pkpass,
// placing all files at the root of the archive (no directory prefix).
func zipPass(passName string) error {
	passDir := passName + ".pass"
	pkpassPath := passName + ".pkpass"

	entries, err := os.ReadDir(passDir)
	if err != nil {
		return fmt.Errorf("reading pass directory: %w", err)
	}

	out, err := os.Create(pkpassPath)
	if err != nil {
		return fmt.Errorf("creating pkpass file: %w", err)
	}
	defer out.Close()

	w := zip.NewWriter(out)
	defer w.Close()

	for _, entry := range entries {
		if entry.IsDir() {
			continue
		}

		src, err := os.Open(filepath.Join(passDir, entry.Name()))
		if err != nil {
			return fmt.Errorf("opening %s: %w", entry.Name(), err)
		}

		dst, err := w.Create(entry.Name())
		if err != nil {
			src.Close()
			return fmt.Errorf("adding %s to archive: %w", entry.Name(), err)
		}

		if _, err = io.Copy(dst, src); err != nil {
			src.Close()
			return fmt.Errorf("writing %s to archive: %w", entry.Name(), err)
		}

		src.Close()
	}

	slog.Info("created pkpass", "path", pkpassPath)
	return nil
}

// ensureRequiredFields validates pass fields and gives default values for optional
// fields. FormatVersion defaults to 1, SerialNumber to a nanosecond-precision
// timestamp, Description to "Pass", and OrganizationName to "Organization"
// If no pass style is set, Generic is used as the default. It returns an error if
// PassTypeIdentifier or TeamIdentifier is absent, or if more than one pass
// style is provided.
func ensureRequiredFields(p *models.Pass) error {
	if p.FormatVersion == 0 {
		p.FormatVersion = 1
		slog.Debug("Set default FormatVersion to 1")
	}

	if p.SerialNumber == "" {
		// simple unique serial using timestamp
		p.SerialNumber = "SN-" + strconv.FormatInt(time.Now().UnixNano(), 10)
		slog.Debug("Generated SerialNumber", "serial", p.SerialNumber)
	}

	if p.Description == "" {
		p.Description = "Pass"
		slog.Debug("Set default Description")
	}

	if p.OrganizationName == "" {
		p.OrganizationName = "Organization"
		slog.Debug("Set default OrganizationName")
	}

	// Required Apple Developer identifiers
	if p.PassTypeIdentifier == "" {
		return fmt.Errorf("missing required field: PassTypeIdentifier")
	}
	if p.TeamIdentifier == "" {
		return fmt.Errorf("missing required field: TeamIdentifier")
	}

	// Ensure exactly one style is present. If none, default to `generic`.
	styleCount := 0
	if p.BoardingPass != nil {
		styleCount++
	}
	if p.Coupon != nil {
		styleCount++
	}
	if p.EventTicket != nil {
		styleCount++
	}
	if p.Generic != nil {
		styleCount++
	}
	if p.StoreCard != nil {
		styleCount++
	}

	if styleCount == 0 {
		p.Generic = &models.PassFields{}
		slog.Debug("No pass style provided; defaulted to generic")
	}

	if styleCount > 1 {
		return fmt.Errorf("exactly one pass style must be present")
	}

	return nil
}
