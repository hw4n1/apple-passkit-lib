package passkit

// An object that represents a date interval that the system uses to show a relevant pass.
//
// See: https://developer.apple.com/documentation/walletpasses/pass/relevantdates-data.dictionary

type RelevantDates struct {

	// ---------------- OPTIONAL ----------------
	Date      *string `json:"date,omitempty"`
	StartDate *string `json:"startDate,omitempty"`
	EndDate   *string `json:"endDate,omitempty"`
}
