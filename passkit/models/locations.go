package models

//An object that represents a location that the system uses to show a relevant pass.
//
//See: https://developer.apple.com/documentation/walletpasses/pass/locations-data.dictionary

type Locations struct {
	// ---------------- REQUIRED ----------------
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`

	// ---------------- OPTIONAL ----------------
	Altitude     *float64 `json:"altitude,omitempty"`
	RelevantText *string  `json:"relevantText,omitempty"`
}
