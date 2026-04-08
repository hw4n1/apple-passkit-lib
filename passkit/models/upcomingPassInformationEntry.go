package models

type URLs struct {
	AccessibilityURL         *string `json:"accessibilityURL,omitempty"`
	AddOnURL                 *string `json:"addOnURL,omitempty"`
	BagPolicyURL             *string `json:"bagPolicyURL,omitempty"`
	ContactVenueEmail        *string `json:"contactVenueEmail,omitempty"`
	ContactVenuePhoneNumber  *string `json:"contactVenuePhoneNumber,omitempty"`
	ContactVenueWebsite      *string `json:"contactVenueWebsite,omitempty"`
	DirectionsInformationURL *string `json:"directionsInformationURL,omitempty"`
	MerchandiseURL           *string `json:"merchandiseURL,omitempty"`
	OrderFoodURL             *string `json:"orderFoodURL,omitempty"`
	ParkingInformationURL    *string `json:"parkingInformationURL,omitempty"`
	PurchaseParkingURL       *string `json:"purchaseParkingURL,omitempty"`
	SellURL                  *string `json:"sellURL,omitempty"`
	TransferURL              *string `json:"transferURL,omitempty"`
	TransitInformationURL    *string `json:"transitInformationURL,omitempty"`
}

type Images struct {
	HeaderImage *Image `json:"headerImage,omitempty"`
	VenueMap    *Image `json:"venueMap,omitempty"`
}
type Image struct {
	URLs          *ImageURLEntry `json:"URLs,omitempty"`
	ReuseExisting *bool          `json:"reuseExisting,omitempty"`
}

type ImageURLEntry struct {
	Sha256 string `json:"SHA256"`
	URL    string `json:"URL"`

	Scale *float64 `json:"scale,omitempty"`
	Size  *int     `json:"size,omitempty"`
}
type DateInformation struct {
	Date                 *string `json:"date,omitempty"`
	IgnoreTimeComponents *bool   `json:"ignoreTimeComponents,omitempty"`
	IsAllDay             *bool   `json:"isAllDay,omitempty"`
	IsUnannounced        *bool   `json:"isUnannounced,omitempty"`
	IsIndetermined       *bool   `json:"isIndetermined,omitempty"`
	TimeZone             *string `json:"timeZone,omitempty"`
}

type UpcomingPassInformationEntry struct {

	// ---------------- REQUIRED ----------------

	Identifier string `json:"identifier"`
	Name       string `json:"name"`
	Type       string `json:"type"`

	// ---------------- OPTIONAL ----------------

	URLs                      *URLs               `json:"URLs,omitempty"`
	AdditionalInfoFields      *[]PassFieldContent `json:"additionalInfoFields,omitempty"`
	AuxiliaryStoreIdentifiers *[]int              `json:"auxiliaryStoreIdentifiers,omitempty"`
	BackFields                *[]PassFieldContent `json:"backFields,omitempty"`
	DateInformation           *DateInformation    `json:"dateInformation,omitempty"`
	Images                    *Images             `json:"images,omitempty"`
	Semantics                 *SemanticTags       `json:"semantics,omitempty"`
}
