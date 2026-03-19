package builder

import (
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"com.naing/apple-passkit/passkit/models"
)

/*
Validates the provided pass
Ensures pass style is given
Writes the result json file
*/
func Build(passName string, pass models.Pass) error {
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

		slog.Warn("Directory " + dirName + " already exists")
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

/*
Verifies required & default values and identifiers
*/
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
