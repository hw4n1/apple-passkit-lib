package models

// SemanticTags contains machine-readable metadata the system uses to offer a pass and suggest related actions.
//
// See: https://developer.apple.com/documentation/walletpasses/semantictags

// SemanticTags is the semantics object in the pass.
type SemanticTags struct {
	// ---------- Event / ticket (localizable strings) ----------
	AdditionalTicketAttributes *string   `json:"additionalTicketAttributes,omitempty"`
	AdmissionLevel             *string   `json:"admissionLevel,omitempty"`
	AdmissionLevelAbbreviation *string   `json:"admissionLevelAbbreviation,omitempty"`
	AlbumIDs                   *[]string `json:"albumIDs,omitempty"`
	ArtistIDs                  *[]string `json:"artistIDs,omitempty"`
	AttendeeName               *string   `json:"attendeeName,omitempty"`
	AwayTeamAbbreviation       *string   `json:"awayTeamAbbreviation,omitempty"`
	AwayTeamLocation           *string   `json:"awayTeamLocation,omitempty"`
	AwayTeamName               *string   `json:"awayTeamName,omitempty"`

	// ---------- Store card ----------
	Balance *CurrencyAmount `json:"balance,omitempty"`

	// ---------- Boarding (localizable strings) ----------
	BoardingGroup          *string `json:"boardingGroup,omitempty"`
	BoardingSequenceNumber *string `json:"boardingSequenceNumber,omitempty"`
	BoardingZone           *string `json:"boardingZone,omitempty"`
	CarNumber              *string `json:"carNumber,omitempty"`
	ConfirmationNumber     *string `json:"confirmationNumber,omitempty"`

	// ---------- Dates (ISO 8601) ----------
	CurrentArrivalDate   *string `json:"currentArrivalDate,omitempty"`
	CurrentBoardingDate  *string `json:"currentBoardingDate,omitempty"`
	CurrentDepartureDate *string `json:"currentDepartureDate,omitempty"`

	// ---------- Airline / departure ----------
	DepartureAirportCode              *string                   `json:"departureAirportCode,omitempty"`
	DepartureAirportName              *string                   `json:"departureAirportName,omitempty"`
	DepartureLocationSecurityPrograms *[]TransitSecurityProgram `json:"departureLocationSecurityPrograms,omitempty"`
	DepartureLocationTimeZone         *string                   `json:"departureLocationTimeZone,omitempty"`
	DepartureCityName                 *string                   `json:"departureCityName,omitempty"`
	DepartureGate                     *string                   `json:"departureGate,omitempty"`
	DepartureLocation                 *Location                 `json:"departureLocation,omitempty"`
	DepartureLocationDescription      *string                   `json:"departureLocationDescription,omitempty"`
	DeparturePlatform                 *string                   `json:"departurePlatform,omitempty"`
	DepartureStationName              *string                   `json:"departureStationName,omitempty"`
	DepartureTerminal                 *string                   `json:"departureTerminal,omitempty"`

	// ---------- Destination ----------
	DestinationAirportCode              *string                   `json:"destinationAirportCode,omitempty"`
	DestinationAirportName              *string                   `json:"destinationAirportName,omitempty"`
	DestinationLocationSecurityPrograms *[]TransitSecurityProgram `json:"destinationLocationSecurityPrograms,omitempty"`
	DestinationLocationTimeZone         *string                   `json:"destinationLocationTimeZone,omitempty"`
	DestinationCityName                 *string                   `json:"destinationCityName,omitempty"`
	DestinationGate                     *string                   `json:"destinationGate,omitempty"`
	DestinationLocation                 *Location                 `json:"destinationLocation,omitempty"`
	DestinationLocationDescription      *string                   `json:"destinationLocationDescription,omitempty"`
	DestinationPlatform                 *string                   `json:"destinationPlatform,omitempty"`
	DestinationStationName              *string                   `json:"destinationStationName,omitempty"`
	DestinationTerminal                 *string                   `json:"destinationTerminal,omitempty"`

	// ---------- Event dates & info ----------
	Duration            *int           `json:"duration,omitempty"` // Duration in seconds
	EntranceDescription *string        `json:"entranceDescription,omitempty"`
	EventEndDate        *string        `json:"eventEndDate,omitempty"`
	EventName           *string        `json:"eventName,omitempty"`
	EventStartDate      *string        `json:"eventStartDate,omitempty"`
	EventStartDateInfo  *EventDateInfo `json:"eventStartDateInfo,omitempty"`
	EventType           *EventType     `json:"eventType,omitempty"`

	// ---------- Airline ----------
	AirlineCode  *string `json:"airlineCode,omitempty"`
	FlightCode   *string `json:"flightCode,omitempty"`
	FlightNumber *int    `json:"flightNumber,omitempty"`

	// ---------- Event / genre ----------
	Genre *string `json:"genre,omitempty"`

	// ---------- Sports ----------
	HomeTeamAbbreviation *string `json:"homeTeamAbbreviation,omitempty"`
	HomeTeamLocation     *string `json:"homeTeamLocation,omitempty"`
	HomeTeamName         *string `json:"homeTeamName,omitempty"`

	// ---------- International docs (airline) ----------
	InternationalDocumentsAreVerified             *bool   `json:"internationalDocumentsAreVerified,omitempty"`
	InternationalDocumentsVerifiedDeclarationName *string `json:"internationalDocumentsVerifiedDeclarationName,omitempty"`

	// ---------- Sports league ----------
	LeagueAbbreviation *string `json:"leagueAbbreviation,omitempty"`
	LeagueName         *string `json:"leagueName,omitempty"`

	// ---------- Lounge / membership ----------
	LoungePlaceIDs          *[]string `json:"loungePlaceIDs,omitempty"`
	MembershipProgramName   *string   `json:"membershipProgramName,omitempty"`
	MembershipProgramNumber *string   `json:"membershipProgramNumber,omitempty"`
	MembershipProgramStatus *string   `json:"membershipProgramStatus,omitempty"`

	// ---------- Original dates ----------
	OriginalArrivalDate   *string `json:"originalArrivalDate,omitempty"`
	OriginalBoardingDate  *string `json:"originalBoardingDate,omitempty"`
	OriginalDepartureDate *string `json:"originalDepartureDate,omitempty"`

	// ---------- Passenger (airline) ----------
	PassengerAirlineSSRs              *[]string                 `json:"passengerAirlineSSRs,omitempty"`
	PassengerCapabilities             *[]PassengerCapability    `json:"passengerCapabilities,omitempty"`
	PassengerEligibleSecurityPrograms *[]TransitSecurityProgram `json:"passengerEligibleSecurityPrograms,omitempty"`
	PassengerInformationSSRs          *[]string                 `json:"passengerInformationSSRs,omitempty"`
	PassengerName                     *PersonNameComponents     `json:"passengerName,omitempty"`
	PassengerServiceSSRs              *[]string                 `json:"passengerServiceSSRs,omitempty"`

	// ---------- Event performers ----------
	PerformerNames *[]string `json:"performerNames,omitempty"`
	PlaylistIDs    *[]string `json:"playlistIDs,omitempty"`

	// ---------- Boarding priority ----------
	PriorityStatus    *string `json:"priorityStatus,omitempty"`
	Seats             *[]Seat `json:"seats,omitempty"`
	SecurityScreening *string `json:"securityScreening,omitempty"`
	SilenceRequested  *bool   `json:"silenceRequested,omitempty"`

	// ---------- Sports / event ----------
	SportName         *string         `json:"sportName,omitempty"`
	TailgatingAllowed *bool           `json:"tailgatingAllowed,omitempty"`
	TicketFareClass   *string         `json:"ticketFareClass,omitempty"`
	TotalPrice        *CurrencyAmount `json:"totalPrice,omitempty"`

	// ---------- Transit ----------
	TransitProvider     *string `json:"transitProvider,omitempty"`
	TransitStatus       *string `json:"transitStatus,omitempty"`
	TransitStatusReason *string `json:"transitStatusReason,omitempty"`

	// ---------- Vehicle ----------
	VehicleName   *string `json:"vehicleName,omitempty"`
	VehicleNumber *string `json:"vehicleNumber,omitempty"`
	VehicleType   *string `json:"vehicleType,omitempty"`

	// ---------- Venue dates ----------
	VenueBoxOfficeOpenDate   *string   `json:"venueBoxOfficeOpenDate,omitempty"`
	VenueCloseDate           *string   `json:"venueCloseDate,omitempty"`
	VenueDoorsOpenDate       *string   `json:"venueDoorsOpenDate,omitempty"`
	VenueEntrance            *string   `json:"venueEntrance,omitempty"`
	VenueEntranceDoor        *string   `json:"venueEntranceDoor,omitempty"`
	VenueEntranceGate        *string   `json:"venueEntranceGate,omitempty"`
	VenueEntrancePortal      *string   `json:"venueEntrancePortal,omitempty"`
	VenueFanZoneOpenDate     *string   `json:"venueFanZoneOpenDate,omitempty"`
	VenueGatesOpenDate       *string   `json:"venueGatesOpenDate,omitempty"`
	VenueLocation            *Location `json:"venueLocation,omitempty"`
	VenueName                *string   `json:"venueName,omitempty"`
	VenueOpenDate            *string   `json:"venueOpenDate,omitempty"`
	VenueParkingLotsOpenDate *string   `json:"venueParkingLotsOpenDate,omitempty"`
	VenuePhoneNumber         *string   `json:"venuePhoneNumber,omitempty"`
	VenueRegionName          *string   `json:"venueRegionName,omitempty"`
	VenueRoom                *string   `json:"venueRoom,omitempty"`

	// ---------- Wi‑Fi ----------
	WifiAccess *[]WifiNetwork `json:"wifiAccess,omitempty"`
}
