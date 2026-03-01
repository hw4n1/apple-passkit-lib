package passkit

// Pass represents the root pass.json object.
//
// Apple Documentation:
// https://developer.apple.com/documentation/walletpasses/pass
//
// Exactly ONE pass style must be present.
type Pass struct {

	// ---------------- REQUIRED ----------------

	Description        string `json:"description"`
	FormatVersion      int    `json:"formatVersion"`
	OrganizationName   string `json:"organizationName"`
	PassTypeIdentifier string `json:"passTypeIdentifier"`
	SerialNumber       string `json:"serialNumber"`
	TeamIdentifier     string `json:"teamIdentifier"`

	// ---------------- OPTIONAL ----------------

	LogoText *string `json:"logoText,omitempty"`
	Voided   *bool   `json:"voided,omitempty"`

	// ---------------- PASS STYLES ----------------
	BoardingPass *BoardingPass `json:"boardingPass,omitempty"`
	Coupon       *Coupon       `json:"coupon,omitempty"`
	EventTicket  *EventTicket  `json:"eventTicket,omitempty"`
	Generic      *Generic      `json:"generic,omitempty"`
	StoreCard    *StoreCard    `json:"storeCard,omitempty"`
}