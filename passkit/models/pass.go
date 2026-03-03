package models

// Pass represents the root pass.json object.
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

	AccessibilityURL           *string `json:"accessibilityURL,omitempty"`
	AddOnURL                   *string `json:"addOnURL,omitempty"`
	AppLaunchURL               *string `json:"appLaunchURL,omitempty"`
	AssociatedStoreIdentifiers *[]int  `json:"associatedStoreIdentifiers,omitempty"`
	AuxiliaryStoreIdentifiers  *[]int  `json:"auxiliaryStoreIdentifiers,omitempty"`
	AuthenticationToken        *string `json:"authenticationToken,omitempty"`
	BackgroundColor            *string `json:"backgroundColor,omitempty"`
	BagPolicyURL               *string `json:"bagPolicyURL,omitempty"`

	Barcodes *Barcodes `json:"barcodes,omitempty"`
	Beacons  *Beacons  `json:"beacons,omitempty"`

	ContactVenueEmail        *string `json:"contactVenueEmail,omitempty"`
	ContactVenuePhoneNumber  *string `json:"contactVenuePhoneNumber,omitempty"`
	ContactVenueWebsite      *string `json:"contactVenueWebsite,omitempty"`
	DirectionsInformationURL *string `json:"directionsInformationURL,omitempty"`
	EventLogoText            *string `json:"eventLogoText,omitempty"`
	ExpirationDate           *string `json:"expirationDate,omitempty"`

	// A background color for the footer of the pass, specified as a CSS-style RGB triple, such as rgb(100, 10, 110).
	// This key works only for poster event tickets.
	FooterBackgroundColor *string `json:"footerBackgroundColor,omitempty"`

	//A foreground color for the pass, specified as a CSS-style RGB triple, such as rgb(100, 10, 110).
	ForegroundColor    *string      `json:"foregroundColor,omitempty"`
	GroupingIdentifier *string      `json:"groupingIdentifier,omitempty"`
	LabelColor         *string      `json:"labelColor,omitempty"`
	Locations          *[]Locations `json:"locations,omitempty"`

	LogoText                     *string                         `json:"logoText,omitempty"`
	MaxDistance                  *int                            `json:"maxDistance,omitempty"`
	MerchandiseURL               *string                         `json:"merchandiseURL,omitempty"`
	NFC                          *NFC                            `json:"nfc,omitempty"`
	OrderFoodURL                 *string                         `json:"orderFoodURL,omitempty"`
	ParkingInformationURL        *string                         `json:"parkingInformationURL,omitempty"`
	PreferredStyleSchemes        *[]string                       `json:"preferredStyleSchemes,omitempty"`
	PurchaseParkingURL           *string                         `json:"purchaseParkingURL,omitempty"`
	RelevantDates                *[]RelevantDates                `json:"relevantDates,omitempty"`
	SellURL                      *string                         `json:"sellURL,omitempty"`
	Semantics                    *SemanticTags                   `json:"semantics,omitempty"`
	SharingProhibited            *bool                           `json:"sharingProhibited,omitempty"`
	SupressStripShine            *bool                           `json:"suppressStripShine,omitempty"`
	SupressHeaderDarkening       *bool                           `json:"suppressHeaderDarkening,omitempty"`
	TransferURL                  *string                         `json:"transferURL,omitempty"`
	TransitInformationURL        *string                         `json:"transitInformationURL,omitempty"`
	UseAutomaticColors           *bool                           `json:"useAutomaticColors,omitempty"`
	UserInfo                     *RawData                        `json:"userInfo,omitempty"`
	Voided                       *bool                           `json:"voided,omitempty"`
	WebServiceURL                *string                         `json:"webServiceURL,omitempty"`
	ChangeSeatURL                *string                         `json:"changeSeatURL,omitempty"`
	EntertainmentURL             *string                         `json:"entertainmentURL,omitempty"`
	PurchaseAdditionalBaggageURL *string                         `json:"purchaseAdditionalBaggageURL,omitempty"`
	PurchaseLoungeAccessURL      *string                         `json:"purchaseLoungeAccessURL,omitempty"`
	PurchaseWifiURL              *string                         `json:"purchaseWifiURL,omitempty"`
	UpgradeURL                   *string                         `json:"upgradeURL,omitempty"`
	UpcomingPassInformation      *[]UpcomingPassInformationEntry `json:"upcomingPassInformation,omitempty"`
	ManagementURL                *string                         `json:"managementURL,omitempty"`
	RegisterServiceAnimalURL     *string                         `json:"registerServiceAnimalURL,omitempty"`
	ReportLostBagURL             *string                         `json:"reportLostBagURL,omitempty"`
	RequestWheelchairURL         *string                         `json:"requestWheelchairURL,omitempty"`
	TrackBagURL                  *string                         `json:"trackBagURL,omitempty"`
	TransitProviderEmail         *string                         `json:"transitProviderEmail,omitempty"`
	TransitProviderPhoneNumber   *string                         `json:"transitProviderPhoneNumber,omitempty"`
	TransitProviderWebsiteURL    *string                         `json:"transitProviderWebsiteURL,omitempty"`

	// ---------------- PASS STYLES ----------------
	BoardingPass *BoardingPass `json:"boardingPass,omitempty"`
	Coupon       *PassFields   `json:"coupon,omitempty"`
	EventTicket  *PassFields   `json:"eventTicket,omitempty"`
	Generic      *PassFields   `json:"generic,omitempty"`
	StoreCard    *PassFields   `json:"storeCard,omitempty"`
}

// RawData is custom key-value data (e.g. for pass userInfo). Serializes as a JSON object.
type RawData map[string]interface{}
