package passkit

// A compilation of data object types for semantic tags.
// Types that embed SemanticTagTypes: CurrencyAmount, EventDateInfo, Location, PersonNameComponents, Seat, WifiNetwork.
//
// See: https://developer.apple.com/documentation/walletpasses/semantictagtype

type SemanticTagTypes struct{}

// An object that represents the coordinates of a location.
// See: https://developer.apple.com/documentation/walletpasses/semantictagtype/location-data.dictionary
type Location struct {
	SemanticTagTypes
	// ---------------- REQUIRED ----------------
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// CurrencyAmount represents an amount of money and type of currency.
// See: https://developer.apple.com/documentation/walletpasses/semantictagtype/currencyamount-data.dictionary
type CurrencyAmount struct {
	SemanticTagTypes
	Amount       string `json:"amount"`
	CurrencyCode string `json:"currencyCode"`
}

// An object that represents a date for an event.
// See: https://developer.apple.com/documentation/walletpasses/semantictagtype/eventdateinfo-data.dictionary
type EventDateInfo struct {
	SemanticTagTypes
	Date                 string  `json:"date"`
	IgnoreTimeComponents *bool   `json:"ignoreTimeComponents,omitempty"`
	TimeZone             *string `json:"timeZone,omitempty"`
	Unannounced          *bool   `json:"unannounced,omitempty"`
	Undetermined         *bool   `json:"undetermined,omitempty"`
}

// An object that represents the parts of a person’s name.
// See: https://developer.apple.com/documentation/walletpasses/semantictagtype/personnamecomponents-data.dictionary
type PersonNameComponents struct {
	SemanticTagTypes
	FamilyName             *string `json:"familyName,omitempty"`
	GivenName              *string `json:"givenName,omitempty"`
	MiddleName             *string `json:"middleName,omitempty"`
	NamePrefix             *string `json:"namePrefix,omitempty"`
	NameSuffix             *string `json:"nameSuffix,omitempty"`
	Nickname               *string `json:"nickname,omitempty"`
	PhoneticRepresentation *string `json:"phoneticRepresentation,omitempty"`
}

// A compilation of data object types for semantic tags.
// See: https://developer.apple.com/documentation/walletpasses/semantictagtype/seat-data.dictionary
type Seat struct {
	SemanticTagTypes
	SeatAisle        *string `json:"seatAisle,omitempty"`
	SeatDescription  *string `json:"seatDescription,omitempty"`
	SeatIdentifier   *string `json:"seatIdentifier,omitempty"`
	SeatLevel        *string `json:"seatLevel,omitempty"`
	SeatNumber       *string `json:"seatNumber,omitempty"`
	SeatRow          *string `json:"seatRow,omitempty"`
	SeatSection      *string `json:"seatSection,omitempty"`
	SeatSectionColor *string `json:"seatSectionColor,omitempty"`
	SeatType         *string `json:"seatType,omitempty"`
}

// An object that contains information required to connect to a Wi-Fi network. Optionally, this object may contain keys required to perform authentication with captive portal.
// See: https://developer.apple.com/documentation/walletpasses/semantictagtype/wifinetwork-data.dictionary
type WifiNetwork struct {
	SemanticTagTypes
	// ---------------- REQUIRED ----------------
	Password string `json:"password"`
	SSID     string `json:"ssid"`

	// ---------------- OPTIONAL ----------------
	CaptiveToken        *string `json:"captiveToken,omitempty"`
	CaptiveTokenAuthURL *string `json:"captiveTokenAuthURL,omitempty"`
}
