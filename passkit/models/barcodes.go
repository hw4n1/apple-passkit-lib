package models

type Barcodes struct {

	// ---------------- REQUIRED ----------------
	Format          BarcodeFormat `json:"format"`
	Message         string        `json:"message"`
	MessageEncoding string        `json:"messageEncoding"`

	// ---------------- OPTIONAL ----------------
	AltText *string `json:"altText,omitempty"`
}
