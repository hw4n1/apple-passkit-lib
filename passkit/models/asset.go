package models

// Asset represents a single file asset to be placed inside the .pass folder.
type Asset struct {
	Name string
	Data []byte
}
