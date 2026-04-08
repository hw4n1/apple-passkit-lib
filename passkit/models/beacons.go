package models

//An object that represents the identifier of a Bluetooth Low Energy beacon the system uses to show a relevant pass.
//
//See: https://developer.apple.com/documentation/walletpasses/pass/beacons-data.dictionary

type Beacons struct {
	// ---------------- REQUIRED ----------------
	ProximityUUID string `json:"proximityUUID"`

	// ---------------- OPTIONAL ----------------
	Major        *uint16 `json:"major,omitempty"`
	Minor        *uint16 `json:"minor,omitempty"`
	RelevantText *string `json:"relevantText,omitempty"`
}
