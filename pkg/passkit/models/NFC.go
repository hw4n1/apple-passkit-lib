package passkit

// An object that represents the near-field communication (NFC) payload the device passes to an Apple Pay terminal.
//
// See: https://developer.apple.com/documentation/walletpasses/pass/nfc-data.dictionary

type NFC struct {
	// ---------------- REQUIRED ----------------
	EncryptionPublicKey string `json:"encryptionPublicKey"`
	Message             string `json:"message"`

	// ---------------- OPTIONAL ----------------
	RequiresAuthentication *bool `json:"requiresAuthentication,omitempty"`
}
